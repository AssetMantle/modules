// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/mappables"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type block struct {
	mapper              helpers.Mapper
	parameters          helpers.Parameters
	supplementAuxiliary helpers.Auxiliary
	transferAuxiliary   helpers.Auxiliary
	scrubAuxiliary      helpers.Auxiliary
}

var _ helpers.Block = (*block)(nil)

func (block block) Begin(_ sdkTypes.Context, _ abciTypes.RequestBeginBlock) {

}

func (block block) End(context sdkTypes.Context, _ abciTypes.RequestEndBlock) {
	executeOrders := make(map[ids.OrderID]bool)
	orders := block.mapper.NewCollection(context)

	orders.Iterate(
		// TODO ***** test this case
		key.NewKey(baseIDs.NewOrderID(nil, nil, nil, nil, nil, nil, nil)),
		func(order helpers.Mappable) bool {
			metaProperties, err := supplement.GetMetaPropertiesFromResponse(block.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.(mappables.Order).GetExpiry(), order.(mappables.Order).GetMakerOwnableSplit())))
			if err != nil {
				panic(err)
			}
			if expiryProperty := metaProperties.GetMetaProperty(constants.ExpiryProperty); expiryProperty != nil {
				expiry := expiryProperty.GetData().(data.HeightData).Get()

				if expiry.Compare(baseTypes.NewHeight(context.BlockHeight())) <= 0 {
					makerOwnableSplitProperty := metaProperties.GetMetaProperty(constants.MakerOwnableSplitProperty)
					if makerOwnableSplitProperty == nil {
						panic(errorConstants.MetaDataError)
					}
					makerOwnableSplit := makerOwnableSplitProperty.GetData().(data.DecData).Get()
					// TODO ***** check security of sending and receiving from module and module account security
					if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, order.(mappables.Order).GetMakerID(), order.(mappables.Order).GetMakerOwnableID(), makerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
						panic(auxiliaryResponse.GetError())
					}
					orders.Remove(order)
				} else {
					// TODO ***** test
					id1 := baseIDs.NewOrderID(order.(mappables.Order).GetClassificationID(), order.(mappables.Order).GetMakerOwnableID(), order.(mappables.Order).GetTakerOwnableID(), baseIDs.NewStringID(""), baseIDs.NewStringID(""), baseIDs.NewIdentityID(nil, nil), baseQualified.NewImmutables(baseLists.NewPropertyList()))
					// TODO ***** test
					id2 := baseIDs.NewOrderID(order.(mappables.Order).GetClassificationID(), order.(mappables.Order).GetTakerOwnableID(), order.(mappables.Order).GetMakerOwnableID(), baseIDs.NewStringID(""), baseIDs.NewStringID(""), baseIDs.NewIdentityID(nil, nil), baseQualified.NewImmutables(baseLists.NewPropertyList()))
					if !executeOrders[id1] && !executeOrders[id2] {
						executeOrders[id1] = true
					}
				}
			}
			return false
		},
	)

	for partialOrderID := range executeOrders {
		nextPartialOrderID := false

		orders.Iterate(key.NewKey(partialOrderID), func(orderMappable helpers.Mappable) bool {
			orders.Iterate(
				key.NewKey(baseIDs.NewOrderID(orderMappable.(mappables.Order).GetClassificationID(), orderMappable.(mappables.Order).GetTakerOwnableID(), orderMappable.(mappables.Order).GetMakerOwnableID(), baseIDs.NewStringID(""), baseIDs.NewStringID(""), baseIDs.NewIdentityID(nil, nil), baseQualified.NewImmutables(baseLists.NewPropertyList()))),
				func(executableMappableOrder helpers.Mappable) bool {

					var leftOrder mappables.Order
					var rightOrder mappables.Order

					orderHeight := orderMappable.(mappables.Order).GetCreation().GetData().(data.HeightData).Get()

					executableOrderHeight := executableMappableOrder.(mappables.Order).GetCreation().GetData().(data.HeightData).Get()

					switch {
					case orderHeight.Compare(executableOrderHeight) > 0:
						leftOrder = orderMappable.(mappables.Order)
						rightOrder = executableMappableOrder.(mappables.Order)
					case executableOrderHeight.Compare(orderHeight) > 0:
						leftOrder = executableMappableOrder.(mappables.Order)
						rightOrder = orderMappable.(mappables.Order)
					default:
						leftOrder = orderMappable.(mappables.Order)
						rightOrder = executableMappableOrder.(mappables.Order)
					}

					leftOrderExchangeRate := leftOrder.GetExchangeRate().GetData().(data.DecData).Get()

					leftOrderMetaProperties, err := supplement.GetMetaPropertiesFromResponse(block.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(leftOrder.GetMakerOwnableSplit())))
					if err != nil {
						panic(err)
					}

					leftOrderMakerOwnableSplit := leftOrderMetaProperties.GetMetaProperty(constants.MakerOwnableSplitProperty).GetData().(data.DecData).Get()

					rightOrderExchangeRate := rightOrder.GetExchangeRate().GetData().(data.DecData).Get()

					rightOrderMetaProperties, err := supplement.GetMetaPropertiesFromResponse(block.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(rightOrder.GetMakerOwnableSplit())))
					if err != nil {
						panic(err)
					}

					rightOrderMakerOwnableSplit := rightOrderMetaProperties.GetMetaProperty(constants.MakerOwnableSplitProperty).GetData().(data.DecData).Get()

					rightOrderTakerOwnableSplitDemanded := rightOrderExchangeRate.MulTruncate(rightOrderMakerOwnableSplit).MulTruncate(sdkTypes.SmallestDec())

					if leftOrderExchangeRate.MulTruncate(rightOrderExchangeRate).MulTruncate(sdkTypes.SmallestDec()).MulTruncate(sdkTypes.SmallestDec()).LTE(sdkTypes.OneDec()) {
						switch {
						case leftOrderMakerOwnableSplit.GT(rightOrderTakerOwnableSplitDemanded):
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, leftOrder.GetMakerID(), leftOrder.GetTakerOwnableID(), rightOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, rightOrder.GetMakerID(), leftOrder.GetMakerOwnableID(), rightOrderTakerOwnableSplitDemanded)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}

							mutableProperties, err := scrub.GetPropertiesFromResponse(block.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(leftOrderMakerOwnableSplit.Sub(rightOrderTakerOwnableSplitDemanded))))))
							if err != nil {
								panic(err)
							}

							orders.Mutate(mappable.NewOrder(leftOrder.GetClassificationID(), leftOrder.GetImmutables(), leftOrder.Mutate(mutableProperties.GetList()...).GetMutables()))
							orders.Remove(rightOrder)

							if executableOrderHeight.Compare(orderHeight) > 0 {
								return true
							}
						case leftOrderMakerOwnableSplit.LT(rightOrderTakerOwnableSplitDemanded):
							sendToLeftOrder := leftOrderMakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(rightOrderExchangeRate)
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, leftOrder.GetMakerID(), leftOrder.GetTakerOwnableID(), sendToLeftOrder)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, rightOrder.GetMakerID(), leftOrder.GetMakerOwnableID(), leftOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}

							mutableProperties, err := scrub.GetPropertiesFromResponse(block.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(rightOrderMakerOwnableSplit.Sub(sendToLeftOrder))))))
							if err != nil {
								panic(err)
							}

							orders.Mutate(mappable.NewOrder(rightOrder.GetClassificationID(), rightOrder.GetImmutables(), rightOrder.GetMutables().Mutate(mutableProperties.GetList()...)))
							orders.Remove(leftOrder)

							if orderHeight.Compare(executableOrderHeight) >= 0 {
								return true
							}
						default:
							// case leftOrderMakerOwnableSplit.Equal(rightOrderTakerOwnableSplitDemanded):
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, leftOrder.GetMakerID(), leftOrder.GetTakerOwnableID(), rightOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(module.ModuleIdentityID, rightOrder.GetMakerID(), leftOrder.GetMakerOwnableID(), leftOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}

							orders.Remove(rightOrder)
							orders.Remove(leftOrder)
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

func (block block) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaryKeepers ...interface{}) helpers.Block {
	block.mapper, block.parameters = mapper, parameters

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
			}
		default:
			panic(errorConstants.UninitializedUsage)
		}
	}

	return block
}
