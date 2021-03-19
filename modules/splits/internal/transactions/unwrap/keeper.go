/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package unwrap

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/utilities"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	parameters      helpers.Parameters
	supplyKeeper    supply.Keeper
	verifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	splits := transactionKeeper.mapper.NewCollection(context)
	if _, Error := utilities.SubtractSplits(splits, message.FromID, message.OwnableID, sdkTypes.NewDecFromInt(message.Value)); Error != nil {
		return newTransactionResponse(Error)
	}

	if Error := transactionKeeper.supplyKeeper.SendCoinsFromModuleToAccount(context, module.Name, message.From, sdkTypes.NewCoins(sdkTypes.NewCoin(message.OwnableID.String(), message.Value))); Error != nil {
		return newTransactionResponse(Error)
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case supply.Keeper:
			transactionKeeper.supplyKeeper = value
		case helpers.Auxiliary:
			switch value.GetName() {
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			default:
				break
			}
		default:
			panic(errors.UninitializedUsage)
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
