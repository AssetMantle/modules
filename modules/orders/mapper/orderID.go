package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

type orderID struct {
	ChainID          types.ID
	MaintainersID    types.ID
	ClassificationID types.ID
	HashID           types.ID
}

var _ types.ID = (*orderID)(nil)

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

func (orderID orderID) Compare(id types.ID) int {
	return bytes.Compare(orderID.Bytes(), id.Bytes())
}

func readOrderID(orderIDString string) types.ID {
	idList := strings.Split(orderIDString, constants.IDSeparator)
	if len(idList) == 4 {
		return orderID{
			ChainID:          base.NewID(idList[0]),
			MaintainersID:    base.NewID(idList[1]),
			ClassificationID: base.NewID(idList[2]),
			HashID:           base.NewID(idList[3]),
		}
	}
	return orderID{ChainID: base.NewID(""), MaintainersID: base.NewID(""), ClassificationID: base.NewID(""), HashID: base.NewID("")}
}

func orderIDFromInterface(id types.ID) orderID {
	switch value := id.(type) {
	case orderID:
		return value
	default:
		return orderIDFromInterface(readOrderID(id.String()))
	}
}

func NewOrderID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) types.ID {
	return orderID{
		ChainID:          chainID,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		HashID:           hashID,
	}
}
