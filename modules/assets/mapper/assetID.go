package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"

	"strings"
)

type assetID struct {
	ChainID          schema.ID `json:"chainID" valid:"required~Enter the ChainID"`
	MaintainersID    schema.ID `json:"maintainersID" valid:"required~Enter the MaintainersID"`
	ClassificationID schema.ID `json:"classificationID" valid:"required~Enter the ClassificationID"`
	HashID           schema.ID `json:"hashID" valid:"required~Enter the HashID"`
}

var _ schema.ID = (*assetID)(nil)

func (assetID assetID) Bytes() []byte {
	return append(append(append(
		assetID.ChainID.Bytes(),
		assetID.MaintainersID.Bytes()...),
		assetID.ClassificationID.Bytes()...),
		assetID.HashID.Bytes()...)
}

func (assetID assetID) String() string {
	var values []string
	values = append(values, assetID.ChainID.String())
	values = append(values, assetID.MaintainersID.String())
	values = append(values, assetID.ClassificationID.String())
	values = append(values, assetID.HashID.String())
	return strings.Join(values, constants.IDSeparator)
}

func (assetID assetID) Compare(id schema.ID) int {
	return bytes.Compare(assetID.Bytes(), id.Bytes())
}

func readAssetID(assetIDString string) schema.ID {
	idList := strings.Split(assetIDString, constants.IDSeparator)
	if len(idList) == 4 {
		return assetID{
			ChainID:          schema.NewID(idList[0]),
			MaintainersID:    schema.NewID(idList[1]),
			ClassificationID: schema.NewID(idList[2]),
			HashID:           schema.NewID(idList[3]),
		}
	}
	return assetID{ChainID: schema.NewID(""), MaintainersID: schema.NewID(""), ClassificationID: schema.NewID(""), HashID: schema.NewID("")}
}

func assetIDFromInterface(id schema.ID) assetID {
	switch value := id.(type) {
	case assetID:
		return value
	default:
		return assetIDFromInterface(readAssetID(id.String()))
	}
}

func NewAssetID(chainID schema.ID, maintainersID schema.ID, classificationID schema.ID, hashID schema.ID) schema.ID {
	return assetID{
		ChainID:          chainID,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		HashID:           hashID,
	}
}
