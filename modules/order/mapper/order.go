package mapper

import (
	"github.com/persistenceOne/persistenceSDK/types"
)

type order struct {
	ID         types.ID
	Burn       types.Height
	Lock       types.Height
	Immutables types.Immutables
	Mutables   types.Mutables
}

var _ types.InterNFT = (*order)(nil)

func (order order) GetID() types.ID {
	return order.ID
}

func (order order) GetChainID() types.ID {
	return orderIDFromInterface(order.ID).ChainID
}

func (order order) GetClassificationID() types.ID {
	return orderIDFromInterface(order.ID).ClassificationID
}

func (order order) GetBurn() types.Height {
	return order.Burn
}

func (order order) CanBurn(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(order.Burn)
}

func (order order) GetLock() types.Height {
	return order.Lock
}

func (order order) CanSend(currentHeight types.Height) bool {
	return currentHeight.IsGreaterThan(order.Lock)
}

func (order order) GetImmutables() types.Immutables {
	return order.Immutables
}

func (order order) GetMutables() types.Mutables {
	return order.Mutables
}

func NewOrder(orderID types.ID, burn types.Height, lock types.Height, immutables types.Immutables, mutables types.Mutables) types.InterNFT {
	return order{
		ID:         orderID,
		Burn:       burn,
		Lock:       lock,
		Immutables: immutables,
		Mutables:   mutables,
	}
}
