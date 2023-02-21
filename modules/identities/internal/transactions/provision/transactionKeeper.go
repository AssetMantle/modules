// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package provision

import (
	"context"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/schema/data"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/properties/constants"
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
	Mappable := identities.Get(key.NewKey(message.IdentityID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound
	}
	identity := mappable.GetIdentity(Mappable)

	if identity.GetProvisionedAddressCount() >= transactionKeeper.parameterManager.GetParameter(constants.MaxProvisionAddressCountProperty.GetID()).GetMetaProperty().GetData().Get().(data.NumberData).Get() {
		return nil, errorConstants.NotAuthorized
	}

	fromAddress, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	toAddress, err := sdkTypes.AccAddressFromBech32(message.To)
	if err != nil {
		panic("Could not get To address from Bech32 string")
	}

	if !identity.IsProvisioned(fromAddress) {
		return nil, errorConstants.NotAuthorized
	}

	if identity.IsProvisioned(toAddress) {
		return nil, errorConstants.EntityAlreadyExists
	}

	identities.Mutate(mappable.NewMappable(identity.ProvisionAddress(toAddress)))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper
	transactionKeeper.parameterManager = parameterManager

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

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case supplement.Auxiliary.GetName():
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
