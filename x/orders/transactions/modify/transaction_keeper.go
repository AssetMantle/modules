// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"context"

	baseLists "github.com/AssetMantle/schema/go/lists/base"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	propertyConstants "github.com/AssetMantle/schema/go/properties/constants"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/mappable"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	conformAuxiliary      helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	fromAddress, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}

	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.OrderID))

	Mappable := orders.GetMappable(key.NewKey(message.OrderID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("order with ID %s not found", message.OrderID.AsString())
	}
	order := mappable.GetOrder(Mappable)
	makerOwnableSplit, ok := sdkTypes.NewIntFromString(message.MakerOwnableSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("invalid maker ownable split %s", message.MakerOwnableSplit)
	}
	transferMakerOwnableSplit := makerOwnableSplit.Sub(order.GetMakerOwnableSplit())

	if transferMakerOwnableSplit.LT(sdkTypes.ZeroInt()) {
		if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentityID, message.FromID, order.GetMakerOwnableID(), transferMakerOwnableSplit.Abs())); err != nil {
			return nil, err
		}
	} else if transferMakerOwnableSplit.GT(sdkTypes.ZeroInt()) {
		if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, constants.ModuleIdentityID, order.GetMakerOwnableID(), transferMakerOwnableSplit)); err != nil {
			return nil, err
		}
	}

	mutableMetaProperties := message.MutableMetaProperties.
		Add(baseProperties.NewMetaProperty(propertyConstants.MakerOwnableSplitProperty.GetKey(), baseData.NewNumberData(makerOwnableSplit))).
		Add(baseProperties.NewMetaProperty(propertyConstants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(message.ExpiresIn.Get()+sdkTypes.UnwrapSDKContext(context).BlockHeight()))))

	updatedMutables := order.GetMutables().Mutate(baseLists.AnyPropertiesToProperties(append(mutableMetaProperties.Get(), message.MutableProperties.Get()...)...)...)

	if _, err := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(order.GetClassificationID(), order.GetImmutables(), updatedMutables)); err != nil {
		return nil, err
	}

	orders.Remove(mappable.NewMappable(order))
	orders.Add(mappable.NewMappable(base.NewOrder(order.GetClassificationID(), order.GetImmutables(), updatedMutables)))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

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
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
