/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package reveal

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper     helpers.Mapper
	parameters helpers.Parameters
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	metaID := key.NewMetaID(base.NewID(message.MetaFact.GetData().GenerateHash()))
	metas := transactionKeeper.mapper.NewCollection(context).Fetch(key.New(metaID))
	meta := metas.Get(key.New(metaID)).(mappables.Meta)
	if meta != nil {
		return newTransactionResponse(errors.EntityAlreadyExists)
	}
	metas.Add(mappable.NewMeta(message.MetaFact.GetData()))
	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, _ []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
