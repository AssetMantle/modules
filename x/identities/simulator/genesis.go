// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/assets"
	"github.com/AssetMantle/modules/simulation/simulatedDatabase/identities"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	mappableAssets "github.com/AssetMantle/modules/x/assets/mappable"
	"github.com/AssetMantle/modules/x/identities/genesis"
	mappableIdentities "github.com/AssetMantle/modules/x/identities/mappable"
	identitiesModule "github.com/AssetMantle/modules/x/identities/module"
	"github.com/AssetMantle/modules/x/identities/parameters/maxProvisionAddressCount"
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		maxProvisionAddressCount.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, len(assets.ClassificationIDMappableBytesMap))

	identities.ClearAll()
	index := 0
	var classificationIDString string

	for i := 0; i < len(assets.ClassificationIDMappableBytesMap); i++ {
		assetMap := assets.GetAssetData(simulationState.Accounts[i].Address.String())
		if assetMap == nil {
			continue
		}
		for class, _ := range assetMap {
			classificationIDString = class
		}
		mappable := &mappableAssets.Mappable{}
		baseHelpers.CodecPrototype().MustUnmarshal(assets.ClassificationIDMappableBytesMap[classificationIDString], mappable)
		immutables := mappable.Asset.Immutables
		mutables := baseQualified.NewMutables(mappable.Asset.Mutables.GetMutablePropertyList().Add(constantProperties.AuthenticationProperty))
		classificationID := baseIDs.NewClassificationID(immutables, mutables)
		identityID := baseIDs.NewIdentityID(classificationID, immutables)
		identity := base.NewIdentity(classificationID, immutables, mutables).ProvisionAddress(simulationState.Accounts[index].Address)

		mappableList[index] = mappableIdentities.NewMappable(identity)
		identities.AddIDData(simulationState.Accounts[index].Address.String(), classificationID.AsString(), identityID.AsString())
		identities.AddMappableBytes(classificationID.AsString(), baseHelpers.CodecPrototype().MustMarshal(mappableIdentities.NewMappable(identity)))
		index++
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseParameters.NewParameterList(maxProvisionAddressCount.Parameter.Mutate(Data)))

	simulationState.GenState[identitiesModule.Name] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
