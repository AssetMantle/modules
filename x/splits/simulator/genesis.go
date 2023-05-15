// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/assets"
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/identities"
	"github.com/AssetMantle/modules/x/splits/mappable"
	"github.com/AssetMantle/schema/go/ids/base"
	base2 "github.com/AssetMantle/schema/go/types/base"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/genesis"
	splitsModule "github.com/AssetMantle/modules/x/splits/module"
	"github.com/AssetMantle/modules/x/splits/parameters/wrapAllowedCoins"
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		wrapAllowedCoins.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, len(assets.ClassificationIDMappableBytesMap))

	var assetIDString, identityIDString string
	for i := 0; i < len(assets.ClassificationIDMappableBytesMap); i++ {
		assetMap := assets.GetAssetData(simulationState.Accounts[i].Address.String())
		for _, id := range assetMap {
			assetIDString = id
		}
		assetID, _ := base.ReadAssetID(assetIDString)

		identityMap := identities.GetIDData(simulationState.Accounts[i].Address.String())

		for _, id := range identityMap {
			identityIDString = id
		}
		identityID, _ := base.ReadIdentityID(identityIDString)

		mappableList[i] = mappable.NewMappable(base2.NewSplit(identityID, assetID.ToAnyOwnableID().Get(), sdkTypes.NewInt(1)))
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseParameters.NewParameterList(wrapAllowedCoins.Parameter.Mutate(Data)))

	simulationState.GenState[splitsModule.Name] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
