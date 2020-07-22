package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type orders struct {
	ID   schema.ID
	List []schema.Order

	mapper  ordersMapper
	context sdkTypes.Context
}

var _ schema.Orders = (*orders)(nil)

func (orders orders) GetID() schema.ID { return orders.ID }
func (orders orders) Get(id schema.ID) schema.Order {
	orderID := orderIDFromInterface(id)
	for _, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(orderID) == 0 {
			return oldOrder
		}
	}
	return nil
}
func (orders orders) GetList() []schema.Order {
	return orders.List
}

func (orders orders) Fetch(id schema.ID) schema.Orders {
	var orderList []schema.Order
	orderID := orderIDFromInterface(id)
	if len(orderID.HashID.Bytes()) > 0 {
		order := orders.mapper.read(orders.context, orderID)
		if order != nil {
			orderList = append(orderList, order)
		}
	} else {
		appendOrderList := func(order schema.Order) bool {
			orderList = append(orderList, order)
			return false
		}
		orders.mapper.iterate(orders.context, orderID, appendOrderList)
	}
	orders.ID, orders.List = id, orderList
	return orders
}
func (orders orders) Add(order schema.Order) schema.Orders {
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
func (orders orders) Remove(order schema.Order) schema.Orders {
	orders.mapper.delete(orders.context, order.GetID())
	for i, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(order.GetID()) == 0 {
			orders.List = append(orders.List[:i], orders.List[i+1:]...)
			break
		}
	}
	return orders
}
func (orders orders) Mutate(order schema.Order) schema.Orders {
	orders.mapper.update(orders.context, order)
	for i, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(order.GetID()) == 0 {
			orders.List[i] = order
			break
		}
	}
	return orders
}

func NewOrders(Mapper utility.Mapper, context sdkTypes.Context) schema.Orders {
	switch mapper := Mapper.(type) {
	case ordersMapper:
		return orders{
			ID:      readOrderID(""),
			List:    []schema.Order{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleName)))
	}

}
