/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package wrap

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper                    helpers.Mapper
	bankKeeper                bank.Keeper
	identitiesVerifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	if Error := transactionKeeper.identitiesVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); Error != nil {
		return Error
	}
	if Error := transactionKeeper.bankKeeper.SendCoinsFromAccountToModule(context, message.From, mapper.ModuleName, message.Coins); Error != nil {
		return Error
	}
	for _, coin := range message.Coins {
		splitID := mapper.NewSplitID(message.FromID, base.NewID(coin.Denom))
		splits := mapper.NewSplits(transactionKeeper.mapper, context).Fetch(splitID)
		split := splits.Get(splitID)
		if split == nil {
			splits.Add(mapper.NewSplit(splitID, sdkTypes.NewDecFromInt(coin.Amount)))
		} else {
			splits.Mutate(split.Receive(sdkTypes.NewDecFromInt(coin.Amount)).(mappables.Split))
		}
	}
	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case bank.Keeper:
			transactionKeeper.bankKeeper = value
		case helpers.Auxiliary:
			switch value.GetName() {
			case verify.Auxiliary.GetName():
				transactionKeeper.identitiesVerifyAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
