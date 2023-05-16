// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"context"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/burn"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/mappable"
	"github.com/AssetMantle/modules/x/orders/module"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type block struct {
	mapper              helpers.Mapper
	parameterManager    helpers.ParameterManager
	burnAuxiliary       helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
	transferAuxiliary   helpers.Auxiliary
	scrubAuxiliary      helpers.Auxiliary
}

var _ helpers.Block = (*block)(nil)

func (block block) Begin(_ context.Context, _ abciTypes.RequestBeginBlock) {

}

func (block block) End(context context.Context, _ abciTypes.RequestEndBlock) {
	executeOrders := make(map[ids.OrderID]bool)
	orders := block.mapper.NewCollection(context)

	orders.Iterate(
		// TODO ***** define a proper new key
		key.NewKey(baseIDs.PrototypeOrderID()),
		func(Mappable helpers.Mappable) bool {
			order := mappable.GetOrder(Mappable)

			if order.GetExpiryHeight().Compare(baseTypes.CurrentHeight(context)) <= 0 {
				// TODO ***** check security of sending and receiving from module and module account security
				if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, order.GetMakerID(), order.GetMakerOwnableID(), order.GetMakerOwnableSplit())); err != nil {
					panic(err)
				}
				if _, err := block.burnAuxiliary.GetKeeper().Help(context, burn.NewAuxiliaryRequest(order.GetClassificationID())); err != nil {
					panic(err)
				}
				orders.Remove(mappable.NewMappable(order))
			} else {
				// TODO ***** figure out use case
				// // TODO ***** test
				// id1 := baseIDs.NewOrderID(order.GetClassificationID(), order.GetMakerOwnableID(), order.GetTakerOwnableID(), sdkTypes.SmallestDec(), baseTypes.NewHeight(0), baseIDs.PrototypeIdentityID(), baseQualified.NewImmutables(baseLists.NewPropertyList()))
				// // TODO ***** test
				// id2 := baseIDs.NewOrderID(order.GetClassificationID(), order.GetTakerOwnableID(), order.GetMakerOwnableID(), sdkTypes.SmallestDec(), baseTypes.NewHeight(0), baseIDs.PrototypeIdentityID(), baseQualified.NewImmutables(baseLists.NewPropertyList()))
				// if !executeOrders[id1] && !executeOrders[id2] {
				// 	executeOrders[id1] = true
				// }
			}
			return false
		},
	)

	for partialOrderID := range executeOrders {
		nextPartialOrderID := false

		orders.Iterate(key.NewKey(partialOrderID),
			func(Mappable helpers.Mappable) bool {
				order := mappable.GetOrder(Mappable)
				orders.Iterate(
					key.NewKey(baseIDs.PrototypeOrderID()),
					func(Mappable helpers.Mappable) bool {
						executableOrder := mappable.GetOrder(Mappable)
						var leftOrder documents.Order
						var rightOrder documents.Order

						orderHeight := order.GetCreationHeight()

						executableOrderHeight := executableOrder.GetCreationHeight()

						switch {
						case orderHeight.Compare(executableOrderHeight) > 0:
							leftOrder = order
							rightOrder = executableOrder
						case executableOrderHeight.Compare(orderHeight) > 0:
							leftOrder = executableOrder
							rightOrder = order
						default:
							leftOrder = order
							rightOrder = executableOrder
						}

						leftOrderExchangeRate := leftOrder.GetExchangeRate()

						leftOrderMakerOwnableSplit := leftOrder.GetMakerOwnableSplit()

						rightOrderExchangeRate := rightOrder.GetExchangeRate()

						rightOrderMakerOwnableSplit := rightOrder.GetMakerOwnableSplit()

						rightOrderTakerOwnableSplitDemanded := rightOrderExchangeRate.MulTruncate(rightOrderMakerOwnableSplit.ToDec()).MulTruncate(sdkTypes.SmallestDec()).TruncateInt()

						if leftOrderExchangeRate.MulTruncate(rightOrderExchangeRate).MulTruncate(sdkTypes.SmallestDec()).MulTruncate(sdkTypes.SmallestDec()).LTE(sdkTypes.OneDec()) {
							switch {
							case leftOrderMakerOwnableSplit.GT(rightOrderTakerOwnableSplitDemanded):
								if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, leftOrder.GetMakerID(), leftOrder.GetTakerOwnableID(), rightOrderMakerOwnableSplit)); err != nil {
									panic(err)
								}
								if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, rightOrder.GetMakerID(), leftOrder.GetMakerOwnableID(), rightOrderTakerOwnableSplitDemanded)); err != nil {
									panic(err)
								}

								mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewNumberData(leftOrderMakerOwnableSplit.Sub(rightOrderTakerOwnableSplitDemanded))))

								orders.Mutate(mappable.NewMappable(base.NewOrder(leftOrder.GetClassificationID(), leftOrder.GetImmutables(), leftOrder.Mutate(baseLists.AnyPropertiesToProperties(mutableProperties.Get()...)...).GetMutables())))
								orders.Remove(mappable.NewMappable(rightOrder))

								if executableOrderHeight.Compare(orderHeight) > 0 {
									return true
								}
							case leftOrderMakerOwnableSplit.LT(rightOrderTakerOwnableSplitDemanded):
								sendToLeftOrder := leftOrderMakerOwnableSplit.ToDec().QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(rightOrderExchangeRate)
								if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, leftOrder.GetMakerID(), leftOrder.GetTakerOwnableID(), sendToLeftOrder.TruncateInt())); err != nil {
									panic(err)
								}
								if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, rightOrder.GetMakerID(), leftOrder.GetMakerOwnableID(), leftOrderMakerOwnableSplit)); err != nil {
									panic(err)
								}

								auxiliaryResponse, err := block.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewNumberData(rightOrderMakerOwnableSplit.Sub(sendToLeftOrder.TruncateInt()))))))
								if err != nil {
									panic(err)
								}
								mutableProperties := scrub.GetPropertiesFromResponse(auxiliaryResponse)

								orders.Mutate(mappable.NewMappable(base.NewOrder(rightOrder.GetClassificationID(), rightOrder.GetImmutables(), rightOrder.GetMutables().Mutate(baseLists.AnyPropertiesToProperties(mutableProperties.Get()...)...))))
								orders.Remove(mappable.NewMappable(leftOrder))

								if orderHeight.Compare(executableOrderHeight) >= 0 {
									return true
								}
							default:
								// case leftOrderMakerOwnableSplit.Equal(rightOrderTakerOwnableSplitDemanded):
								if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, leftOrder.GetMakerID(), leftOrder.GetTakerOwnableID(), rightOrderMakerOwnableSplit)); err != nil {
									panic(err)
								}
								if _, err := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, rightOrder.GetMakerID(), leftOrder.GetMakerOwnableID(), leftOrderMakerOwnableSplit)); err != nil {
									panic(err)
								}

								orders.Remove(mappable.NewMappable(rightOrder))
								orders.Remove(mappable.NewMappable(leftOrder))
								return true
							}
						} else {
							nextPartialOrderID = true
							return true
						}
						return false
					},
				)
				return nextPartialOrderID
			})

		if nextPartialOrderID {
			continue
		}
	}
}

func Prototype() helpers.Block {
	return block{}
}

func (block block) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaryKeepers ...interface{}) helpers.Block {
	block.mapper, block.parameterManager = mapper, parameterManager

	for _, auxiliaryKeeper := range auxiliaryKeepers {
		switch value := auxiliaryKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case supplement.Auxiliary.GetName():
				block.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				block.transferAuxiliary = value
			case scrub.Auxiliary.GetName():
				block.scrubAuxiliary = value
			case burn.Auxiliary.GetName():
				block.burnAuxiliary = value
			}
		}
	}

	return block
}
