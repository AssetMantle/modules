package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Orders interface {
	GetID() types.ID

	Get(types.ID) mappables.Order
	GetList() []mappables.Order

	Fetch(types.ID) Orders
	Add(mappables.Order) Orders
	Remove(mappables.Order) Orders
	Mutate(mappables.Order) Orders
}
