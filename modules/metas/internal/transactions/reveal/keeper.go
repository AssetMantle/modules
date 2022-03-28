/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package reveal

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionKeeper struct {
	mapper     helpers.Mapper
	parameters helpers.Parameters
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	msgServer := NewMsgServerImpl(transactionKeeper)

	_, Error := msgServer.Reveal(sdkTypes.WrapSDKContext(context), &message)
	return newTransactionResponse(Error)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, _ []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters
	return transactionKeeper
}
func (transactionKeeper transactionKeeper) RegisterService(configurator sdkModule.Configurator) {
	RegisterMsgServer(configurator.MsgServer(), NewMsgServerImpl(transactionKeeper))
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
