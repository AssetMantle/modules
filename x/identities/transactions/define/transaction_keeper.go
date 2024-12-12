// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/utilities"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties/constants"
	"github.com/AssetMantle/schema/qualified/base"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	defineAuxiliary     helpers.Auxiliary
	superAuxiliary      helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}
func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	Mappable := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.GetFromIdentityID())).GetMappable(key.NewKey(message.GetFromIdentityID()))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.GetFromIdentityID().AsString())
	}
	identity := mappable.GetIdentity(Mappable)

	if !identity.IsProvisioned(message.GetFromAddress()) {
		return nil, errorConstants.NotAuthorized.Wrapf("address %s is not provisioned by identity with ID %s", message.GetFromAddress(), message.GetFromIdentityID().AsString())
	}

	immutables := base.NewImmutables(message.ImmutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(message.ImmutableProperties.Get()...)...))

	mutables := base.NewMutables(
		message.MutableMetaProperties.Add(
			baseLists.AnyPropertiesToProperties(
				message.MutableProperties.Add(
					constants.AuthenticationProperty,
				).Get()...,
			)...,
		),
	)

	auxiliaryResponse, err := transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(message.GetFromAddress(), immutables, mutables))
	if err != nil {
		return nil, err
	}
	classificationID := define.GetClassificationIDFromResponse(auxiliaryResponse)

	if _, err := transactionKeeper.superAuxiliary.GetKeeper().Help(context, super.NewAuxiliaryRequest(classificationID, message.GetFromIdentityID(), base.NewMutables(mutables.GetMutablePropertyList().Remove(constants.AuthenticationProperty)), utilities.SetModulePermissions(true, true)...)); err != nil {
		return nil, err
	}

	return newTransactionResponse(classificationID), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case define.Auxiliary.GetName():
				transactionKeeper.defineAuxiliary = value
			case super.Auxiliary.GetName():
				transactionKeeper.superAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			}
		}
	}

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
