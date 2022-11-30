package base

import (
	"bytes"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.OrderID = (*OrderID)(nil)

func (orderID *OrderID) Compare(listable traits.Listable) int {
	return bytes.Compare(orderID.Bytes(), orderIDFromInterface(listable).Bytes())
}

func (orderID *OrderID) Bytes() []byte {
	return *orderID.OrderId
}

func (orderID OrderID) IsOrderID() {}

func orderIDFromInterface(i interface{}) *OrderID {
	switch value := i.(type) {
	case OrderID:
		return &value
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
