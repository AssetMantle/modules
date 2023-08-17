// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"context"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	propertyConstants "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/record"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
}

// TODO move to proper package
var PutOrderClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(propertyConstants.MakerIDProperty, propertyConstants.MakerAssetIDProperty, propertyConstants.TakerAssetIDProperty, propertyConstants.MakerSplitProperty, propertyConstants.TakerSplitProperty, propertyConstants.ExpiryHeightProperty)), baseQualified.NewMutables(baseLists.NewPropertyList()))

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

	makerSplit, ok := sdkTypes.NewIntFromString(message.MakerSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("maker split is not a valid integer")
	}

	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, constants.ModuleIdentity.GetModuleIdentityID(), message.MakerAssetID, makerSplit)); err != nil {
		return nil, err
	}

	takerSplit, ok := sdkTypes.NewIntFromString(message.TakerSplit)
	if !ok {
		return nil, errorConstants.IncorrectFormat.Wrapf("taker split is not a valid integer")
	}

	if message.ExpiryHeight.Compare(baseTypes.CurrentHeight(context)) <= 0 {
		return nil, errorConstants.InvalidRequest.Wrapf("order expiry is in the past")
	} else if message.ExpiryHeight.Get()-baseTypes.CurrentHeight(context).Get() > transactionKeeper.parameterManager.Fetch(context).GetParameter(propertyConstants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get() {
		return nil, errorConstants.InvalidRequest.Wrapf("order expiry exceeds maximum allowed %d", transactionKeeper.parameterManager.Fetch(context).GetParameter(propertyConstants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get())
	}

	immutables := baseQualified.NewImmutables(baseLists.PrototypePropertyList().
		Add(baseProperties.NewMetaProperty(propertyConstants.MakerIDProperty.GetKey(), baseData.NewIDData(message.FromID))).
		Add(baseProperties.NewMetaProperty(propertyConstants.MakerAssetIDProperty.GetKey(), baseData.NewIDData(message.MakerAssetID))).
		Add(baseProperties.NewMetaProperty(propertyConstants.TakerAssetIDProperty.GetKey(), baseData.NewIDData(message.TakerAssetID))).
		Add(baseProperties.NewMetaProperty(propertyConstants.MakerSplitProperty.GetKey(), baseData.NewNumberData(makerSplit))).
		Add(baseProperties.NewMetaProperty(propertyConstants.TakerSplitProperty.GetKey(), baseData.NewNumberData(takerSplit))).
		Add(baseProperties.NewMetaProperty(propertyConstants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(message.ExpiryHeight))))

	orderID := baseIDs.NewOrderID(PutOrderClassificationID, immutables)

	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(orderID))
	if orders.GetMappable(key.NewKey(orderID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("order with ID %s already exists", orderID.AsString())
	}

	orders.Add(record.NewRecord(base.NewOrder(PutOrderClassificationID, immutables, baseQualified.NewMutables(baseLists.PrototypePropertyList()))))

	return newTransactionResponse(orderID), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, externalKeeper := range auxiliaries {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			}
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
