package base

import (
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

type orderID struct {
	ids.HashID
}

var _ ids.OrderID = (*orderID)(nil)

func (orderID orderID) IsOrderID() {}
func (orderID orderID) Compare(listable traits.Listable) int {
	return orderID.HashID.Compare(orderIDFromInterface(listable).HashID)
}
func orderIDFromInterface(i interface{}) orderID {
	switch value := i.(type) {
	case orderID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewOrderID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.OrderID {
	return orderID{
		HashID: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()),
	}
}

func PrototypeOrderID() ids.OrderID {
	return orderID{
		HashID: PrototypeHashID(),
	}
}

func ReadOrderID(orderIDString string) (ids.OrderID, error) {
	if hashID, err := ReadHashID(orderIDString); err == nil {
		return orderID{
			HashID: hashID,
		}, nil
	}

	if orderIDString == "" {
		return PrototypeOrderID(), nil
	}

	return orderID{}, nil
}
