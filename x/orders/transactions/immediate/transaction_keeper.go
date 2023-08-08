// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

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
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/mappable"
	"github.com/AssetMantle/modules/x/orders/record"
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

	makerSplit, ok := sdkTypes.NewIntFromString(message.MakerSplit)
	if !ok {
		return nil, errorConstants.InvalidParameter.Wrapf("maker split %s is not a valid integer", message.MakerSplit)
	}

	takerSplit, ok := sdkTypes.NewIntFromString(message.TakerSplit)
	if !ok {
		return nil, errorConstants.InvalidParameter.Wrapf("taker split %s is not a valid integer", message.TakerSplit)
	}

	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, constants.ModuleIdentityID, message.MakerAssetID, makerSplit)); err != nil {
		return nil, err
	}
	immutableMetaProperties := message.ImmutableMetaProperties.
		Add(baseProperties.NewMetaProperty(propertyConstants.ExchangeRateProperty.GetKey(), baseData.NewNumberData(takerSplit.Quo(sdkTypes.OneInt()).Quo(makerSplit)))).
		Add(baseProperties.NewMetaProperty(propertyConstants.CreationHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(sdkTypes.UnwrapSDKContext(context).BlockHeight())))).
		Add(baseProperties.NewMetaProperty(propertyConstants.MakerAssetIDProperty.GetKey(), baseData.NewIDData(message.MakerAssetID))).
		Add(baseProperties.NewMetaProperty(propertyConstants.TakerAssetIDProperty.GetKey(), baseData.NewIDData(message.TakerAssetID))).
		Add(baseProperties.NewMetaProperty(propertyConstants.MakerIDProperty.GetKey(), baseData.NewIDData(message.FromID))).
		Add(baseProperties.NewMetaProperty(propertyConstants.TakerIDProperty.GetKey(), baseData.NewIDData(message.TakerID)))

	immutables := baseQualified.NewImmutables(immutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.ImmutableProperties.Get()...)...))
	orderID := baseIDs.NewOrderID(message.ClassificationID, immutables)
	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(orderID))

	if order := orders.GetMappable(key.NewKey(orderID)); order != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("order with ID %s already exists", orderID.AsString())
	}

	if message.ExpiresIn.Get() > transactionKeeper.parameterManager.Fetch(context).GetParameter(propertyConstants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get() {
		return nil, errorConstants.InvalidRequest.Wrapf("order expiry exceeds maximum allowed %d", transactionKeeper.parameterManager.Fetch(context).GetParameter(propertyConstants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get())
	}

	mutableMetaProperties := message.MutableMetaProperties.
		Add(baseProperties.NewMetaProperty(propertyConstants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(message.ExpiresIn.Get()+sdkTypes.UnwrapSDKContext(context).BlockHeight())))).
		Add(baseProperties.NewMetaProperty(propertyConstants.MakerSplitProperty.GetKey(), baseData.NewNumberData(makerSplit)))

	mutables := baseQualified.NewMutables(mutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.MutableProperties.Get()...)...))

	if _, err := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); err != nil {
		return nil, err
	}

	order := base.NewOrder(message.ClassificationID, immutables, mutables)
	orders = orders.Add(record.NewRecord(order))

	// Order execution
	orderMutated := false
	orderLeftOverMakerSplit := makerSplit

	accumulator := func(Record helpers.Record) bool {
		executableOrder := mappable.GetOrder(Record.GetMappable())

		executableOrderTakerSplitDemanded := executableOrder.GetExchangeRate().MulTruncate(executableOrder.GetMakerSplit().ToDec()).MulTruncate(sdkTypes.SmallestDec()).TruncateInt()

		if order.GetExchangeRate().MulTruncate(executableOrder.GetExchangeRate()).MulTruncate(sdkTypes.SmallestDec()).MulTruncate(sdkTypes.SmallestDec()).LTE(sdkTypes.OneDec()) {
			switch {
			case orderLeftOverMakerSplit.GT(executableOrderTakerSplitDemanded):
				// sending to buyer
				if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentityID, order.GetMakerID(), order.GetTakerAssetID(), executableOrder.GetMakerSplit())); err != nil {
					panic(err)
				}
				// sending to executableOrder
				if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentityID, executableOrder.GetMakerID(), order.GetMakerAssetID(), executableOrderTakerSplitDemanded)); err != nil {
					panic(err)
				}

				orderLeftOverMakerSplit = orderLeftOverMakerSplit.Sub(executableOrderTakerSplitDemanded)

				orders.Remove(record.NewRecord(executableOrder))
			case orderLeftOverMakerSplit.LT(executableOrderTakerSplitDemanded):
				// sending to buyer
				sendToBuyer := orderLeftOverMakerSplit.Quo(sdkTypes.OneInt()).ToDec().QuoTruncate(executableOrder.GetExchangeRate()).TruncateInt()
				if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentityID, order.GetMakerID(), order.GetTakerAssetID(), sendToBuyer)); err != nil {
					panic(err)
				}
				// sending to executableOrder
				if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentityID, executableOrder.GetMakerID(), order.GetMakerAssetID(), orderLeftOverMakerSplit)); err != nil {
					panic(err)
				}

				mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(propertyConstants.MakerSplitProperty.GetKey(), baseData.NewNumberData(executableOrder.GetMakerSplit().Sub(sendToBuyer))))

				orders.Mutate(record.NewRecord(base.NewOrder(executableOrder.GetClassificationID(), executableOrder.GetImmutables(), executableOrder.GetMutables().Mutate(baseLists.AnyPropertiesToProperties(mutableProperties.Get()...)...))))

				orderLeftOverMakerSplit = sdkTypes.ZeroInt()
			default:
				// case orderLeftOverMakerSplit.Equal(executableOrderTakerSplitDemanded):
				// sending to buyer
				if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentityID, order.GetMakerID(), order.GetTakerAssetID(), executableOrder.GetMakerSplit())); err != nil {
					panic(err)
				}
				// sending to seller
				if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentityID, executableOrder.GetMakerID(), order.GetMakerAssetID(), orderLeftOverMakerSplit)); err != nil {
					panic(err)
				}

				orders.Remove(record.NewRecord(executableOrder))

				orderLeftOverMakerSplit = sdkTypes.ZeroInt()
			}

			orderMutated = true
		}

		if orderLeftOverMakerSplit.Equal(sdkTypes.ZeroInt()) {
			orders.Remove(record.NewRecord(order))
			return true
		}

		return false
	}

	orders.Iterate(record.NewRecord(order).GetKey(), accumulator)

	if !orderLeftOverMakerSplit.Equal(sdkTypes.ZeroInt()) && orderMutated {
		mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(propertyConstants.MakerSplitProperty.GetKey(), baseData.NewNumberData(orderLeftOverMakerSplit)))

		orders.Mutate(record.NewRecord(base.NewOrder(order.GetClassificationID(), order.GetImmutables(), order.GetMutables().Mutate(baseLists.AnyPropertiesToProperties(mutableProperties.Get()...)...))))
	}

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
