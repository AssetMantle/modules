// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/revoke"
)

type transactionKeeper struct {
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
	revokeAuxiliary  helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if Mappable := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.GetFromIdentityID())).GetMappable(key.NewKey(message.GetFromIdentityID())); Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.GetFromIdentityID().AsString())
	} else if identity := mappable.GetIdentity(Mappable); !identity.IsProvisioned(message.GetFromAddress()) {
		return nil, errorConstants.NotAuthorized.Wrapf("address %s is not provisioned for identity with ID %s", message.GetFromAddress().String(), message.GetFromIdentityID().AsString())
	}

	if _, err := transactionKeeper.revokeAuxiliary.GetKeeper().Help(context, revoke.NewAuxiliaryRequest(message.GetFromIdentityID(), message.ToID, message.ClassificationID)); err != nil {
		return nil, err
	}

	return newTransactionResponse(), nil
}
func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case revoke.Auxiliary.GetName():
				transactionKeeper.revokeAuxiliary = value
			}
		}
	}

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
