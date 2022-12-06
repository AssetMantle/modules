package base

import (
	"bytes"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// type splitID struct {
//	OwnerID ids.IdentityID
//	ids.OwnableID
// }

type splitID base.SplitID

func (splitID *splitID) String() string {
	// TODO implement me
	panic("implement me")
}

var _ ids.SplitID = (*splitID)(nil)

func (splitID *splitID) IsSplitID() {}
func (splitID *splitID) Bytes() []byte {
	return append(
		splitID.OwnerId.Bytes(),
		splitID.OwnableId.Bytes()...)
}

// func (splitID *splitID) String() string {
//	return stringUtilities.JoinIDStrings(splitID.OwnerID.String(), splitID.OwnableID.String())
// }
func (splitID *splitID) Compare(listable traits.Listable) int {
	return bytes.Compare(splitID.Bytes(), splitIDFromInterface(listable).Bytes())
}
func (splitID *splitID) GetOwnableID() ids.ID {
	return splitID.OwnableId
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
	return &splitID{
		OwnerId:   ownerID.(*identityID),
		OwnableId: ownableID.(*ownableID),
	}
}

func PrototypeSplitID() ids.SplitID {
	return &splitID{
		OwnerId:   PrototypeIdentityID().(*identityID),
		OwnableId: PrototypeOwnableID().(*ownableID),
	}
}

func ReadSplitID(splitIDString string) (ids.SplitID, error) {
	if splitIDStringSplit := stringUtilities.SplitCompositeIDString(splitIDString); len(splitIDStringSplit) == 2 {
		if ownerID, err := ReadIdentityID(splitIDStringSplit[0]); err == nil {
			if ownableID, err := ReadOwnableID(splitIDStringSplit[1]); err == nil {
				return &splitID{
					OwnerId:   ownerID.(*identityID),
					OwnableId: ownableID.(*ownableID),
				}, nil
			}
		}
	}

	if splitIDString == "" {
		return PrototypeSplitID(), nil
	}

	return &splitID{}, constants.MetaDataError
}
