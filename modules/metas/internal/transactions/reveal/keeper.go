// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/modules/metas/internal/mappable"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables/base"
)

type transactionKeeper struct {
	mapper     helpers.Mapper
	parameters helpers.Parameters
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	metaID := baseIDs.NewMetaID(message.Data.GetType(), message.Data.GenerateHashID())
	metas := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(metaID))

	meta := metas.Get(key.NewKey(metaID))
	if meta != nil {
		return newTransactionResponse(constants.EntityAlreadyExists)
	}

	if message.Data.GenerateHashID().Compare(baseIDs.GenerateHashID()) != 0 {
		metas.Add(mappable.NewMappable(base.NewMeta(message.Data)))
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
