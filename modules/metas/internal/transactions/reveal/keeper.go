// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/modules/metas/internal/mappable"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type transactionKeeper struct {
	mapper     helpers.Mapper
	parameters helpers.Parameters
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	metaID := key.GenerateMetaID(message.Data)
	metas := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(metaID))

	meta := metas.Get(key.FromID(metaID))
	if meta != nil {
		return newTransactionResponse(constants.EntityAlreadyExists)
	}

	if message.Data.GenerateHash().Compare(baseIDs.NewID("")) != 0 {
		metas.Add(mappable.NewMeta(message.Data))
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, _ []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
