// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	simulatorIdentities "github.com/AssetMantle/modules/simulation/simulatedDatabase/identities"
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

	mappableList := make([]helpers.Mappable, len(simulationState.Accounts))

	var classificationID ids.ClassificationID
	var identityID ids.IdentityID

	for i := range mappableList {
		identityMap := simulatorIdentities.GetIDData(simulationState.Accounts[i].Address.String())
		for class, id := range identityMap {
			classificationID, _ = baseIDs.ReadClassificationID(class)
			identityID, _ = baseIDs.ReadIdentityID(id)
		}
		mappable := &mappableIdentities.Mappable{}
		baseHelpers.CodecPrototype().Unmarshal(simulatorIdentities.GetMappableBytes(classificationID.AsString()), mappable)
		mutables := mappable.GetIdentity().Get().GetMutables()
		mappableList[i] = mappableMaintainers.NewMappable(base.NewMaintainer(identityID, classificationID, mutables.GetMutablePropertyList().GetPropertyIDList(), utilities.SetPermissions(true, true, true, true, true, true)))
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseParameters.NewParameterList(deputizeAllowed.Parameter.Mutate(Data)))

	simulationState.GenState[maintainersModule.Name] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
