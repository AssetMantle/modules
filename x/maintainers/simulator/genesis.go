// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	simulatorAssets "github.com/AssetMantle/modules/simulation/simulated_database/assets"
	simulatorIdentities "github.com/AssetMantle/modules/simulation/simulated_database/identities"
	mappableAssets "github.com/AssetMantle/modules/x/assets/mappable"
	assetUtilities "github.com/AssetMantle/modules/x/assets/utilities"
	mappableIdentities "github.com/AssetMantle/modules/x/identities/mappable"
	identityUtilities "github.com/AssetMantle/modules/x/identities/utilities"
	"github.com/AssetMantle/modules/x/maintainers/constants"
	"github.com/AssetMantle/modules/x/maintainers/genesis"
	mappableMaintainers "github.com/AssetMantle/modules/x/maintainers/mappable"
	"github.com/AssetMantle/modules/x/maintainers/parameters/deputize_allowed"
	"github.com/AssetMantle/modules/x/maintainers/utilities"
	orderUtilities "github.com/AssetMantle/modules/x/orders/utilities"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		deputize_allowed.ID.AsString(),
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
		mappableList[index] = mappableMaintainers.NewMappable(base.NewMaintainer(identityID, classificationID, mutables.GetMutablePropertyList().GetPropertyIDList(), utilities.SetModulePermissions(true, true, true).Add(baseIDs.StringIDsToIDs(identityUtilities.SetModulePermissions(true, true))...)))

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
		mappableList[index+1] = mappableMaintainers.NewMappable(base.NewMaintainer(identityID, classificationID, mutables.GetMutablePropertyList().GetPropertyIDList(), utilities.SetModulePermissions(true, true, true).Add(baseIDs.StringIDsToIDs(assetUtilities.SetModulePermissions(true, true, true))...)))

		immutables := baseQualified.NewImmutables(assetMappable.Asset.Immutables.GetImmutablePropertyList().Add(baseLists.AnyPropertiesToProperties(constantProperties.ExchangeRateProperty.ToAnyProperty(),
			constantProperties.CreationHeightProperty.ToAnyProperty(),
			constantProperties.MakerOwnableIDProperty.ToAnyProperty(),
			constantProperties.TakerOwnableIDProperty.ToAnyProperty(),
			constantProperties.MakerIDProperty.ToAnyProperty(),
			constantProperties.TakerIDProperty.ToAnyProperty())...))

		mutables = baseQualified.NewMutables(assetMappable.Asset.Mutables.GetMutablePropertyList().Add(baseLists.AnyPropertiesToProperties(
			constantProperties.ExpiryHeightProperty.ToAnyProperty(),
			constantProperties.MakerOwnableSplitProperty.ToAnyProperty(),
		)...))

		orderClassificationID := baseIDs.NewClassificationID(immutables, mutables)
		mappableList[index+2] = mappableMaintainers.NewMappable(base.NewMaintainer(identityID, orderClassificationID, mutables.GetMutablePropertyList().GetPropertyIDList(), utilities.SetModulePermissions(true, true, true).Add(baseIDs.StringIDsToIDs(orderUtilities.SetModulePermissions(true, true))...)))

		index += 3
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseLists.NewParameterList(deputize_allowed.Parameter.Mutate(Data)))

	simulationState.GenState[constants.ModuleName] = baseHelpers.CodecPrototype().MustMarshalJSON(genesisState)
}
