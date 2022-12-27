// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"context"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/cosmos/cosmos-sdk/types"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameters            helpers.Parameters
	conformAuxiliary      helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
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

	orders := transactionKeeper.mapper.NewCollection(types.UnwrapSDKContext(context)).Fetch(key.NewKey(message.OrderID))

	Mappable := orders.Get(key.NewKey(message.OrderID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound
	}
	order := Mappable.(documents.Order)

	transferMakerOwnableSplit := message.MakerOwnableSplit.Sub(order.GetMakerOwnableSplit())

	if transferMakerOwnableSplit.LT(types.ZeroDec()) {
		if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(module.ModuleIdentityID, message.FromID, order.GetMakerOwnableID(), transferMakerOwnableSplit.Abs())); !auxiliaryResponse.IsSuccessful() {
			return nil, auxiliaryResponse.GetError()
		}
	} else if transferMakerOwnableSplit.GT(types.ZeroDec()) {
		if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), transfer.NewAuxiliaryRequest(message.FromID, module.ModuleIdentityID, order.GetMakerOwnableID(), transferMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
			return nil, auxiliaryResponse.GetError()
		}
	}

	mutableMetaProperties := message.MutableMetaProperties.
		Add(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(message.MakerOwnableSplit))).
		Add(baseProperties.NewMetaProperty(constants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(message.ExpiresIn.Get()+types.UnwrapSDKContext(context).BlockHeight()))))

	updatedMutables := order.GetMutables().Mutate(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), conform.NewAuxiliaryRequest(order.GetClassificationID(), order.GetImmutables(), updatedMutables)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	orders.Remove(mappable.NewMappable(order))
	orders.Add(mappable.NewMappable(base.NewOrder(order.GetClassificationID(), order.GetImmutables(), updatedMutables)))

	return &Response{}, nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, externalKeeper := range auxiliaries {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
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
