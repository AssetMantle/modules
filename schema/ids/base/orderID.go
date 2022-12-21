package base

import (
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

//type orderID struct {
//	ids.HashID
//}

var _ ids.OrderID = (*OrderID)(nil)

func (orderID *OrderID) Bytes() []byte {
	return orderID.OrderId.IdBytes
}
func (orderID *OrderID) IsOrderID() {}
func (orderID *OrderID) Compare(listable traits.Listable) int {
	return orderID.OrderId.Compare(orderIDFromInterface(listable).OrderId)
}
func (orderID *OrderID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_OrderId{
			OrderId: orderID,
		},
	}
}

func orderIDFromInterface(i interface{}) *OrderID {
	switch value := i.(type) {
	case *OrderID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewOrderID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.OrderID {
	return &OrderID{
		OrderId: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeOrderID() ids.OrderID {
	return &OrderID{
		OrderId: PrototypeHashID().(*HashID),
	}
}

func ReadOrderID(orderIDString string) (ids.OrderID, error) {
	if hashID, err := ReadHashID(orderIDString); err == nil {
		return &OrderID{
			OrderId: hashID.(*HashID),
		}, nil
	}

	if orderIDString == "" {
		return PrototypeOrderID(), nil
	}

	return &OrderID{}, nil
}
