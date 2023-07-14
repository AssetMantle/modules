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
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/modules/x/splits/parameters/wrap_allowed_coins"
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

	mappableList := make([]helpers.Mappable, 2*len(assets.ClassificationIDMappableBytesMap))

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

		mappableList[index] = mappable.NewMappable(baseTypes.NewSplit(assetID.(ids.AssetID).ToAnyOwnableID().Get(), identityID.(ids.IdentityID), sdkTypes.NewInt(1)))
		mappableList[index+1] = mappable.NewMappable(baseTypes.NewSplit(base.NewCoinID(base.NewStringID("stake")), identityID.(ids.IdentityID), sdkTypes.NewInt(1000)))
		index += 2
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseLists.NewParameterList(wrap_allowed_coins.Parameter.Mutate(Data)))

	simulationState.GenState[constants.ModuleName] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
