package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type orders struct {
	ID   types.ID
	List []mappables.Order

	mapper  helpers.Mapper
	context sdkTypes.Context
}

var _ mappers.Orders = (*orders)(nil)

func (orders orders) GetID() types.ID { return orders.ID }
func (orders orders) Get(id types.ID) mappables.Order {
	orderID := orderIDFromInterface(id)
	for _, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(orderID) == 0 {
			return oldOrder
		}
	}
	return nil
}
func (orders orders) GetList() []mappables.Order {
	return orders.List
}

func (orders orders) Fetch(id types.ID) mappers.Orders {
	var orderList []mappables.Order
	ordersID := orderIDFromInterface(id)
	if len(ordersID.HashID.Bytes()) > 0 {
		mappable := orders.mapper.Read(orders.context, ordersID)
		if mappable != nil {
			orderList = append(orderList, mappable.(order))
		}
	} else {
		appendOrderList := func(mappable traits.Mappable) bool {
			orderList = append(orderList, mappable.(order))
			return false
		}
		orders.mapper.Iterate(orders.context, ordersID, appendOrderList)
	}
	orders.ID, orders.List = id, orderList
	return orders
}
func (orders orders) Add(order mappables.Order) mappers.Orders {
	orders.ID = readOrderID("")
	orders.mapper.Create(orders.context, order)
	for i, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(order.GetID()) < 0 {
			orders.List = append(append(orders.List[:i], order), orders.List[i+1:]...)
			break
		}
	}
	return orders
}
func (orders orders) Remove(order mappables.Order) mappers.Orders {
	orders.mapper.Delete(orders.context, order.GetID())
	for i, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(order.GetID()) == 0 {
			orders.List = append(orders.List[:i], orders.List[i+1:]...)
			break
		}
	}
	return orders
}
func (orders orders) Mutate(order mappables.Order) mappers.Orders {
	orders.mapper.Update(orders.context, order)
	for i, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(order.GetID()) == 0 {
			orders.List[i] = order
			break
		}
	}
	return orders
}

func NewOrders(mapper helpers.Mapper, context sdkTypes.Context) mappers.Orders {
	return orders{
		ID:      readOrderID(""),
		List:    []mappables.Order{},
		mapper:  mapper,
		context: context,
	}

}
