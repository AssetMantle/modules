package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

type splitID struct {
	OwnerID   ids.ID
	OwnableID ids.ID
}

var _ ids.SplitID = (*splitID)(nil)

func (splitID splitID) Bytes() []byte {
	return append(
		splitID.OwnerID.Bytes(),
		splitID.OwnableID.Bytes()...)
}
func (splitID splitID) String() string {
	return stringUtilities.JoinIDStrings(splitID.OwnerID.String(), splitID.OwnableID.String())
}
func (splitID splitID) Compare(listable traits.Listable) int {
	return bytes.Compare(splitID.Bytes(), splitIDFromInterface(listable).Bytes())
}
func (splitID splitID) GetOwnableID() ids.ID {
	return splitID.OwnableID
}
func splitIDFromInterface(i interface{}) splitID {
	switch value := i.(type) {
	case splitID:
		return value
	default:
		panic(i)
	}
}

func NewSplitID(ownerID ids.ID, ownableID ids.ID) ids.SplitID {
	return splitID{
		OwnerID:   ownerID,
		OwnableID: ownableID,
	}
}
