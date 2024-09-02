// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"context"
	"github.com/AssetMantle/schema/types/base"
	abciTypes "github.com/cometbft/cometbft/abci/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/orders/mappable"
)

type block struct {
	mapper helpers.Mapper
}

var _ helpers.Block = (*block)(nil)

func (block block) Begin(_ context.Context, _ abciTypes.RequestBeginBlock) {}
func (block block) End(context context.Context, _ abciTypes.RequestEndBlock) {
	orders := block.mapper.NewCollection(context)
	orders.IterateAll(func(record helpers.Record) bool {
		if mappable.GetOrder(record.GetMappable()).GetExpiryHeight().Compare(base.CurrentHeight(context)) <= 0 {
			orders.Remove(record)
		}
		return false
	})
}

//executeOrders := make(map[ids.OrderID]bool)
//orders.Iterate(
//	// TODO ***** define a proper new key
//	key.NewKey(baseIDs.PrototypeOrderID()),
//	func(Record helpers.Record) bool {
//		order := mappable.GetOrder(Record.GetMappable())
//		if order.GetExpiryHeight().Compare(baseTypes.CurrentHeight(context)) <= 0 {
//			// TODO ***** check security of sending and receiving from module and module account security
//			if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentity.GetModuleIdentityID(), order.GetMakerID(), order.GetMakerAssetID(), order.GetMakerSplit())); err != nil {
//				panic(err)
//			}
//			if _, err := block.burnAuxiliary.GetKeeper().Help(context, burn.NewAuxiliaryRequest(order.GetClassificationID())); err != nil {
//				panic(err)
//			}
//			orders.Remove(record.NewRecord(order))
//		} else {
//			// TODO ***** figure out use case
//			// // TODO ***** test
//			// id1 := baseIDs.NewOrderID(order.GetClassificationID(), order.GetMakerAssetID(), order.GetTakerAssetID(), sdkTypes.SmallestDec(), baseTypes.NewHeight(0), baseIDs.PrototypeIdentityID(), baseQualified.NewImmutables(baseLists.NewPropertyList()))
//			// // TODO ***** test
//			// id2 := baseIDs.NewOrderID(order.GetClassificationID(), order.GetTakerAssetID(), order.GetMakerAssetID(), sdkTypes.SmallestDec(), baseTypes.NewHeight(0), baseIDs.PrototypeIdentityID(), baseQualified.NewImmutables(baseLists.NewPropertyList()))
//			// if !executeOrders[id1] && !executeOrders[id2] {
//			// 	executeOrders[id1] = true
//			// }
//		}
//		return false
//	},
//)
//
//for partialOrderID := range executeOrders {
//	nextPartialOrderID := false
//
//	orders.Iterate(key.NewKey(partialOrderID),
//		func(Record helpers.Record) bool {
//			order := mappable.GetOrder(Record.GetMappable())
//			orders.Iterate(
//				key.NewKey(baseIDs.PrototypeOrderID()),
//				func(Record helpers.Record) bool {
//					executableOrder := mappable.GetOrder(Record.GetMappable())
//					var leftOrder documents.Order
//					var rightOrder documents.Order
//
//					orderHeight := order.GetExecutionHeight()
//
//					executableOrderHeight := executableOrder.GetExecutionHeight()
//
//					switch {
//					case orderHeight.Compare(executableOrderHeight) > 0:
//						leftOrder = order
//						rightOrder = executableOrder
//					case executableOrderHeight.Compare(orderHeight) > 0:
//						leftOrder = executableOrder
//						rightOrder = order
//					default:
//						leftOrder = order
//						rightOrder = executableOrder
//					}
//
//					leftOrderExchangeRate := leftOrder.GetExchangeRate()
//
//					leftOrderMakerSplit := leftOrder.GetMakerSplit()
//
//					rightOrderExchangeRate := rightOrder.GetExchangeRate()
//
//					rightOrderMakerSplit := rightOrder.GetMakerSplit()
//
//					rightOrderTakerSplitDemanded := rightOrderExchangeRate.MulTruncate(rightOrderMakerSplit.ToLegacyDec()).MulTruncate(sdkTypes.SmallestDec()).TruncateInt()
//
//					if leftOrderExchangeRate.MulTruncate(rightOrderExchangeRate).MulTruncate(sdkTypes.SmallestDec()).MulTruncate(sdkTypes.SmallestDec()).LTE(sdkTypes.OneDec()) {
//						switch {
//						case leftOrderMakerSplit.GT(rightOrderTakerSplitDemanded):
//							if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentity.GetModuleIdentityID(), leftOrder.GetMakerID(), leftOrder.GetTakerAssetID(), rightOrderMakerSplit)); err != nil {
//								panic(err)
//							}
//							if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentity.GetModuleIdentityID(), rightOrder.GetMakerID(), leftOrder.GetMakerAssetID(), rightOrderTakerSplitDemanded)); err != nil {
//								panic(err)
//							}
//
//							mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.MakerSplitProperty.GetKey(), baseData.NewNumberData(leftOrderMakerSplit.Sub(rightOrderTakerSplitDemanded))))
//
//							orders.Mutate(record.NewRecord(base.NewOrder(leftOrder.GetClassificationID(), leftOrder.GetImmutables(), leftOrder.Mutate(baseLists.AnyPropertiesToProperties(mutableProperties.Get()...)...).GetMutables())))
//							orders.Remove(record.NewRecord(rightOrder))
//
//							if executableOrderHeight.Compare(orderHeight) > 0 {
//								return true
//							}
//						case leftOrderMakerSplit.LT(rightOrderTakerSplitDemanded):
//							sendToLeftOrder := leftOrderMakerSplit.ToLegacyDec().QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(rightOrderExchangeRate)
//							if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentity.GetModuleIdentityID(), leftOrder.GetMakerID(), leftOrder.GetTakerAssetID(), sendToLeftOrder.TruncateInt())); err != nil {
//								panic(err)
//							}
//							if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentity.GetModuleIdentityID(), rightOrder.GetMakerID(), leftOrder.GetMakerAssetID(), leftOrderMakerSplit)); err != nil {
//								panic(err)
//							}
//
//							auxiliaryResponse, err := block.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.MakerSplitProperty.GetKey(), baseData.NewNumberData(rightOrderMakerSplit.Sub(sendToLeftOrder.TruncateInt()))))))
//							if err != nil {
//								panic(err)
//							}
//							mutableProperties := scrub.GetPropertiesFromResponse(auxiliaryResponse)
//
//							orders.Mutate(record.NewRecord(base.NewOrder(rightOrder.GetClassificationID(), rightOrder.GetImmutables(), rightOrder.GetMutables().Mutate(baseLists.AnyPropertiesToProperties(mutableProperties.Get()...)...))))
//							orders.Remove(record.NewRecord(leftOrder))
//
//							if orderHeight.Compare(executableOrderHeight) >= 0 {
//								return true
//							}
//						default:
//							// case leftOrderMakerSplit.Equal(rightOrderTakerSplitDemanded):
//							if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentity.GetModuleIdentityID(), leftOrder.GetMakerID(), leftOrder.GetTakerAssetID(), rightOrderMakerSplit)); err != nil {
//								panic(err)
//							}
//							if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentity.GetModuleIdentityID(), rightOrder.GetMakerID(), leftOrder.GetMakerAssetID(), leftOrderMakerSplit)); err != nil {
//								panic(err)
//							}
//
//							orders.Remove(record.NewRecord(rightOrder))
//							orders.Remove(record.NewRecord(leftOrder))
//							return true
//						}
//					} else {
//						nextPartialOrderID = true
//						return true
//					}
//					return false
//				},
//			)
//			return nextPartialOrderID
//		})
//
//	if nextPartialOrderID {
//		continue
//	}
//}

func (block block) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ ...interface{}) helpers.Block {
	block.mapper = mapper
	return block
}

func Prototype() helpers.Block {
	return block{}
}
