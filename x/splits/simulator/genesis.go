// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/simulation/simulated_database/assets"
	"github.com/AssetMantle/modules/simulation/simulated_database/identities"
	"github.com/AssetMantle/modules/x/splits/constants"
	"github.com/AssetMantle/modules/x/splits/genesis"
	"github.com/AssetMantle/modules/x/splits/parameters/transfer_enabled"
	"github.com/AssetMantle/modules/x/splits/record"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		transfer_enabled.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	records := make([]helpers.Record, 2*len(assets.ClassificationIDMappableBytesMap))

	var assetIDString, identityIDString string
	index := 0
	for i := 0; i < len(assets.ClassificationIDMappableBytesMap); i++ {
		assetMap := assets.GetAssetData(simulationState.Accounts[i].Address.String())
		for _, id := range assetMap {
			assetIDString = id
		}
		assetID, _ := base.PrototypeAssetID().FromString(assetIDString)

		identityMap := identities.GetIDData(simulationState.Accounts[i].Address.String())

		for _, id := range identityMap {
			identityIDString = id
		}
		identityID, _ := base.PrototypeIdentityID().FromString(identityIDString)

		records[index] = record.NewRecord(base.NewSplitID(assetID.(ids.AssetID), identityID.(ids.IdentityID)), baseTypes.NewSplit(sdkTypes.NewInt(1)))
		records[index+1] = record.NewRecord(base.NewSplitID(baseDocuments.NewCoinAsset("stake").GetCoinAssetID(), identityID.(ids.IdentityID)), baseTypes.NewSplit(sdkTypes.NewInt(1000)))
		index += 2
	}

	genesisState := genesis.Prototype().Initialize(records, baseLists.NewParameterList(transfer_enabled.Parameter.Mutate(Data)))

	simulationState.GenState[constants.ModuleName] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
