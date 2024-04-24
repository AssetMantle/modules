// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package update

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/record"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/schema/go/documents/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
)

type transactionKeeper struct {
	mapper            helpers.Mapper
	maintainAuxiliary helpers.Auxiliary
	memberAuxiliary   helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	fromAddress := message.GetFromAddress()

	if Mappable := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.FromID)).GetMappable(key.NewKey(message.FromID)); Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.FromID.AsString())
	} else if identity := mappable.GetIdentity(Mappable); !identity.IsProvisioned(fromAddress) {
		return nil, errorConstants.NotAuthorized.Wrapf("address %s is not provisioned for identity with ID %s", fromAddress.String(), message.FromID.AsString())
	}

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.IdentityID))

	Mappable := identities.GetMappable(key.NewKey(message.IdentityID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.IdentityID.AsString())
	}
	identity := mappable.GetIdentity(Mappable)

	mutables := baseQualified.NewMutables(message.MutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.MutableProperties.Get()...)...))

	if _, err := transactionKeeper.memberAuxiliary.GetKeeper().Help(context, member.NewAuxiliaryRequest(identity.GetClassificationID(), nil, mutables)); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.maintainAuxiliary.GetKeeper().Help(context, maintain.NewAuxiliaryRequest(identity.GetClassificationID(), message.FromID, mutables)); err != nil {
		return nil, err
	}

	updatedIdentity := base.NewIdentity(identity.GetClassificationID(), identity.GetImmutables(), baseQualified.NewMutables(identity.GetMutables().GetMutablePropertyList().Mutate(baseLists.AnyPropertiesToProperties(mutables.GetMutablePropertyList().Get()...)...)))

	if err := updatedIdentity.ValidateBasic(); err != nil {
		return nil, err
	}

	identities.Mutate(record.NewRecord(updatedIdentity))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case maintain.Auxiliary.GetName():
				transactionKeeper.maintainAuxiliary = value
			case member.Auxiliary.GetName():
				transactionKeeper.memberAuxiliary = value
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
