// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package quash

import (
	"context"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/unbond"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseTypes "github.com/AssetMantle/schema/go/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	authenticateAuxiliary helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	unbondAuxiliary       helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	fromAddress, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.IdentityID))

	Mappable := identities.Get(key.NewKey(message.IdentityID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("identity with ID %s not found", message.IdentityID.AsString())
	}
	identity := mappable.GetIdentity(Mappable)

	// TODO: Check if identity expired identities must be quashed at the end of block
	if identity.GetExpiry().Compare(baseTypes.NewHeight(sdkTypes.UnwrapSDKContext(context).BlockHeight())) > 0 {
		return nil, errorConstants.NotAuthorized.Wrapf("identity with ID %s is not expired yet", message.IdentityID.AsString())
	}

	if _, err := transactionKeeper.unbondAuxiliary.GetKeeper().Help(context, unbond.NewAuxiliaryRequest(identity.GetClassificationID(), fromAddress)); err != nil {
		return nil, err
	}

	identities.Remove(mappable.NewMappable(identity))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case unbond.Auxiliary.GetName():
				transactionKeeper.unbondAuxiliary = value
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
