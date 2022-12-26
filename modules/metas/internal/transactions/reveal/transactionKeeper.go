// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/modules/metas/internal/mappable"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type transactionKeeper struct {
	mapper     helpers.Mapper
	parameters helpers.Parameters
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context types.Context, message helpers.Message) helpers.TransactionResponse {
	_, err := transactionKeeper.Handle(context.Context(), message.(*Message))
	return newTransactionResponse(err)
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*Response, error) {
	dataID := baseIDs.GenerateDataID(message.Data)
	metas := transactionKeeper.mapper.NewCollection(types.UnwrapSDKContext(context)).Fetch(key.NewKey(dataID))

	data := metas.Get(key.NewKey(dataID))
	if data != nil {
		return nil, constants.EntityAlreadyExists
	}

	if message.Data.GenerateHashID().Compare(baseIDs.GenerateHashID()) != 0 {
		metas.Add(mappable.NewMappable(message.Data))
	}

	return &Response{}, nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, _ []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
