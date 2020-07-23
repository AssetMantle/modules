package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"strings"
)

type identityID struct {
	ChainID          schema.ID `json:"chain id" valid:"required~Enter the ChainID"`
	MaintainersID    schema.ID `json:"maintainers id" valid:"required~Enter the MaintainersID"`
	ClassificationID schema.ID `json:"classification ID" valid:"required~Enter the ClassificationID"`
	HashID           schema.ID `json:"hash id" valid:"required~Enter the HashID"`
}

var _ schema.ID = (*identityID)(nil)

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

func (identityID identityID) Compare(id schema.ID) int {
	return bytes.Compare(identityID.Bytes(), id.Bytes())
}

func readIdentityID(identityIDString string) schema.ID {
	idList := strings.Split(identityIDString, constants.IDSeparator)
	if len(idList) == 4 {
		return identityID{
			ChainID:          schema.NewID(idList[0]),
			MaintainersID:    schema.NewID(idList[1]),
			ClassificationID: schema.NewID(idList[2]),
			HashID:           schema.NewID(idList[3]),
		}
	}
	return identityID{ChainID: schema.NewID(""), MaintainersID: schema.NewID(""), ClassificationID: schema.NewID(""), HashID: schema.NewID("")}
}

func identityIDFromInterface(id schema.ID) identityID {
	switch value := id.(type) {
	case identityID:
		return value
	default:
		return identityIDFromInterface(readIdentityID(id.String()))
	}
}

func NewIdentityID(chainID schema.ID, maintainersID schema.ID, classificationID schema.ID, hashID schema.ID) schema.ID {
	return identityID{
		ChainID:          chainID,
		MaintainersID:    maintainersID,
		ClassificationID: classificationID,
		HashID:           hashID,
	}
}
