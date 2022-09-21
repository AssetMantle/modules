// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/utilities"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
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

	mappable := identities.Get(key.FromID(identityID))
	if mappable == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}
	identity := mappable.(mappables.Identity)

	if found, err := utilities.IsProvisioned(context, transactionKeeper.supplementAuxiliary, identity, message.From); err != nil {
		return newTransactionResponse(err)
	} else if !found {
		return newTransactionResponse(errors.NotAuthorized)
	}

	if found, err := utilities.IsProvisioned(context, transactionKeeper.supplementAuxiliary, identity, message.To); err != nil {
		return newTransactionResponse(err)
	} else if !found {
		return newTransactionResponse(errors.NotAuthorized)
	}

	if identity, err := utilities.UnprovisionAddress(context, transactionKeeper.supplementAuxiliary, identity, message.To); err != nil {
		return newTransactionResponse(err)
	} else {
		identities.Mutate(identity)
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			}
		default:
			panic(errors.UninitializedUsage)
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
