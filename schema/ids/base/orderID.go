package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.OrderID = (*OrderID)(nil)

func (orderID *OrderID) AsString() string {
	return orderID.HashID.AsString()
}
func (orderID *OrderID) Bytes() []byte {
	return orderID.HashID.IDBytes
}
func (orderID *OrderID) IsOrderID() {}
func (orderID *OrderID) Compare(listable traits.Listable) int {
	return orderID.HashID.Compare(orderIDFromInterface(listable).HashID)
}
func (orderID *OrderID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_OrderID{
			OrderID: orderID,
		},
	}
}

func orderIDFromInterface(i interface{}) *OrderID {
	switch value := i.(type) {
	case *OrderID:
		return value
	default:
		panic(errorConstants.IncorrectFormat.Wrapf("expected *OrderID, got %T", i))
	}
}

func NewOrderID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.OrderID {
	return &OrderID{
		HashID: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeOrderID() ids.OrderID {
	return &OrderID{
		HashID: PrototypeHashID().(*HashID),
	}
}

func ReadOrderID(orderIDString string) (ids.OrderID, error) {
	if hashID, err := ReadHashID(orderIDString); err == nil {
		return &OrderID{
			HashID: hashID.(*HashID),
		}, nil
	}

	if orderIDString == "" {
		return PrototypeOrderID(), nil
	}

	return &OrderID{}, nil
}
