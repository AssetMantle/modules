package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

var _ types.InterNFT = (*Order)(nil)

type Order struct {
	ID         types.ID
	From       sdkTypes.AccAddress
	SellOrder  sdkTypes.Coin
	BuyOrder   sdkTypes.Coin
	Mutables   types.Mutables
	Immutables types.Immutables
	Lock       types.Height
	Burn       types.Height
}

func (order Order) GetID() types.ID {
	return order.ID
}

func (order Order) GetBuyOrder() sdkTypes.Coin {
	return order.BuyOrder
}
func (order Order) GetSellOrder() sdkTypes.Coin {
	return order.SellOrder
}

func (order Order) GetChainID() types.ID {
	return orderIDFromInterface(order.ID).ChainID
}

func (order Order) GetClassificationID() types.ID {
	return orderIDFromInterface(order.ID).ClassificationID
}

func (order Order) GetMaintainersID() types.ID {
	return orderIDFromInterface(order.ID).MaintainersID
}

func (order Order) GetHashID() types.ID {
	return order.Immutables.GetHashID()
}

func (order Order) GetMutables() types.Mutables {
	return order.Mutables
}

func (order Order) GetImmutables() types.Immutables {
	return order.Immutables
}

func (order Order) GetLock() types.Height {
	return order.Lock
}

func (order Order) CanSend(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(order.Lock)
}

func (order Order) GetBurn() types.Height {
	return order.Burn
}

func (order Order) CanBurn(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(order.Burn)
}

func NewOrder(orderID types.ID, from sdkTypes.AccAddress, buyOrder sdkTypes.Coin, sellOrder sdkTypes.Coin, mutables types.Mutables, immutables types.Immutables, lock types.Height, burn types.Height) types.InterNFT {
	return Order{
		ID:         orderID,
		From:       from,
		BuyOrder:   buyOrder,
		SellOrder:  sellOrder,
		Mutables:   mutables,
		Immutables: immutables,
		Lock:       lock,
		Burn:       burn,
	}
}
