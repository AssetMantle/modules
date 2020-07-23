package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type orders struct {
	ID   types.ID
	List []entities.Order

	mapper  ordersMapper
	context sdkTypes.Context
}

var _ mappers.Orders = (*orders)(nil)

func (orders orders) GetID() types.ID { return orders.ID }
func (orders orders) Get(id types.ID) entities.Order {
	orderID := orderIDFromInterface(id)
	for _, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(orderID) == 0 {
			return oldOrder
		}
	}
	return nil
}
func (orders orders) GetList() []entities.Order {
	return orders.List
}

func (orders orders) Fetch(id types.ID) mappers.Orders {
	var orderList []entities.Order
	ordersID := orderIDFromInterface(id)
	if len(ordersID.HashID.Bytes()) > 0 {
		order := orders.mapper.read(orders.context, ordersID)
		if order != nil {
			orderList = append(orderList, order)
		}
	} else {
		appendOrderList := func(order entities.Order) bool {
			orderList = append(orderList, order)
			return false
		}
		orders.mapper.iterate(orders.context, ordersID, appendOrderList)
	}
	orders.ID, orders.List = id, orderList
	return orders
}
func (orders orders) Add(order entities.Order) mappers.Orders {
	orders.ID = readOrderID("")
	orders.mapper.create(orders.context, order)
	for i, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(order.GetID()) < 0 {
			orders.List = append(append(orders.List[:i], order), orders.List[i+1:]...)
			break
		}
	}
	return orders
}
func (orders orders) Remove(order entities.Order) mappers.Orders {
	orders.mapper.delete(orders.context, order.GetID())
	for i, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(order.GetID()) == 0 {
			orders.List = append(orders.List[:i], orders.List[i+1:]...)
			break
		}
	}
	return orders
}
func (orders orders) Mutate(order entities.Order) mappers.Orders {
	orders.mapper.update(orders.context, order)
	for i, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(order.GetID()) == 0 {
			orders.List[i] = order
			break
		}
	}
	return orders
}

func NewOrders(Mapper utilities.Mapper, context sdkTypes.Context) mappers.Orders {
	switch mapper := Mapper.(type) {
	case ordersMapper:
		return orders{
			ID:      readOrderID(""),
			List:    []entities.Order{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleName)))
	}

}
