/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package unprovision

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	identityID := message.IdentityID
	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(identityID))

	identity := identities.Get(key.FromID(identityID))
	if identity == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}

	//  TODO implement
	//  if !identity.(mappables.InterIdentity).IsProvisioned(message.From) {
	//	return newTransactionResponse(errors.NotAuthorized)
	//  }
	//
	//  if !identity.(mappables.InterIdentity).IsProvisioned(message.To) {
	//	return newTransactionResponse(errors.EntityNotFound)
	//  }
	//
	//  if identity.(mappables.InterIdentity).IsUnprovisioned(message.To) {
	//	return newTransactionResponse(errors.DeletionNotAllowed)
	//  }
	//
	//  identities.Mutate(identity.(mappables.InterIdentity).UnprovisionAddress(message.To))

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper
	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
