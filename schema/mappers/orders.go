package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Orders interface {
	GetID() types.ID

	Get(types.ID) entities.Order
	GetList() []entities.Order

	Fetch(types.ID) Orders
	Add(entities.Order) Orders
	Remove(entities.Order) Orders
	Mutate(entities.Order) Orders
}
