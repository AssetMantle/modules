package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.OrderID = (*ID_OrderID)(nil)

func (orderID *ID_OrderID) String() string {
	return orderID.OrderID.String()
}
func (orderID *ID_OrderID) Compare(listable traits.Listable) int {
	return bytes.Compare(orderID.Bytes(), idFromInterface(listable).Bytes())
}
func (orderID *ID_OrderID) Bytes() []byte {
	return orderID.OrderID.OrderId.IdBytes
}
func (orderID *ID_OrderID) IsOrderID() {}

func NewOrderID(hashID ids.ID) ids.ID {
	if hashID.(*ID).GetHashID() == nil {
		panic(constants.MetaDataError)
	}
	return &ID{
		Impl: &ID_OrderID{
			OrderID: &OrderID{
				OrderId: hashID.(*ID).GetHashID(),
			},
		},
	}
}

func PrototypeOrderID() ids.ID {
	return NewOrderID(PrototypeHashID())
}

func ReadOrderID(orderIDString string) (ids.ID, error) {
	if hashID, err := ReadHashID(orderIDString); err == nil {
		return NewOrderID(hashID), nil
	}

	if orderIDString == "" {
		return PrototypeOrderID(), nil
	}

	return PrototypeOrderID(), nil
}
