// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"context"
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/mint"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/types"
)

type transactionKeeper struct {
	mapper                     helpers.Mapper
	parameters                 helpers.Parameters
	conformAuxiliary           helpers.Auxiliary
	mintAuxiliary              helpers.Auxiliary
	authenticateAuxiliary      helpers.Auxiliary
	maintainersVerifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context types.Context, message helpers.Message) helpers.TransactionResponse {
	_, err := transactionKeeper.Handle(context.Context(), message.(*Message))
	return newTransactionResponse(err)
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*Response, error) {

	if auxiliaryResponse := transactionKeeper.maintainersVerifyAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), verify.NewAuxiliaryRequest(message.ClassificationID, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	fromAddress, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(append(message.ImmutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...))

	assetID := baseIDs.NewAssetID(message.ClassificationID, immutables)

	assets := transactionKeeper.mapper.NewCollection(types.UnwrapSDKContext(context)).Fetch(key.NewKey(assetID))
	if assets.Get(key.NewKey(assetID)) != nil {
		return nil, errorConstants.EntityAlreadyExists
	}

	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(append(message.MutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...))

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	split := types.SmallestDec()

	if metaPropertyList := baseLists.NewPropertyList(append(message.ImmutableMetaProperties.GetList(), message.MutableMetaProperties.GetList()...)...); metaPropertyList.GetProperty(constants.SupplyProperty.GetID()) != nil {
		split = metaPropertyList.GetProperty(constants.SupplyProperty.GetID()).(properties.MetaProperty).GetData().(data.DecData).Get()
	}

	if auxiliaryResponse := transactionKeeper.mintAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), mint.NewAuxiliaryRequest(message.ToID, assetID, split)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	assets.Add(mappable.NewMappable(base.NewAsset(message.ClassificationID, immutables, mutables)))

	return &Response{}, nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case mint.Auxiliary.GetName():
				transactionKeeper.mintAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.maintainersVerifyAuxiliary = value
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
