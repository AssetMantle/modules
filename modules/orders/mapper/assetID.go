package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"strings"
)

//var _ types.ID = (*orderID)(nil)

type orderID struct {
	ChainID          types.ID
	MaintainersID    types.ID
	ClassificationID types.ID
	HashID           types.ID
}

func (assetID orderID) Bytes() []byte {
	return append(append(append(
		assetID.ChainID.Bytes(),
		assetID.MaintainersID.Bytes()...),
		assetID.ClassificationID.Bytes()...),
		assetID.HashID.Bytes()...)
}

func (assetID orderID) String() string {
	var values []string
	values = append(values, assetID.ChainID.String())
	values = append(values, assetID.MaintainersID.String())
	values = append(values, assetID.ClassificationID.String())
	values = append(values, assetID.HashID.String())
	return strings.Join(values, constants.IDSeparator)
}

func (assetID orderID) Compare(id types.ID) int {
	return bytes.Compare(assetID.Bytes(), id.Bytes())
}

func readAssetID(assetIDString string) types.ID {
	idList := strings.Split(assetIDString, constants.IDSeparator)
	if len(idList) == 4 {
		return orderID{
			ChainID:          types.NewID(idList[0]),
			MaintainersID:    types.NewID(idList[1]),
			ClassificationID: types.NewID(idList[2]),
			HashID:           types.NewID(idList[3]),
		}
	}
	return orderID{ChainID: types.NewID(""), MaintainersID: types.NewID(""), ClassificationID: types.NewID(""), HashID: types.NewID("")}
}

func orderIDFromInterface(id types.ID) orderID {
	switch value := id.(type) {
	case orderID:
		return value
	default:
		return orderIDFromInterface(readAssetID(id.String()))
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
