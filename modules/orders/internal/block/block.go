/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package block

import (
	"sort"
	"strings"

	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type block struct {
	mapper              helpers.Mapper
	parameters          helpers.Parameters
	supplementAuxiliary helpers.Auxiliary
	transferAuxiliary   helpers.Auxiliary
}

var _ helpers.Block = (*block)(nil)

type orderBook struct {
	buyOrders  []mappables.Order
	sellOrders []mappables.Order
}

func (orderBook orderBook) addBuyOrder(order mappables.Order) orderBook {
	if len(orderBook.buyOrders) == 0 {
		orderBook.buyOrders = []mappables.Order{order}
	} else {
		orderBook.buyOrders = append(orderBook.buyOrders, order)
	}

	return orderBook
}

func (orderBook orderBook) addSellOrder(order mappables.Order) orderBook {
	if len(orderBook.sellOrders) == 0 {
		orderBook.sellOrders = []mappables.Order{order}
	} else {
		orderBook.sellOrders = append(orderBook.sellOrders, order)
	}

	return orderBook
}

func (orderBook orderBook) removeBuyOrder(order mappables.Order) orderBook {
	for i, buyOrder := range orderBook.buyOrders {
		if buyOrder.GetClassificationID().Equals(order.GetClassificationID()) && buyOrder.GetMakerID().Equals(order.GetMakerID()) && buyOrder.GetMakerOwnableID().Equals(order.GetMakerOwnableID()) && buyOrder.GetTakerOwnableID().Equals(order.GetTakerOwnableID()) && buyOrder.GetImmutables().GetHashID().Equals(order.GetImmutables().GetHashID()) {
			orderBook.buyOrders = append(orderBook.buyOrders[:i], orderBook.buyOrders[i:]...)
			break
		}
	}

	return orderBook
}

func (orderBook orderBook) removeSellOrder(order mappables.Order) orderBook {
	for i, sellOrder := range orderBook.sellOrders {
		if sellOrder.GetClassificationID().Equals(order.GetClassificationID()) && sellOrder.GetMakerID().Equals(order.GetMakerID()) && sellOrder.GetMakerOwnableID().Equals(order.GetMakerOwnableID()) && sellOrder.GetTakerOwnableID().Equals(order.GetTakerOwnableID()) && sellOrder.GetImmutables().GetHashID().Equals(order.GetImmutables().GetHashID()) {
			orderBook.sellOrders = append(orderBook.sellOrders[:i], orderBook.sellOrders[i:]...)
			break
		}
	}

	return orderBook
}

func (block block) Begin(context sdkTypes.Context, _ abciTypes.RequestBeginBlock) {
	orderBookAssetsMap := make(map[string]orderBook)
	accumulator := func(mappable helpers.Mappable) bool {
		order := mappable.(mappables.Order)
		keysList := []string{order.GetMakerOwnableID().String(), order.GetTakerOwnableID().String()}
		sort.Strings(keysList)
		orderBookKey := strings.Join(keysList, "")
		orderBook := orderBookAssetsMap[orderBookKey]

		switch {
		case len(orderBook.buyOrders) == 0 && len(orderBook.sellOrders) == 0:
			orderBook = orderBook.addBuyOrder(order)
		case len(orderBook.buyOrders) != 0 && len(orderBook.sellOrders) == 0:
			if orderBook.buyOrders[0].GetMakerOwnableID().Equals(order.GetMakerOwnableID()) {
				orderBook = orderBook.addBuyOrder(order)
			} else {
				orderBook = orderBook.addSellOrder(order)
			}
		case len(orderBook.buyOrders) == 0 && len(orderBook.sellOrders) != 0:
			if orderBook.sellOrders[0].GetMakerOwnableID().Equals(order.GetMakerOwnableID()) {
				orderBook = orderBook.addSellOrder(order)
			} else {
				orderBook = orderBook.addBuyOrder(order)
			}
		default:
			if orderBook.buyOrders[0].GetMakerOwnableID().Equals(order.GetMakerOwnableID()) {
				orderBook = orderBook.addBuyOrder(order)
			} else {
				orderBook = orderBook.addSellOrder(order)
			}
		}

		orderBookAssetsMap[orderBookKey] = orderBook

		return false
	}
	block.mapper.NewCollection(context).Iterate(key.New(base.NewID("")), accumulator)
}

func (block block) End(_ sdkTypes.Context, _ abciTypes.RequestEndBlock) {

}

func (block block) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaryKeepers ...interface{}) helpers.Block {
	block.mapper, block.parameters = mapper, parameters

	for _, auxiliaryKeeper := range auxiliaryKeepers {
		switch value := auxiliaryKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case supplement.Auxiliary.GetName():
				block.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				block.transferAuxiliary = value
			}
		default:
			panic(errors.UninitializedUsage)
		}
	}

	return block
}
