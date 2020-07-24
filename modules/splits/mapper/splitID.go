package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

type splitID struct {
	OwnerID   types.ID `json:"OwnerId" valid:"required~required field ownerID missing"`
	OwnableID types.ID `json:"OwnableId" valid:"required~required field ownableid missing"`
}

var _ types.ID = (*splitID)(nil)

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

func (splitID splitID) Compare(id types.ID) int {
	return bytes.Compare(splitID.Bytes(), id.Bytes())
}

func readSplitID(splitIDString string) types.ID {
	idList := strings.Split(splitIDString, constants.IDSeparator)
	if len(idList) == 2 {
		return splitID{
			OwnerID:   base.NewID(idList[0]),
			OwnableID: base.NewID(idList[1]),
		}
	}
	return splitID{OwnerID: base.NewID(""), OwnableID: base.NewID("")}
}

func splitIDFromInterface(id types.ID) splitID {
	switch value := id.(type) {
	case splitID:
		return value
	default:
		return splitIDFromInterface(readSplitID(id.String()))
	}
}
func generateKey(splitID types.ID) []byte {
	return append(StoreKeyPrefix, splitIDFromInterface(splitID).Bytes()...)
}
func NewSplitID(ownerID types.ID, ownableID types.ID) types.ID {
	return splitID{
		OwnerID:   ownerID,
		OwnableID: ownableID,
	}
}
