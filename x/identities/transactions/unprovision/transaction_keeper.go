// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/record"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	supplementAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	identityID := message.IdentityID
	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(identityID))

	Mappable := identities.GetMappable(key.NewKey(identityID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", identityID.AsString())
	}
	identity := mappable.GetIdentity(Mappable)

	toAddress, _ := sdkTypes.AccAddressFromBech32(message.To)

	if !identity.IsProvisioned(message.GetFromAddress()) {
		return nil, errorConstants.NotAuthorized.Wrapf("address %s is not provisioned", message.GetFromAddress().String())
	}

	if !identity.IsProvisioned(toAddress) {
		return nil, errorConstants.EntityNotFound.Wrapf("address %s is not provisioned", toAddress.String())
	}

	updatedIdentity := identity.UnprovisionAddress(toAddress)

	if updatedIdentity.GetProvisionedAddressCount().IsZero() {
		return nil, errorConstants.InvalidRequest.Wrapf("identity %s must have at least one provisioned address", identityID.AsString())
	}

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
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
