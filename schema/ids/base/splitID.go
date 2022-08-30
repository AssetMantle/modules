package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

type splitID struct {
	OwnerID   ids.IdentityID
	OwnableID ids.OwnableID
}

func (splitID splitID) IsSplitID() {
	// TODO implement me
	panic("implement me")
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

func NewSplitID(ownerID ids.IdentityID, ownableID ids.OwnableID) ids.SplitID {
	return splitID{
		OwnerID:   ownerID,
		OwnableID: ownableID,
	}
}

func ReadSplitID(splitIDString string) (ids.SplitID, error) {
	if splitIDStringSplit := stringUtilities.SplitCompositeIDString(splitIDString); len(splitIDStringSplit) == 3 {
		if ownerID, err := ReadIdentityID(stringUtilities.JoinIDStrings(stringUtilities.SplitCompositeIDString(splitIDString)[:1]...)); err == nil {
			if ownableID, err := ReadOwnableID(stringUtilities.SplitCompositeIDString(splitIDString)[2]); err == nil {
				return splitID{
					OwnerID:   ownerID,
					OwnableID: ownableID,
				}, nil
			}
		}
	}
	return splitID{}, constants.MetaDataError
}
