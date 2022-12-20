package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

//type splitID struct {
//	OwnerID   ids.IdentityID
//	OwnableID ids.OwnableID
//}

var _ ids.SplitID = (*SplitID)(nil)

func (splitID *SplitID) IsSplitID() {}
func (splitID *SplitID) Bytes() []byte {
	return append(
		splitID.OwnerID.Bytes(),
		splitID.OwnableID.Bytes()...)
}
func (splitID *SplitID) SplitIDString() string {
	return stringUtilities.JoinIDStrings(splitID.OwnerID.String(), splitID.OwnableID.String())
}
func (splitID *SplitID) Compare(listable traits.Listable) int {
	return bytes.Compare(splitID.Bytes(), splitIDFromInterface(listable).Bytes())
}
func (splitID *SplitID) ToAnyID() *AnyID {
	return &AnyID{
		Impl: &AnyID_SplitID{
			SplitID: splitID,
		},
	}
}

func splitIDFromInterface(i interface{}) *SplitID {
	switch value := i.(type) {
	case *SplitID:
		return value
	default:
		panic(i)
	}
}

func NewSplitID(ownerID ids.IdentityID, ownableID ids.OwnableID) ids.SplitID {
	return &SplitID{
		OwnerID:   ownerID.(*IdentityID),
		OwnableID: ownableID.(*OwnableID),
	}
}

func PrototypeSplitID() ids.SplitID {
	return &SplitID{
		OwnerID:   PrototypeIdentityID().(*IdentityID),
		OwnableID: PrototypeOwnableID().(*OwnableID),
	}
}

func ReadSplitID(splitIDString string) (ids.SplitID, error) {
	if splitIDStringSplit := stringUtilities.SplitCompositeIDString(splitIDString); len(splitIDStringSplit) == 2 {
		if ownerID, err := ReadIdentityID(splitIDStringSplit[0]); err == nil {
			if ownableID, err := ReadOwnableID(splitIDStringSplit[1]); err == nil {
				return &SplitID{
					OwnerID:   ownerID.(*IdentityID),
					OwnableID: ownableID.(*OwnableID),
				}, nil
			}
		}
	}

	if splitIDString == "" {
		return PrototypeSplitID(), nil
	}

	return &SplitID{}, constants.MetaDataError
}
