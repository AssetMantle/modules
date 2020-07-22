package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"

	"github.com/persistenceOne/persistenceSDK/types"
	"strings"
)

type orderID struct {
	ChainID          schema.ID
	MaintainersID    schema.ID
	ClassificationID schema.ID
	HashID           schema.ID
}

var _ schema.ID = (*orderID)(nil)

func (orderID orderID) Bytes() []byte {
	return append(append(append(
		orderID.ChainID.Bytes(),
		orderID.MaintainersID.Bytes()...),
		orderID.ClassificationID.Bytes()...),
		orderID.HashID.Bytes()...)
}

func (orderID orderID) String() string {
	var values []string
	values = append(values, orderID.ChainID.String())
	values = append(values, orderID.MaintainersID.String())
	values = append(values, orderID.ClassificationID.String())
	values = append(values, orderID.HashID.String())
	return strings.Join(values, constants.IDSeparator)
}

func (orderID orderID) Compare(id schema.ID) int {
	return bytes.Compare(orderID.Bytes(), id.Bytes())
}

func readOrderID(orderIDString string) schema.ID {
	idList := strings.Split(orderIDString, constants.IDSeparator)
	if len(idList) == 4 {
		return orderID{
			ChainID:          schema.NewID(idList[0]),
			MaintainersID:    schema.NewID(idList[1]),
			ClassificationID: schema.NewID(idList[2]),
			HashID:           schema.NewID(idList[3]),
		}
	}
	return orderID{ChainID: schema.NewID(""), MaintainersID: schema.NewID(""), ClassificationID: schema.NewID(""), HashID: schema.NewID("")}
}

func orderIDFromInterface(id schema.ID) orderID {
	switch value := id.(type) {
	case orderID:
		return value
	default:
		return orderIDFromInterface(readOrderID(id.String()))
	}
}

func NewOrderID(chainID schema.ID, maintainersID schema.ID, classificationID schema.ID, hashID schema.ID) schema.ID {
	return orderID{
		ChainID:          chainID,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		HashID:           hashID,
	}
}
