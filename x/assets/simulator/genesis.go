// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/assets"
	"github.com/AssetMantle/modules/x/assets/mappable"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/genesis"
	assetsModule "github.com/AssetMantle/modules/x/assets/module"
	"github.com/AssetMantle/modules/x/assets/parameters/mintEnabled"
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		mintEnabled.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, len(simulationState.Accounts))

	assets.ClearAll()

	for i := range mappableList {
		immutables := baseQualified.NewImmutables(baseSimulation.GenerateRandomMetaPropertyListWithoutData(simulationState.Rand))
		mutables := baseQualified.NewMutables(baseSimulation.GenerateRandomPropertyList(simulationState.Rand).Add(baseProperties.NewMetaProperty(constants.SupplyProperty.GetKey(), baseData.NewNumberData(sdkTypes.NewInt(100)))))
		//for _, property := range immutables.GetImmutablePropertyList().GetList() {
		//	immutables = base.NewImmutables(immutables.GetImmutablePropertyList().Mutate(baseProperties.NewMetaProperty(property.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(simulationState.Rand, property.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))))
		//}
		classificationID := baseIDs.NewClassificationID(immutables, mutables)
		assetID := baseIDs.NewAssetID(classificationID, immutables)
		asset := baseDocuments.NewAsset(classificationID, immutables, mutables)
		mappableList[i] = mappable.NewMappable(asset)
		assets.AddAssetData(simulationState.Accounts[i].Address.String(), classificationID.AsString(), assetID.AsString())
		assets.AddMappableBytes(classificationID.AsString(), baseHelpers.CodecPrototype().MustMarshal(mappable.NewMappable(asset)))
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseParameters.NewParameterList(mintEnabled.Parameter.Mutate(Data)))

	simulationState.GenState[assetsModule.Name] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
