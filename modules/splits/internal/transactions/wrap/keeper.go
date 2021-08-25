/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package wrap

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	parameters      helpers.Parameters
	bankKeeper      bankKeeper.Keeper
	verifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	msgServer := NewMsgServerImpl(transactionKeeper)

	_, Error := msgServer.Wrap(sdkTypes.WrapSDKContext(context), &message)
	return newTransactionResponse(Error)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case bankKeeper.Keeper:
			transactionKeeper.bankKeeper = value
		case helpers.Auxiliary:
			switch value.GetName() {
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			}
		}
	}

	return transactionKeeper
}
func (transactionKeeper transactionKeeper) RegisterService(configurator module.Configurator) {
	RegisterMsgServer(configurator.MsgServer(), NewMsgServerImpl(transactionKeeper))
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
