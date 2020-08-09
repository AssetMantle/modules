/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package take

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	assetsMapper "github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/auxiliaries/swap"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionKeeper struct {
	mapper                    helpers.Mapper
	exchangesSwapAuxiliary    helpers.Auxiliary
	identitiesVerifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	orderID := message.OrderID
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(orderID)
	order := orders.Get(orderID)
	if order == nil {
		return constants.EntityNotFound
	}
	if Error := transactionKeeper.identitiesVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); Error != nil {
		return Error
	}

	makerID := base.NewID(order.GetImmutables().Get().Get(base.NewID(constants.MakerIDProperty)).GetFact().GetHash())
	makerSplitID := base.NewID(order.GetImmutables().Get().Get(base.NewID(constants.MakerSplitIDProperty)).GetFact().GetHash())
	makerSplit, Error := sdkTypes.NewDecFromStr(order.GetMutables().Get().Get(base.NewID(constants.MakerSplitProperty)).GetFact().GetHash())
	takerID := base.NewID(order.GetImmutables().Get().Get(base.NewID(constants.TakerIDProperty)).GetFact().GetHash())
	takerSplitID := base.NewID(order.GetImmutables().Get().Get(base.NewID(constants.TakerSplitIDProperty)).GetFact().GetHash())
	exchangeRate, Error := sdkTypes.NewDecFromStr(order.GetImmutables().Get().Get(base.NewID(constants.ExchangeRateProperty)).GetFact().GetHash())
	if Error != nil {
		return Error
	}

	makerIsAsset := assetsMapper.NewAssets(assetsMapper.Mapper, context).Fetch(makerSplitID).Get(makerSplitID) != nil
	takerIsAsset := assetsMapper.NewAssets(assetsMapper.Mapper, context).Fetch(takerSplitID).Get(takerSplitID) != nil

	if makerIsAsset && takerIsAsset {
		if !sdkTypes.OneDec().Equal(message.TakerSplit) || !sdkTypes.OneDec().Equal(makerSplit) {
			return constants.IncorrectMessage
		}
	} else if !makerIsAsset && takerIsAsset {
		if !sdkTypes.OneDec().Equal(message.TakerSplit) {
			return constants.IncorrectMessage
		}
	} else if makerIsAsset && !takerIsAsset {
		if makerSplit.Mul(exchangeRate).GT(message.TakerSplit) {
			return constants.IncorrectMessage
		} else {
			message.TakerSplit = makerSplit.Mul(exchangeRate)
		}
	} else {
		if makerSplit.Mul(exchangeRate).LTE(message.TakerSplit) {
			message.TakerSplit = makerSplit.Mul(exchangeRate)
		}
	}

	makerSplitDeduction := message.TakerSplit.Quo(exchangeRate)

	if takerID.String() != "" && message.FromID.Compare(takerID) != 0 {
		return constants.NotAuthorized
	}
	if Error := transactionKeeper.exchangesSwapAuxiliary.GetKeeper().Help(context, swap.NewAuxiliaryRequest(makerID,
		makerSplitDeduction, makerSplitID, message.FromID, message.TakerSplit, takerSplitID)); Error != nil {
		return Error
	}

	makerSplit = makerSplit.Sub(makerSplitDeduction)
	mutables := base.NewMutables(order.GetMutables().Get().Mutate(base.NewProperty(base.NewID(constants.MakerSplitProperty),
		base.NewFact(makerSplit.String(), true))), order.GetMutables().GetMaintainersID())
	order = mapper.NewOrder(order.GetID(), mutables, order.GetImmutables())
	orders = orders.Mutate(order)
	if makerSplit.IsZero() {
		orders.Remove(order)
	}

	return nil
}

func initializeTransactionKeeper(mapper helpers.Mapper, auxiliaries []interface{}) helpers.TransactionKeeper {
	transactionKeeper := transactionKeeper{mapper: mapper}
	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case swap.Auxiliary.GetName():
				transactionKeeper.exchangesSwapAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.identitiesVerifyAuxiliary = value
			}
		}
	}
	return transactionKeeper
}
