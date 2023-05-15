// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	simulatorAssets "github.com/AssetMantle/modules/simulation/simulatedDatabase/assets"
	simulatorIdentities "github.com/AssetMantle/modules/simulation/simulatedDatabase/identities"
	mappableAssets "github.com/AssetMantle/modules/x/assets/mappable"
	mappableIdentities "github.com/AssetMantle/modules/x/identities/mappable"
	mappableMaintainers "github.com/AssetMantle/modules/x/maintainers/mappable"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/maintainers/genesis"
	maintainersModule "github.com/AssetMantle/modules/x/maintainers/module"
	"github.com/AssetMantle/modules/x/maintainers/parameters/deputizeAllowed"
	"github.com/AssetMantle/modules/x/maintainers/utilities"
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		deputizeAllowed.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, 2*len(simulatorAssets.ClassificationIDMappableBytesMap))

	var classificationID ids.ClassificationID
	var identityID ids.IdentityID
	index := 0
	for i := 0; i < len(simulatorAssets.ClassificationIDMappableBytesMap); i++ {
		identityMap := simulatorIdentities.GetIDData(simulationState.Accounts[i].Address.String())
		if identityMap == nil {
			continue
		}
		for class, id := range identityMap {
			classificationID, _ = baseIDs.ReadClassificationID(class)
			identityID, _ = baseIDs.ReadIdentityID(id)
		}
		identityMappable := &mappableIdentities.Mappable{}
		baseHelpers.CodecPrototype().Unmarshal(simulatorIdentities.GetMappableBytes(classificationID.AsString()), identityMappable)
		mutables := identityMappable.GetIdentity().Get().GetMutables()
		mappableList[index] = mappableMaintainers.NewMappable(base.NewMaintainer(identityID, classificationID, mutables.GetMutablePropertyList().GetPropertyIDList(), utilities.SetPermissions(true, true, true, true, true, true)))

		assetMap := simulatorAssets.GetAssetData(simulationState.Accounts[i].Address.String())
		if assetMap == nil {
			continue
		}
		for class, _ := range assetMap {
			classificationID, _ = baseIDs.ReadClassificationID(class)
		}
		assetMappable := &mappableAssets.Mappable{}
		baseHelpers.CodecPrototype().Unmarshal(simulatorAssets.GetMappableBytes(classificationID.AsString()), assetMappable)
		mutables = assetMappable.GetAsset().Get().GetMutables()
		mappableList[index+1] = mappableMaintainers.NewMappable(base.NewMaintainer(identityID, classificationID, mutables.GetMutablePropertyList().GetPropertyIDList(), utilities.SetPermissions(true, true, true, true, true, true)))
		index += 2
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseParameters.NewParameterList(deputizeAllowed.Parameter.Mutate(Data)))

	simulationState.GenState[maintainersModule.Name] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
