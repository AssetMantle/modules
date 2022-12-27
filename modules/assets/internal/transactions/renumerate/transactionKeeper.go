// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"context"
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/documents"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/types"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	maintainAuxiliary     helpers.Auxiliary
	renumerateAuxiliary   helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context types.Context, message helpers.Message) helpers.TransactionResponse {
	_, err := transactionKeeper.Handle(context.Context(), message.(*Message))
	return newTransactionResponse(err)
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*Response, error) {

	fromAddress, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	assets := transactionKeeper.mapper.NewCollection(types.UnwrapSDKContext(context)).Fetch(key.NewKey(message.AssetID))

	Mappable := assets.Get(key.NewKey(message.AssetID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound
	}
	asset := Mappable.(documents.Asset)

	if auxiliaryResponse := transactionKeeper.maintainAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), maintain.NewAuxiliaryRequest(asset.GetClassificationID(), message.FromID, base.NewMutables(baseLists.NewPropertyList(constants.SupplyProperty)))); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	metaProperties, err := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), supplement.NewAuxiliaryRequest(asset.GetSupply())))
	if err != nil {
		return nil, err
	}

	if supplyMetaProperty := metaProperties.GetProperty(constants.SupplyProperty.GetID()); supplyMetaProperty != nil && supplyMetaProperty.IsMeta() {
		value := supplyMetaProperty.(properties.MetaProperty).GetData().(data.DecData).Get()
		if auxiliaryResponse := transactionKeeper.renumerateAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), renumerate.NewAuxiliaryRequest(message.FromID, message.AssetID, value)); !auxiliaryResponse.IsSuccessful() {
			return nil, auxiliaryResponse.GetError()
		}
	} else {
		return nil, errorConstants.MetaDataError
	}

	return &Response{}, nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case maintain.Auxiliary.GetName():
				transactionKeeper.maintainAuxiliary = value
			case renumerate.Auxiliary.GetName():
				transactionKeeper.renumerateAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			}
		default:
			panic(errorConstants.UninitializedUsage)
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
