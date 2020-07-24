package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"

	"strings"
)

type assetID struct {
	ChainID          types.ID `json:"chainid" valid:"required~required field chainid missing"`
	MaintainersID    types.ID `json:"maintainersid" valid:"required~required field maintainersid missing"`
	ClassificationID types.ID `json:"classificationid" valid:"required~required field classificationid missing"`
	HashID           types.ID `json:"hashid" valid:"required~required field hashid missing"`
}

var _ types.ID = (*assetID)(nil)

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

func (assetID assetID) Compare(id types.ID) int {
	return bytes.Compare(assetID.Bytes(), id.Bytes())
}

func readAssetID(assetIDString string) types.ID {
	idList := strings.Split(assetIDString, constants.IDSeparator)
	if len(idList) == 4 {
		return assetID{
			ChainID:          base.NewID(idList[0]),
			MaintainersID:    base.NewID(idList[1]),
			ClassificationID: base.NewID(idList[2]),
			HashID:           base.NewID(idList[3]),
		}
	}
	return assetID{ChainID: base.NewID(""), MaintainersID: base.NewID(""), ClassificationID: base.NewID(""), HashID: base.NewID("")}
}

func assetIDFromInterface(id types.ID) assetID {
	switch value := id.(type) {
	case assetID:
		return value
	default:
		return assetIDFromInterface(readAssetID(id.String()))
	}
}
func generateKey(assetID types.ID) []byte {
	return append(StoreKeyPrefix, assetIDFromInterface(assetID).Bytes()...)
}
func NewAssetID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) types.ID {
	return assetID{
		ChainID:          chainID,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		HashID:           hashID,
	}
}
