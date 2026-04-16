// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	baseData "github.com/AssetMantle/schema/data/base"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"

	"github.com/AssetMantle/modules/helpers"
	simulationModules "github.com/AssetMantle/modules/simulation"
	baseTypes "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/simulation/simulated_database/identities"
	"github.com/AssetMantle/modules/x/assets/transactions/revoke"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/transactions/define"
	"github.com/AssetMantle/modules/x/identities/transactions/deputize"
	"github.com/AssetMantle/modules/x/identities/transactions/issue"
	"github.com/AssetMantle/modules/x/identities/transactions/name"
	"github.com/AssetMantle/modules/x/identities/transactions/provision"
	"github.com/AssetMantle/modules/x/identities/transactions/quash"
	"github.com/AssetMantle/modules/x/identities/transactions/unprovision"
	"github.com/AssetMantle/modules/x/identities/transactions/update"
)

func (simulator) WeightedOperations(simulationState module.SimulationState, module helpers.Module) simulation.WeightedOperations {
	var weightMsg int

	simulationState.AppParams.GetOrGenerate(OpWeightMsg, &weightMsg, nil,
		func(_ *rand.Rand) {
			weightMsg = DefaultWeightMsg
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateNameMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateDefineMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateIssueMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateProvisionAndUnprovisionMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateDeputizeAndRevokeMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateQuashMsg(module)),
		),
		simulation.NewWeightedOperation(
			weightMsg,
			simulationModules.SafeOperation(simulateMutateMsg(module)),
		),
	}
}

func simulateNameMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		account, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		message := GenerateNameMessage(account.Address, baseTypes.GenerateRandomID(rand))
		result, err := simulationModules.ExecuteMessage(context, module, message)
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data)), nil, nil
	}
}
func simulateDefineMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		var message *define.Message
		account, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		identityIDString, err := simulationModules.LookupIdentityID(account.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "define", "no identity data"), nil, nil
		}
		identityID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)
		message = GenerateDefineMessage(account.Address, identityID.(ids.IdentityID), rand).(*define.Message)
		result, err = simulationModules.ExecuteMessage(context, module, message)
		if err != nil {
			return simulationTypes.NewOperationMsg(message, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(message, true, string(result.Data)), nil, nil
	}
}
func simulateIssueMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "issue", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		// IMPORTANT: PropertyList.Add mutates in place. Snapshot all random
		// properties as slices first and build INDEPENDENT PropertyLists for
		// each step to prevent shared-state corruption.
		immMetaSnap := baseTypes.GenerateRandomMetaPropertyList(rand).Get()
		immSnap := baseTypes.GenerateRandomPropertyList(rand).Get()
		mutMetaSnap := baseTypes.GenerateRandomMetaPropertyList(rand).Get()
		mutSnap := baseTypes.GenerateRandomPropertyList(rand).Get()

		// Step 1: Define classification with BondAmountProperty in mutable meta.
		defineMsg := define.NewMessage(from.Address, fromID.(ids.IdentityID),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immSnap...)...),
			baseLists.NewPropertyList(constants.BondAmountProperty).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(mutSnap...)...))
		_, err = simulationModules.ExecuteMessage(context, module, defineMsg.(helpers.Message))
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "issue", "define failed: "+err.Error()), nil, nil
		}

		// Compute classificationID (fresh lists, no mutation from define).
		immutables := baseQualified.NewImmutables(
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...).Add(baseLists.AnyPropertiesToProperties(immSnap...)...))
		classifMuts := baseLists.NewPropertyList(constants.BondAmountProperty, constants.AuthenticationProperty).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...).Add(baseLists.AnyPropertiesToProperties(mutSnap...)...)
		classificationID := baseIDs.NewClassificationID(immutables, baseQualified.NewMutables(classifMuts))

		// Step 2: Issue (fresh lists, no shared state with define/classifID).
		authPropWithAddr := baseProperties.NewMetaProperty(
			constants.AuthenticationProperty.GetKey(),
			baseData.NewListData(baseData.NewAccAddressData(from.Address)),
		)
		issueMsg := issue.NewMessage(from.Address, fromID.(ids.IdentityID), classificationID,
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immSnap...)...),
			baseLists.NewPropertyList(constants.BondAmountProperty, authPropWithAddr).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...),
			baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(mutSnap...)...))
		result, err := simulationModules.ExecuteMessage(context, module, issueMsg.(helpers.Message))
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "issue", "issue failed: "+err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(issueMsg, true, string(result.Data)), nil, nil
	}
}
func simulateProvisionAndUnprovisionMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "provision", "no identity data"), nil, nil
		}
		identityID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		provisionMessage := provision.NewMessage(from.Address, to.Address, identityID.(ids.IdentityID))
		result, err = simulationModules.ExecuteMessage(context, module, provisionMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(provisionMessage, false, err.Error()), nil, nil
		}
		unprovisionMessage := unprovision.NewMessage(from.Address, to.Address, identityID.(ids.IdentityID))
		result, err = simulationModules.ExecuteMessage(context, module, unprovisionMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(unprovisionMessage, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(unprovisionMessage, true, string(result.Data)), nil, nil
	}
}
func simulateDeputizeAndRevokeMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		var err error
		var result *sdkTypes.Result
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)
		to, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "deputize", "no identity data"), nil, nil
		}

		// For identities module, classification comes from the identity map (key = classification)
		fromIDMap := identities.GetIDData(from.Address.String())
		var classificationIDString string
		for class := range fromIDMap {
			classificationIDString = class
			break
		}

		classificationID, _ := baseIDs.PrototypeClassificationID().FromString(classificationIDString)
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		toIdentityIDString, err := simulationModules.LookupIdentityID(to.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "deputize", "no identity data"), nil, nil
		}

		toID, _ := baseIDs.PrototypeIdentityID().FromString(toIdentityIDString)
		Mappable := &mappable.Mappable{}
		baseHelpers.CodecPrototype().Unmarshal(identities.GetMappableBytes(classificationIDString), Mappable)
		deputizeMessage := deputize.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), classificationID.(ids.ClassificationID), Mappable.Identity.Mutables.PropertyList, true, true, true, true, true)
		result, err = simulationModules.ExecuteMessage(context, module, deputizeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(deputizeMessage, false, err.Error()), nil, nil
		}
		revokeMessage := revoke.NewMessage(from.Address, fromID.(ids.IdentityID), toID.(ids.IdentityID), classificationID.(ids.ClassificationID))
		result, err = simulationModules.ExecuteMessage(context, module, revokeMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NewOperationMsg(revokeMessage, false, err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(revokeMessage, true, string(result.Data)), nil, nil
	}
}
func simulateQuashMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "quash", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		issueMsg, issuedID := DefineAndIssue(context, module, from, rand)
		if issueMsg == nil {
			return simulationTypes.NoOpMsg("identities", "quash", "define+issue failed"), nil, nil
		}

		quashMessage := quash.NewMessage(from.Address, fromID.(ids.IdentityID), issuedID)
		result, err := simulationModules.ExecuteMessage(context, module, quashMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "quash", err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(quashMessage, true, string(result.Data)), nil, nil
	}
}
func simulateMutateMsg(module helpers.Module) simulationTypes.Operation {
	return func(rand *rand.Rand, baseApp *baseapp.BaseApp, context sdkTypes.Context, simulationAccountList []simulationTypes.Account, chainID string) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
		from, _ := simulationTypes.RandomAcc(rand, simulationAccountList)

		identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "mutate", "no identity data"), nil, nil
		}
		fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

		_, issuedID := DefineAndIssue(context, module, from, rand)
		if issuedID == nil {
			return simulationTypes.NoOpMsg("identities", "mutate", "define+issue failed"), nil, nil
		}

		mutateMessage := update.NewMessage(from.Address, fromID.(ids.IdentityID), issuedID, baseLists.NewPropertyList(), baseLists.NewPropertyList())
		result, err := simulationModules.ExecuteMessage(context, module, mutateMessage.(helpers.Message))
		if err != nil {
			return simulationTypes.NoOpMsg("identities", "mutate", err.Error()), nil, nil
		}
		return simulationTypes.NewOperationMsg(mutateMessage, true, string(result.Data)), nil, nil
	}
}

func GenerateNameMessage(from sdkTypes.AccAddress, Name ids.StringID) helpers.Message {
	return name.NewMessage(from, Name).(helpers.Message)
}
func GenerateDefineMessage(from sdkTypes.AccAddress, identityID ids.IdentityID, r *rand.Rand) helpers.Message {
	return define.NewMessage(from, identityID, baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r), baseTypes.GenerateRandomMetaPropertyList(r), baseTypes.GenerateRandomPropertyList(r)).(helpers.Message)
}
// DefineAndIssue defines a new identity classification and issues an identity under it.
// Uses PropertyList snapshots to prevent in-place mutation corruption.
// Returns the issue message and identity ID, or nil on failure.
func DefineAndIssue(context sdkTypes.Context, module helpers.Module, from simulationTypes.Account, rand *rand.Rand) (sdkTypes.Msg, ids.IdentityID) {
	identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
	if err != nil {
		return nil, nil
	}
	fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

	immMetaSnap := baseTypes.GenerateRandomMetaPropertyList(rand).Get()
	immSnap := baseTypes.GenerateRandomPropertyList(rand).Get()
	mutMetaSnap := baseTypes.GenerateRandomMetaPropertyList(rand).Get()
	mutSnap := baseTypes.GenerateRandomPropertyList(rand).Get()

	defineMsg := define.NewMessage(from.Address, fromID.(ids.IdentityID),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immSnap...)...),
		baseLists.NewPropertyList(constants.BondAmountProperty).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(mutSnap...)...))
	_, err = simulationModules.ExecuteMessage(context, module, defineMsg.(helpers.Message))
	if err != nil {
		return nil, nil
	}

	immutables := baseQualified.NewImmutables(
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...).Add(baseLists.AnyPropertiesToProperties(immSnap...)...))
	classifMuts := baseLists.NewPropertyList(constants.BondAmountProperty, constants.AuthenticationProperty).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...).Add(baseLists.AnyPropertiesToProperties(mutSnap...)...)
	classificationID := baseIDs.NewClassificationID(immutables, baseQualified.NewMutables(classifMuts))

	authPropWithAddr := baseProperties.NewMetaProperty(
		constants.AuthenticationProperty.GetKey(),
		baseData.NewListData(baseData.NewAccAddressData(from.Address)),
	)
	issueMsg := issue.NewMessage(from.Address, fromID.(ids.IdentityID), classificationID,
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immMetaSnap...)...),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(immSnap...)...),
		baseLists.NewPropertyList(constants.BondAmountProperty, authPropWithAddr).Add(baseLists.AnyPropertiesToProperties(mutMetaSnap...)...),
		baseLists.NewPropertyList(baseLists.AnyPropertiesToProperties(mutSnap...)...))
	_, err = simulationModules.ExecuteMessage(context, module, issueMsg.(helpers.Message))
	if err != nil {
		return nil, nil
	}

	identityID := baseIDs.NewIdentityID(classificationID, immutables)
	return issueMsg, identityID
}

func GetIssueMessage(from, to simulationTypes.Account, rand *rand.Rand) sdkTypes.Msg {
	identityIDString, err := simulationModules.LookupIdentityID(from.Address.String())
	if err != nil {
		return nil
	}
	fromID, _ := baseIDs.PrototypeIdentityID().FromString(identityIDString)

	// For identities module, classification comes from the identity map (key = classification)
	identityMap := identities.GetIDData(from.Address.String())
	var classificationIDString string
	for class := range identityMap {
		classificationIDString = class
		break
	}
	classificationID, _ := baseIDs.PrototypeClassificationID().FromString(classificationIDString)
	mappable := &mappable.Mappable{}
	baseHelpers.CodecPrototype().Unmarshal(identities.GetMappableBytes(classificationIDString), mappable)
	immutableMetaProperties := &baseLists.PropertyList{}
	immutableProperties := &baseLists.PropertyList{}
	mutableMetaProperties := &baseLists.PropertyList{}
	mutableProperties := &baseLists.PropertyList{}
	if mappable.Identity == nil {
		return nil
	}
	// Regenerate immutable property values with same keys but new random data to
	// produce unique identity IDs while conforming to the classification's structure.
	for _, i := range mappable.GetIdentity().Get().GetImmutables().GetImmutablePropertyList().Get() {
		if i.IsMeta() {
			immutableMetaProperties = immutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
		} else {
			immutableProperties = immutableProperties.Add(baseProperties.NewMesaProperty(i.Get().GetKey(), baseTypes.GenerateRandomData(rand, rand.Intn(8)))).(*baseLists.PropertyList)
		}
	}
	for _, i := range mappable.GetIdentity().Get().GetMutables().GetMutablePropertyList().Get() {
		if i.IsMeta() {
			mutableMetaProperties = mutableMetaProperties.Add(baseProperties.NewMetaProperty(i.Get().GetKey(), baseTypes.GenerateRandomDataForTypeID(rand, i.Get().(*baseProperties.MetaProperty).GetData().GetTypeID()))).(*baseLists.PropertyList)
		} else {
			mutableProperties = mutableProperties.Add(i).(*baseLists.PropertyList)
		}
	}
	return issue.NewMessage(from.Address, fromID.(ids.IdentityID), classificationID.(ids.ClassificationID), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
}
