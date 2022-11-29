package base

import (
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ ids.OrderID = (*HashID)(nil)

func (orderID HashID) IsOrderID() {}

func orderIDFromInterface(i interface{}) *HashID {
	switch value := i.(type) {
	case HashID:
		return &value
	default:
		panic(constants.MetaDataError)
	}
}

func NewOrderID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.OrderID {
	return &HashID{
		HashBytes: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).Bytes(),
	}
}

func PrototypeOrderID() ids.OrderID {
	return &HashID{
		HashBytes: PrototypeHashID().Bytes(),
	}
}

func ReadOrderID(orderIDString string) (ids.OrderID, error) {
	if hashID, err := ReadHashID(orderIDString); err == nil {
		return &HashID{
			HashBytes: hashID.Bytes(),
		}, nil
	}

	if orderIDString == "" {
		return PrototypeOrderID(), nil
	}

	return &HashID{}, nil
}
