package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"strings"
)

type splitID struct {
	OwnerID   schema.ID `json:"OwnerId" valid:"required~Enter the OwnerID"`
	OwnableID schema.ID `json:"OwnableId" valid:"required~Enter the OwnableID"`
}

var _ schema.ID = (*splitID)(nil)

func (splitID splitID) Bytes() []byte {
	return append(
		splitID.OwnerID.Bytes(),
		splitID.OwnableID.Bytes()...)

}

func (splitID splitID) String() string {
	var values []string
	values = append(values, splitID.OwnerID.String())
	values = append(values, splitID.OwnableID.String())
	return strings.Join(values, constants.IDSeparator)
}

func (splitID splitID) Compare(id schema.ID) int {
	return bytes.Compare(splitID.Bytes(), id.Bytes())
}

func readSplitID(splitIDString string) schema.ID {
	idList := strings.Split(splitIDString, constants.IDSeparator)
	if len(idList) == 2 {
		return splitID{
			OwnerID:   schema.NewID(idList[0]),
			OwnableID: schema.NewID(idList[1]),
		}
	}
	return splitID{OwnerID: schema.NewID(""), OwnableID: schema.NewID("")}
}

func splitIDFromInterface(id schema.ID) splitID {
	switch value := id.(type) {
	case splitID:
		return value
	default:
		return splitIDFromInterface(readSplitID(id.String()))
	}
}

func NewSplitID(ownerID schema.ID, ownableID schema.ID) schema.ID {
	return splitID{
		OwnerID:   ownerID,
		OwnableID: ownableID,
	}
}
