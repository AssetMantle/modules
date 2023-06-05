// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"math/rand"

	baseLists "github.com/AssetMantle/schema/go/lists/base"

	simulatorAssets "github.com/AssetMantle/modules/simulation/simulatedDatabase/assets"
	simulatorIdentities "github.com/AssetMantle/modules/simulation/simulatedDatabase/identities"

	"github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"

	mappableAssets "github.com/AssetMantle/modules/x/assets/mappable"
	mappableIdentities "github.com/AssetMantle/modules/x/identities/mappable"
	mappableMaintainers "github.com/AssetMantle/modules/x/maintainers/mappable"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/maintainers/genesis"
	maintainersModule "github.com/AssetMantle/modules/x/maintainers/module"
	"github.com/AssetMantle/modules/x/maintainers/parameters/deputizeAllowed"
	"github.com/AssetMantle/modules/x/maintainers/utilities"
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

	mappableList := make([]helpers.Mappable, 3*len(simulatorAssets.ClassificationIDMappableBytesMap))

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
			break
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

		immutables := baseQualified.NewImmutables(assetMappable.Asset.Immutables.GetImmutablePropertyList().Add(baseLists.AnyPropertiesToProperties(constants.ExchangeRateProperty.ToAnyProperty(),
			constants.CreationHeightProperty.ToAnyProperty(),
			constants.MakerOwnableIDProperty.ToAnyProperty(),
			constants.TakerOwnableIDProperty.ToAnyProperty(),
			constants.MakerIDProperty.ToAnyProperty(),
			constants.TakerIDProperty.ToAnyProperty())...))

		mutables = baseQualified.NewMutables(assetMappable.Asset.Mutables.GetMutablePropertyList().Add(baseLists.AnyPropertiesToProperties(
			constants.ExpiryHeightProperty.ToAnyProperty(),
			constants.MakerOwnableSplitProperty.ToAnyProperty(),
		)...))

		orderClassificationID := baseIDs.NewClassificationID(immutables, mutables)
		x := orderClassificationID.AsString()
		fmt.Println(x)
		mappableList[index+2] = mappableMaintainers.NewMappable(base.NewMaintainer(identityID, orderClassificationID, mutables.GetMutablePropertyList().GetPropertyIDList(), utilities.SetPermissions(true, true, true, true, true, true)))

		index += 3
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseLists.NewParameterList(deputizeAllowed.Parameter.Mutate(Data)))

	simulationState.GenState[maintainersModule.Name] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
