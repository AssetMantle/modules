package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

type identityID struct {
	ChainID          types.ID `json:"chainId" valid:"required~Enter the ChainID"`
	MaintainersID    types.ID `json:"maintainersId" valid:"required~Enter the MaintainersID"`
	ClassificationID types.ID `json:"classificationId" valid:"required~Enter the ClassificationID"`
	HashID           types.ID `json:"hashId" valid:"required~Enter the HashID"`
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
			ChainID:          base.NewID(idList[0]),
			MaintainersID:    base.NewID(idList[1]),
			ClassificationID: base.NewID(idList[2]),
			HashID:           base.NewID(idList[3]),
		}
	}
	return identityID{ChainID: base.NewID(""), MaintainersID: base.NewID(""), ClassificationID: base.NewID(""), HashID: base.NewID("")}
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
