// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/go/documents/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/properties/constants"

	"github.com/AssetMantle/modules/simulation/simulatedDatabase/assets"
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/orders"
	mappableAssets "github.com/AssetMantle/modules/x/assets/mappable"
	mappableOrders "github.com/AssetMantle/modules/x/orders/mappable"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/orders/genesis"
	ordersModule "github.com/AssetMantle/modules/x/orders/module"
	"github.com/AssetMantle/modules/x/orders/parameters/maxOrderLife"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		maxOrderLife.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, len(assets.ClassificationIDMappableBytesMap))
	index := 0
	var classificationIDString string

	orders.ClearAll()
	for i := 0; i < len(assets.ClassificationIDMappableBytesMap); i++ {
		assetMap := assets.GetAssetData(simulationState.Accounts[i].Address.String())
		if assetMap == nil {
			continue
		}
		for class, _ := range assetMap {
			classificationIDString = class
		}
		mappable := &mappableAssets.Mappable{}
		baseHelpers.CodecPrototype().MustUnmarshal(assets.ClassificationIDMappableBytesMap[classificationIDString], mappable)
		immutables := baseQualified.NewImmutables(mappable.Asset.Immutables.GetImmutablePropertyList().Add(baseLists.AnyPropertiesToProperties(constants.ExchangeRateProperty.ToAnyProperty(),
			constants.CreationHeightProperty.ToAnyProperty(),
			constants.MakerOwnableIDProperty.ToAnyProperty(),
			constants.TakerOwnableIDProperty.ToAnyProperty(),
			constants.MakerIDProperty.ToAnyProperty(),
			constants.TakerIDProperty.ToAnyProperty())...))
		mutables := baseQualified.NewMutables(mappable.Asset.Mutables.GetMutablePropertyList().Add(baseLists.AnyPropertiesToProperties(
			constants.ExpiryHeightProperty.ToAnyProperty(),
			constants.MakerOwnableSplitProperty.ToAnyProperty(),
		)...))
		classificationID := baseIDs.NewClassificationID(immutables, mutables)
		orderID := baseIDs.NewOrderID(classificationID, immutables)
		order := base.NewOrder(classificationID, immutables, mutables)
		mappableList[index] = mappableOrders.NewMappable(order)
		orders.AddOrderData(simulationState.Accounts[index].Address.String(), classificationID.AsString(), orderID.AsString())
		orders.AddMappableBytes(classificationID.AsString(), baseHelpers.CodecPrototype().MustMarshal(mappableOrders.NewMappable(order)))
		index++
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseLists.NewParameterList(maxOrderLife.Parameter.Mutate(Data)))

	simulationState.GenState[ordersModule.Name] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
