// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package provision

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"

	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/record"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	parameterManager    helpers.ParameterManager
	supplementAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.IdentityID))
	Mappable := identities.GetMappable(key.NewKey(message.IdentityID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.IdentityID.AsString())
	}
	identity := mappable.GetIdentity(Mappable)

	if identity.GetProvisionedAddressCount().GTE(transactionKeeper.parameterManager.Fetch(context).GetParameter(constants.MaxProvisionAddressCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get()) {
		return nil, errorConstants.NotAuthorized.Wrapf("identity with ID %s has reached the maximum allowed number of provision-able addresses %d", message.IdentityID.AsString(), transactionKeeper.parameterManager.Fetch(context).GetParameter(constants.MaxProvisionAddressCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get())
	}

	fromAddress := message.GetFromAddress()

	toAddress, err := sdkTypes.AccAddressFromBech32(message.To)
	if err != nil {
		panic("Could not get To address from Bech32 string")
	}

	if !identity.IsProvisioned(fromAddress) {
		return nil, errorConstants.NotAuthorized.Wrapf("address %s is not provisioned by identity with ID %s", fromAddress.String(), message.IdentityID.AsString())
	}

	if identity.IsProvisioned(toAddress) {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("address %s is already provisioned", toAddress.String())
	}

	updatedIdentity := identity.ProvisionAddress(toAddress)

	if err := updatedIdentity.ValidateBasic(); err != nil {
		return nil, err
	}

	identities.Mutate(record.NewRecord(updatedIdentity))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper
	transactionKeeper.parameterManager = parameterManager

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
