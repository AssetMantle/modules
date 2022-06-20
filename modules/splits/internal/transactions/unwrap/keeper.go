// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unwrap

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/modules/modules/splits/internal/utilities"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
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
	if _, err := utilities.SubtractSplits(splits, message.FromID, message.OwnableID, sdkTypes.NewDecFromInt(message.Value)); err != nil {
		return newTransactionResponse(err)
	}

	if err := transactionKeeper.supplyKeeper.SendCoinsFromModuleToAccount(context, module.Name, message.From, sdkTypes.NewCoins(sdkTypes.NewCoin(message.OwnableID.String(), message.Value))); err != nil {
		return newTransactionResponse(err)
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
			panic(constants.UninitializedUsage)
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
