package base

import (
	"bytes"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

type orderID base.OrderID

var _ ids.OrderID = (*orderID)(nil)

func (orderID *orderID) String() string {
	// TODO implement me
	panic("implement me")
}
func (orderID *orderID) Compare(listable traits.Listable) int {
	return bytes.Compare(orderID.Bytes(), orderIDFromInterface(listable).Bytes())
}

func (orderID *orderID) Bytes() []byte {
	return orderID.OrderId.Bytes()
}

func (orderID *orderID) IsOrderID() {}

func orderIDFromInterface(i interface{}) *orderID {
	switch value := i.(type) {
	case orderID:
		return &value
	default:
		panic(constants.MetaDataError)
	}
}

func NewOrderID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.OrderID {
	return &orderID{
		OrderId: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*hashID),
	}
}

func PrototypeOrderID() ids.OrderID {
	return &orderID{
		OrderId: PrototypeHashID().(*hashID),
	}
}

func ReadOrderID(orderIDString string) (ids.OrderID, error) {
	if hashID, err := ReadHashID(orderIDString); err == nil {
		return &orderID{
			OrderId: hashID.(*hashID),
		}, nil
	}

	if orderIDString == "" {
		return PrototypeOrderID(), nil
	}

	return &orderID{}, nil
}
