// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/assets"
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/classifications"
	mappableAssets "github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/classifications/parameters/maxPropertyCount"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/genesis"
	mappableClassifications "github.com/AssetMantle/modules/x/classifications/mappable"
	classificationsModule "github.com/AssetMantle/modules/x/classifications/module"
	"github.com/AssetMantle/modules/x/classifications/parameters/bondRate"
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var bondRateData data.Data
	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		bondRate.ID.AsString(),
		&bondRateData,
		simulationState.Rand,
		func(rand *rand.Rand) {
			bondRateData = baseData.NewNumberData(sdkTypes.NewInt(int64(rand.Intn(99))))
		},
	)

	var maxPropertyCountData data.Data
	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		maxPropertyCount.ID.AsString(),
		&maxPropertyCountData,
		simulationState.Rand,
		func(rand *rand.Rand) {
			maxPropertyCountData = baseData.NewNumberData(sdkTypes.NewInt(int64(rand.Intn(99))))
		},
	)

	mappableList := make([]helpers.Mappable, 2*len(assets.ClassificationIDMappableBytesMap))
	index := 0
	accountPosition := 0
	classifications.ClearAll()

	for i := range assets.ClassificationIDMappableBytesMap {
		mappable := &mappableAssets.Mappable{}
		baseHelpers.CodecPrototype().MustUnmarshal(assets.ClassificationIDMappableBytesMap[i], mappable)
		immutables := mappable.Asset.Immutables
		mutables := mappable.Asset.Mutables
		assetClassification := base.NewClassification(immutables, mutables)
		identityClassification := base.NewClassification(immutables, baseQualified.NewMutables(mutables.GetMutablePropertyList().Add(constantProperties.AuthenticationProperty)))
		mappableList[index] = mappableClassifications.NewMappable(assetClassification)
		mappableList[index+1] = mappableClassifications.NewMappable(identityClassification)
		index += 2

		classifications.AddAssetClassificationID(simulationState.Accounts[accountPosition].Address.String(), assetClassification.GetClassificationID().AsString())
		classifications.AddIdentityClassificationID(simulationState.Accounts[accountPosition].Address.String(), identityClassification.GetClassificationID().AsString())

		classifications.AddAssetMappableBytes(assetClassification.GetClassificationID().AsString(), baseHelpers.CodecPrototype().MustMarshal(mappableClassifications.NewMappable(assetClassification)))
		classifications.AddIdentityMappableBytes(identityClassification.GetClassificationID().AsString(), baseHelpers.CodecPrototype().MustMarshal(mappableClassifications.NewMappable(identityClassification)))
		accountPosition++
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseParameters.NewParameterList(bondRate.Parameter.Mutate(bondRateData), maxPropertyCount.Parameter.Mutate(maxPropertyCountData)))

	simulationState.GenState[classificationsModule.Name] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
