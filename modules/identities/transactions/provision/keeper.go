/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package provision

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	identityID := message.IdentityID
	identities := mapper.NewIdentities(transactionKeeper.mapper, context).Fetch(identityID)
	identity := identities.Get(identityID)
	if identity == nil {
		return constants.EntityNotFound
	}
	if !identity.IsProvisioned(message.From) {
		return constants.NotAuthorized
	}
	if identity.IsProvisioned(message.To) {
		return constants.EntityAlreadyExists
	}
	if identity.IsUnprovisioned(message.To) {
		return constants.DeletionNotAllowed
	}
	identities.Mutate(identity.ProvisionAddress(message.To))
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, _ []interface{}) helpers.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
