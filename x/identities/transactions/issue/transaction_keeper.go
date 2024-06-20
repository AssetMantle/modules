// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"

	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/bond"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/constants"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/record"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
)

type transactionKeeper struct {
	mapper             helpers.Mapper
	parameterManager   helpers.ParameterManager
	conformAuxiliary   helpers.Auxiliary
	bondAuxiliary      helpers.Auxiliary
	authorizeAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if !transactionKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.IssueEnabledProperty.GetID()).GetMetaProperty().GetData().Get().(data.BooleanData).Get() {
		return nil, errorConstants.NotAuthorized.Wrapf("identity issuing is not enabled")
	}

	if _, err := transactionKeeper.authorizeAuxiliary.GetKeeper().Help(context, authorize.NewAuxiliaryRequest(message.ClassificationID, message.FromID, constants.CanIssueIdentityPermission)); err != nil {
		return nil, err
	}

	fromAddress, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if Mappable := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.FromID)).GetMappable(key.NewKey(message.FromID)); Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.FromID.AsString())
	} else if identity := mappable.GetIdentity(Mappable); !identity.IsProvisioned(fromAddress) {
		return nil, errorConstants.NotAuthorized.Wrapf("address %s is not provisioned for identity with ID %s", fromAddress.String(), message.FromID.AsString())
	}

	immutables := baseQualified.NewImmutables(message.ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.ImmutableProperties.Get()...)...))

	identityID := baseIDs.NewIdentityID(message.ClassificationID, immutables)

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(identityID))
	if identities.GetMappable(key.NewKey(identityID)) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("identity with ID %s already exists", identityID.AsString())
	}

	mutables := baseQualified.NewMutables(message.MutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.MutableProperties.Get()...)...))

	identity := base.NewIdentity(message.ClassificationID, immutables, mutables)

	if err := identity.ValidateBasic(); err != nil {
		return nil, err
	}

	if identity.GetProvisionedAddressCount().LT(types.OneInt()) {
		return nil, errorConstants.IncorrectFormat.Wrapf("identity with ID %s has no provisioned address", identityID.AsString())
	}

	if identity.GetProvisionedAddressCount().GTE(transactionKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.MaxProvisionAddressCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get()) {
		return nil, errorConstants.NotAuthorized.Wrapf("identity with ID %s has reached the maximum allowed number of provision-able addresses %d", identityID.AsString(), transactionKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.MaxProvisionAddressCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get())
	}

	if _, err := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); err != nil {
		return nil, err
	}

	bondAmount := types.ZeroInt()
	if bondAmountProperty := mutables.GetProperty(constantProperties.BondAmountProperty.GetID()); bondAmountProperty == nil || !bondAmountProperty.IsMeta() {
		return nil, errorConstants.MetaDataError.Wrapf("identity with ID %s has no revealed bond amount", identityID.AsString())
	} else {
		bondAmount = bondAmountProperty.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}

	if _, err := transactionKeeper.bondAuxiliary.GetKeeper().Help(context, bond.NewAuxiliaryRequest(message.ClassificationID, fromAddress, bondAmount)); err != nil {
		return nil, err
	}

	if err := identity.ValidateBasic(); err != nil {
		return nil, err
	}

	identities.Add(record.NewRecord(identity))

	return newTransactionResponse(identityID), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper
	transactionKeeper.parameterManager = parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case bond.Auxiliary.GetName():
				transactionKeeper.bondAuxiliary = value
			case authorize.Auxiliary.GetName():
				transactionKeeper.authorizeAuxiliary = value
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
