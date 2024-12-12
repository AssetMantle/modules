// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package govern

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionKeeper struct {
	parameterManager helpers.ParameterManager
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return nil, constants.InvalidRequest.Wrapf("internal transaction,  please use the gov module to  create proposal to update params")
}
func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if _, err := transactionKeeper.parameterManager.Fetch(context).Set(message.Parameter).Update(context); err != nil {
		return nil, constants.InvalidRequest.Wrapf(err.Error())
	}

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(_ helpers.Mapper, parameterManager helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	transactionKeeper.parameterManager = parameterManager

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
