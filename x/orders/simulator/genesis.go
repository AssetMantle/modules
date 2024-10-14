// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/parameters"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/simulation/simulated_database/assets"
	"github.com/AssetMantle/modules/simulation/simulated_database/orders"
	mappableAssets "github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/genesis"
	mappableOrders "github.com/AssetMantle/modules/x/orders/mappable"
	"github.com/AssetMantle/modules/x/orders/parameters/max_order_life"
	"github.com/AssetMantle/modules/x/orders/record"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		max_order_life.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	records := make([]helpers.Record, len(assets.ClassificationIDMappableBytesMap))
	index := 0
	var classificationIDString string

	orders.ClearAll()
	for i := 0; i < len(assets.ClassificationIDMappableBytesMap); i++ {
		assetMap := assets.GetAssetData(simulationState.Accounts[i].Address.String())
		if assetMap == nil {
			continue
		}
		for class := range assetMap {
			classificationIDString = class
		}
		mappable := &mappableAssets.Mappable{}
		baseHelpers.CodecPrototype().MustUnmarshal(assets.ClassificationIDMappableBytesMap[classificationIDString], mappable)
		immutables := baseQualified.NewImmutables(mappable.Asset.Immutables.GetImmutablePropertyList().Add(baseLists.AnyPropertiesToProperties(constantProperties.ExchangeRateProperty.ToAnyProperty(),
			constantProperties.CreationHeightProperty.ToAnyProperty(),
			constantProperties.MakerAssetIDProperty.ToAnyProperty(),
			constantProperties.TakerAssetIDProperty.ToAnyProperty(),
			constantProperties.MakerIDProperty.ToAnyProperty(),
			constantProperties.TakerIDProperty.ToAnyProperty())...))
		mutables := baseQualified.NewMutables(mappable.Asset.Mutables.GetMutablePropertyList().Add(baseLists.AnyPropertiesToProperties(
			constantProperties.ExpiryHeightProperty.ToAnyProperty(),
			constantProperties.MakerSplitProperty.ToAnyProperty(),
		)...))
		classificationID := baseIDs.NewClassificationID(immutables, mutables)
		orderID := baseIDs.NewOrderID(classificationID, immutables)
		order := base.NewOrder(classificationID, immutables, mutables)
		records[index] = record.NewRecord(order)
		orders.AddOrderData(simulationState.Accounts[index].Address.String(), classificationID.AsString(), orderID.AsString())
		orders.AddMappableBytes(classificationID.AsString(), baseHelpers.CodecPrototype().MustMarshal(mappableOrders.NewMappable(order)))
		index++
	}

	genesisState := genesis.Prototype().Initialize(records, []parameters.Parameter{max_order_life.Parameter.Mutate(Data)})

	simulationState.GenState[constants.ModuleName] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
