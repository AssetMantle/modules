package mapper

import (
	"errors"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

type orders struct {
	ID   types.ID
	List []types.InterNFT

	mapper  ordersMapper
	context sdkTypes.Context
}

var _ types.InterNFTs = (*orders)(nil)

func (orders orders) GetID() types.ID { return orders.ID }
func (orders orders) Get(id types.ID) types.InterNFT {
	orderID := orderIDFromInterface(id)
	for _, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(orderID) == 0 {
			return oldOrder
		}
	}
	return nil
}
func (orders orders) GetList() []types.InterNFT {
	return orders.List
}

func (orders orders) Fetch(id types.ID) types.InterNFTs {
	var orderList []types.InterNFT
	ordersID := orderIDFromInterface(id)
	if len(ordersID.HashID.Bytes()) > 0 {
		order := orders.mapper.read(orders.context, ordersID)
		if order != nil {
			orderList = append(orderList, order)
		}
	} else {
		appendOrderList := func(order types.InterNFT) bool {
			orderList = append(orderList, order)
			return false
		}
		orders.mapper.iterate(orders.context, ordersID, appendOrderList)
	}
	orders.ID, orders.List = id, orderList
	return orders
}
func (orders orders) Add(order types.InterNFT) types.InterNFTs {
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
func (orders orders) Remove(order types.InterNFT) types.InterNFTs {
	orders.mapper.delete(orders.context, order.GetID())
	for i, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(order.GetID()) == 0 {
			orders.List = append(orders.List[:i], orders.List[i+1:]...)
			break
		}
	}
	return orders
}
func (orders orders) Mutate(order types.InterNFT) types.InterNFTs {
	orders.mapper.update(orders.context, order)
	for i, oldOrder := range orders.List {
		if oldOrder.GetID().Compare(order.GetID()) == 0 {
			orders.List[i] = order
			break
		}
	}
	return orders
}

func NewOrders(Mapper types.Mapper, context sdkTypes.Context) types.InterNFTs {
	switch mapper := Mapper.(type) {
	case ordersMapper:
		return orders{
			ID:      readOrderID(""),
			List:    []types.InterNFT{},
			mapper:  mapper,
			context: context,
		}
	default:
		panic(errors.New(fmt.Sprintf("incorrect mapper initialization for module, %v", ModuleName)))
	}

}
