package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"strings"
)

type identityID struct {
	ChainID          types.ID
	MaintainersID    types.ID
	ClassificationID types.ID
	HashID           types.ID
}

var _ types.ID = (*identityID)(nil)

func (identityID identityID) Bytes() []byte {
	return append(append(append(
		identityID.ChainID.Bytes(),
		identityID.MaintainersID.Bytes()...),
		identityID.ClassificationID.Bytes()...),
		identityID.HashID.Bytes()...)
}

func (identityID identityID) String() string {
	var values []string
	values = append(values, identityID.ChainID.String())
	values = append(values, identityID.MaintainersID.String())
	values = append(values, identityID.ClassificationID.String())
	values = append(values, identityID.HashID.String())
	return strings.Join(values, constants.IDSeparator)
}

func (identityID identityID) Compare(id types.ID) int {
	return bytes.Compare(identityID.Bytes(), id.Bytes())
}

func readIdentityID(identityIDString string) types.ID {
	idList := strings.Split(identityIDString, constants.IDSeparator)
	if len(idList) == 4 {
		return identityID{
			ChainID:          types.NewID(idList[0]),
			MaintainersID:    types.NewID(idList[1]),
			ClassificationID: types.NewID(idList[2]),
			HashID:           types.NewID(idList[3]),
		}
	}
	return identityID{ChainID: types.NewID(""), MaintainersID: types.NewID(""), ClassificationID: types.NewID(""), HashID: types.NewID("")}
}

func identityIDFromInterface(id types.ID) identityID {
	switch value := id.(type) {
	case identityID:
		return value
	default:
		return identityIDFromInterface(readIdentityID(id.String()))
	}
}

func NewIdentityID(chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) types.ID {
	return identityID{
		ChainID:          chainID,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		HashID:           hashID,
	}
}
