package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.OrderID = (*OrderIDI_OrderID)(nil)

func (orderID *OrderIDI_OrderID) String() string {
	return orderID.OrderID.String()
}
func (orderID *OrderIDI_OrderID) Compare(listable traits.Listable) int {
	return bytes.Compare(orderID.Bytes(), orderIDFromInterface(listable).Bytes())
}
func (orderID *OrderIDI_OrderID) Bytes() []byte {
	return orderID.OrderID.OrderId.Bytes()
}
func (orderID *OrderIDI_OrderID) IsOrderID() {}
func orderIDFromInterface(i interface{}) *OrderIDI {
	switch value := i.(type) {
	case *OrderIDI:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func GenerateOrderID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.OrderID {
	return NewOrderID(GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()))
}
func NewOrderID(hashID ids.HashID) ids.OrderID {
	return &OrderIDI{
		Impl: &OrderIDI_OrderID{
			OrderID: &OrderID{
				OrderId: hashID.(*HashIDI),
			},
		},
	}
}

func PrototypeOrderID() ids.OrderID {
	return NewOrderID(PrototypeHashID())
}

func ReadOrderID(orderIDString string) (ids.OrderID, error) {
	if hashID, err := ReadHashID(orderIDString); err == nil {
		return NewOrderID(hashID), nil
	}

	if orderIDString == "" {
		return PrototypeOrderID(), nil
	}

	return PrototypeOrderID(), nil
}
