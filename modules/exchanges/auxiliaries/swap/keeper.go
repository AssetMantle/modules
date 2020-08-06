/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package swap

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryKeeper struct {
	mapper     helpers.Mapper
	BankKeeper bank.Keeper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) error {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	order := auxiliaryRequest.Order

	switch makerData := order.GetMakerAssetData().(type) {
	case sdkTypes.Coin:
		switch takerData := order.GetTakerAssetData().(type) {
		case sdkTypes.Coin:
			if Error := swapCoins(context, auxiliaryKeeper.BankKeeper, order.GetMakerAddress(), sdkTypes.NewCoins(makerData),
				order.GetTakerAddress(), sdkTypes.NewCoins(takerData)); Error != nil {
				return Error
			}
		default:
			return errors.New(fmt.Sprintf("Unknown message type, %v, not supported", takerData))
		}
	default:
		return errors.New(fmt.Sprintf("Unknown message type, %v, not supported", makerData))
	}
	return nil
}

func swapCoins(context sdkTypes.Context, bankKeeper bank.Keeper, makerAddress sdkTypes.AccAddress, makeOrderCoins sdkTypes.Coins,
	takerAddress sdkTypes.AccAddress, takeOrderCoins sdkTypes.Coins) error {
	if Error := bankKeeper.SendCoins(context, makerAddress, takerAddress, makeOrderCoins); Error != nil {
		return Error
	}
	if Error := bankKeeper.SendCoins(context, takerAddress, makerAddress, takeOrderCoins); Error != nil {
		return Error
	}
	return nil
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, externalKeepers []interface{}) helpers.AuxiliaryKeeper {
	auxiliaryKeeper := auxiliaryKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case bank.Keeper:
			auxiliaryKeeper.BankKeeper = value
		}
	}
	return auxiliaryKeeper
}
