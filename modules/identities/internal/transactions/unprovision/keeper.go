// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/utilities"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	supplementAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	identityID := message.IdentityID
	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(identityID))

	identity := identities.Get(key.FromID(identityID)).(mappables.Identity)
	if identity == nil {
		return newTransactionResponse(errorConstants.EntityNotFound)
	}

	if provisioned, err := utilities.IsProvisioned(context, transactionKeeper.supplementAuxiliary, identity, message.From); err != nil || !provisioned {
		return newTransactionResponse(errorConstants.NotAuthorized)
	}

	if provisioned, err := utilities.IsProvisioned(context, transactionKeeper.supplementAuxiliary, identity, message.To); err != nil {
		return newTransactionResponse(err)
	} else if provisioned {
		return newTransactionResponse(errorConstants.EntityAlreadyExists)
	}

	if updatedIdentity, err := utilities.UnprovisionAddress(context, transactionKeeper.supplementAuxiliary, identity, message.To); err != nil {
		return newTransactionResponse(err)
	} else {
		identities.Mutate(updatedIdentity)
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case define.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			}
		default:
			panic(errorConstants.UninitializedUsage)
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
