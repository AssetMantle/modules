// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/simulation/simulated_database/assets"
	mappableAssets "github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/classifications/constants"
	"github.com/AssetMantle/modules/x/classifications/genesis"
	"github.com/AssetMantle/modules/x/classifications/parameters/bond_rate"
	"github.com/AssetMantle/modules/x/classifications/parameters/max_property_count"
	"github.com/AssetMantle/modules/x/classifications/record"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var bondRateData data.Data
	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		bond_rate.ID.AsString(),
		&bondRateData,
		simulationState.Rand,
		func(rand *rand.Rand) {
			bondRateData = baseData.NewNumberData(sdkTypes.NewInt(int64(rand.Intn(99))))
		},
	)

	var maxPropertyCountData data.Data
	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		max_property_count.ID.AsString(),
		&maxPropertyCountData,
		simulationState.Rand,
		func(rand *rand.Rand) {
			maxPropertyCountData = baseData.NewNumberData(sdkTypes.NewInt(int64(rand.Intn(99))))
		},
	)

	records := make([]helpers.Record, 3*len(assets.ClassificationIDMappableBytesMap))
	index := 0
	accountPosition := 0

	for i := range assets.ClassificationIDMappableBytesMap {
		mappable := &mappableAssets.Mappable{}
		baseHelpers.CodecPrototype().MustUnmarshal(assets.ClassificationIDMappableBytesMap[i], mappable)

		immutables := mappable.Asset.Immutables
		mutables := mappable.Asset.Mutables

		assetClassification := base.NewClassification(immutables, mutables)
		identityClassification := base.NewClassification(immutables, baseQualified.NewMutables(mutables.GetMutablePropertyList().Add(constantProperties.AuthenticationProperty)))
		orderClassification := base.NewClassification(baseQualified.NewImmutables(immutables.GetImmutablePropertyList().Add(baseLists.AnyPropertiesToProperties(constantProperties.ExchangeRateProperty.ToAnyProperty(),
			constantProperties.CreationHeightProperty.ToAnyProperty(),
			constantProperties.MakerOwnableIDProperty.ToAnyProperty(),
			constantProperties.TakerOwnableIDProperty.ToAnyProperty(),
			constantProperties.MakerIDProperty.ToAnyProperty(),
			constantProperties.TakerIDProperty.ToAnyProperty())...)), baseQualified.NewMutables(mappable.Asset.Mutables.GetMutablePropertyList().Add(baseLists.AnyPropertiesToProperties(
			constantProperties.ExpiryHeightProperty.ToAnyProperty(),
			constantProperties.MakerOwnableSplitProperty.ToAnyProperty(),
		)...)))

		records[index] = record.NewRecord(assetClassification)
		records[index+1] = record.NewRecord(identityClassification)
		records[index+2] = record.NewRecord(orderClassification)

		index += 3
		accountPosition++
	}

	genesisState := genesis.Prototype().Initialize(records, baseLists.NewParameterList(bond_rate.Parameter.Mutate(bondRateData), max_property_count.Parameter.Mutate(maxPropertyCountData)))

	simulationState.GenState[constants.ModuleName] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
