package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.OrderID = (*OrderIDI)(nil)

func (orderIDI *OrderIDI) Compare(listable traits.Listable) int {
	return orderIDI.Impl.(ids2.OrderID).Compare(listable)
}

func (orderIDI *OrderIDI) Bytes() []byte {
	return orderIDI.Impl.(ids2.OrderID).Bytes()
}

func (orderIDI *OrderIDI) IsOrderID() {

}
