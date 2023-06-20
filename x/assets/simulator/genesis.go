// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/simulation/simulated_database/assets"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/genesis"
	"github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/assets/parameters/mint_enabled"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		mint_enabled.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, len(simulationState.Accounts))

	assets.ClearAll()

	for i := range mappableList {
		immutables := baseQualified.NewImmutables(baseSimulation.GenerateRandomMetaPropertyListWithoutData(simulationState.Rand))
		mutables := baseQualified.NewMutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand).Add(baseProperties.NewMetaProperty(constantProperties.SupplyProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(100)))))
		// for _, property := range immutables.GetImmutablePropertyList().GetList() {
		//	immutables = base.NewImmutables(immutables.GetImmutablePropertyList().Mutate(baseProperties.NewMetaProperty(property.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(simulationState.Rand, property.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))))
		// }
		classificationID := baseIDs.NewClassificationID(immutables, mutables)
		assetID := baseIDs.NewAssetID(classificationID, immutables)
		asset := baseDocuments.NewAsset(classificationID, immutables, mutables)
		mappableList[i] = mappable.NewMappable(asset)
		assets.AddAssetData(simulationState.Accounts[i].Address.String(), classificationID.AsString(), assetID.AsString())
		assets.AddMappableBytes(classificationID.AsString(), baseHelpers.CodecPrototype().MustMarshal(mappable.NewMappable(asset)))
	}

	genesisState := genesis.Prototype().Initialize(mappableList, base.NewParameterList(mint_enabled.Parameter.Mutate(Data)))

	simulationState.GenState[constants.ModuleName] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
