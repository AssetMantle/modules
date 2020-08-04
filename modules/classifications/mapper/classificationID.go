package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

type classificationID struct {
	ChainID       types.ID `json:"chainID" valid:"required~required field chainID missing"`
	MaintainersID types.ID `json:"maintainersID" valid:"required~required field maintainersID missing"`
	HashID        types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*classificationID)(nil)

func (classificationID classificationID) Bytes() []byte {
	return append(append(
		classificationID.ChainID.Bytes(),
		classificationID.MaintainersID.Bytes()...),
		classificationID.HashID.Bytes()...)
}

func (classificationID classificationID) String() string {
	var values []string
	values = append(values, classificationID.ChainID.String())
	values = append(values, classificationID.MaintainersID.String())
	values = append(values, classificationID.HashID.String())
	return strings.Join(values, constants.IDSeparator)
}

func (classificationID classificationID) Compare(id types.ID) int {
	return bytes.Compare(classificationID.Bytes(), id.Bytes())
}

func readClassificationID(classificationIDString string) types.ID {
	idList := strings.Split(classificationIDString, constants.IDSeparator)
	if len(idList) == 4 {
		return classificationID{
			ChainID:       base.NewID(idList[0]),
			MaintainersID: base.NewID(idList[1]),
			HashID:        base.NewID(idList[3]),
		}
	}
	return classificationID{ChainID: base.NewID(""), MaintainersID: base.NewID(""), HashID: base.NewID("")}
}

func classificationIDFromInterface(id types.ID) classificationID {
	switch value := id.(type) {
	case classificationID:
		return value
	default:
		return classificationIDFromInterface(readClassificationID(id.String()))
	}
}
func generateKey(classificationID types.ID) []byte {
	return append(StoreKeyPrefix, classificationIDFromInterface(classificationID).Bytes()...)
}
func NewClassificationID(chainID types.ID, maintainersID types.ID, hashID types.ID) types.ID {
	return classificationID{
		ChainID:       chainID,
		MaintainersID: maintainersID,
		HashID:        hashID,
	}
}
