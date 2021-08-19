/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package revoke

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/revoke"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	parameters      helpers.Parameters
	revokeAuxiliary helpers.Auxiliary
	verifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	msgServer := NewMsgServerImpl(transactionKeeper)

	_, Error := msgServer.Revoke(sdkTypes.WrapSDKContext(context), &message)
	return newTransactionResponse(Error)
}
func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case revoke.Auxiliary.GetName():
				transactionKeeper.revokeAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			}
		default:
			panic(errors.UninitializedUsage)
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
