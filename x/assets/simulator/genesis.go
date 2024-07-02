// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/simulation/simulated_database/assets"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/genesis"
	"github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/assets/parameters/burn_enabled"
	"github.com/AssetMantle/modules/x/assets/parameters/mint_enabled"
	"github.com/AssetMantle/modules/x/assets/record"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		burn_enabled.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewBooleanData(random.GenerateRandomBool()) },
	)

	records := make([]helpers.Record, len(simulationState.Accounts))

	assets.ClearAll()

	for i := range records {
		immutables := baseQualified.NewImmutables(baseSimulation.GenerateRandomMetaPropertyListWithoutData(simulationState.Rand))
		mutables := baseQualified.NewMutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand).Add(baseProperties.NewMetaProperty(constantProperties.SupplyProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(100)))))
		// for _, property := range immutables.GetImmutablePropertyList().GetList() {
		//	immutables = base.NewImmutables(immutables.GetImmutablePropertyList().Mutate(baseProperties.NewMetaProperty(property.Get().GenerateKey(), baseTypes.GenerateRandomDataForTypeID(simulationState.Rand, property.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))))
		// }
		classificationID := baseIDs.NewClassificationID(immutables, mutables)
		assetID := baseIDs.NewAssetID(classificationID, immutables)
		asset := baseDocuments.NewAsset(classificationID, immutables, mutables)
		records[i] = record.NewRecord(asset)
		assets.AddAssetData(simulationState.Accounts[i].Address.String(), classificationID.AsString(), assetID.AsString())
		assets.AddMappableBytes(classificationID.AsString(), baseHelpers.CodecPrototype().MustMarshal(mappable.NewMappable(asset)))
	}

	genesisState := genesis.Prototype().Initialize(records, base.NewParameterList(mint_enabled.Parameter.Mutate(Data)))

	simulationState.GenState[constants.ModuleName] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
