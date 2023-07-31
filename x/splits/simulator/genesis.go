// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/simulation/simulated_database/assets"
	"github.com/AssetMantle/modules/simulation/simulated_database/identities"
	"github.com/AssetMantle/modules/x/splits/constants"
	"github.com/AssetMantle/modules/x/splits/genesis"
	"github.com/AssetMantle/modules/x/splits/parameters/wrap_allowed_coins"
	"github.com/AssetMantle/modules/x/splits/record"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		wrap_allowed_coins.ID.AsString(),
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

		records[index] = record.NewRecord(base.NewSplitID(assetID.(ids.AssetID).ToAnyOwnableID().Get(), identityID.(ids.IdentityID)), baseTypes.NewSplit(sdkTypes.NewInt(1)))
		records[index+1] = record.NewRecord(base.NewSplitID(base.NewCoinID(base.NewStringID("stake")), identityID.(ids.IdentityID)), baseTypes.NewSplit(sdkTypes.NewInt(1000)))
		index += 2
	}

	genesisState := genesis.Prototype().Initialize(records, baseLists.NewParameterList(wrap_allowed_coins.Parameter.Mutate(Data)))

	simulationState.GenState[constants.ModuleName] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
