// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package name

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/identities/key"
	"github.com/AssetMantle/modules/x/identities/record"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents/base"
)

type transactionKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	nameIdentity := base.NewNameIdentity(message.Name, baseData.NewListData(baseData.NewAccAddressData(message.GetFromAddress())))

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(nameIdentity.GetNameIdentityID()))
	if identities.GetMappable(key.NewKey(nameIdentity.GetNameIdentityID())) != nil {
		return nil, errorConstants.EntityAlreadyExists.Wrapf("name identity with ID %s already exists", nameIdentity.GetNameIdentityID().AsString())
	}

	if err := nameIdentity.ValidateBasic(); err != nil {
		return nil, err
	}

	identities.Add(record.NewRecord(nameIdentity))

	return newTransactionResponse(nameIdentity.GetNameIdentityID()), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
