package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

var _ ids.SplitID = (*ID_SplitID)(nil)

func (splitID *ID_SplitID) String() string {
	return splitID.SplitID.String()
}
func (splitID *ID_SplitID) IsSplitID() {}
func (splitID *ID_SplitID) Bytes() []byte {
	return append(
		splitID.SplitID.OwnerID.HashID.IdBytes,
		splitID.SplitID.OwnableID.StringID.IdString...)
}
func (splitID *ID_SplitID) Compare(listable traits.Listable) int {
	return bytes.Compare(splitID.Bytes(), idFromInterface(listable).Bytes())
}
func (splitID *ID_SplitID) GetOwnableID() ids.ID {
	return &ID{Impl: &ID_OwnableID{OwnableID: splitID.SplitID.OwnableID}}
}

func GenerateSplitID(split types.Split) ids.ID {
	return NewSplitID(split.GetOwnerID(), split.GetOwnableID())
}
func NewSplitID(ownerID ids.ID, ownableID ids.ID) ids.ID {
	if ownerID.(*ID).GetIdentityID() == nil || ownableID.(*ID).GetOwnableID() == nil {
		panic(constants.MetaDataError)
	}
	return &ID{
		Impl: &ID_SplitID{
			SplitID: &SplitID{
				OwnerID:   ownerID.(*ID).GetIdentityID(),
				OwnableID: ownableID.(*ID).GetOwnableID(),
			},
		},
	}
}
func PrototypeSplitID() ids.ID {
	return NewSplitID(PrototypeIdentityID(), PrototypeOwnableID())
}

func ReadSplitID(splitIDString string) (ids.ID, error) {
	if splitIDStringSplit := stringUtilities.SplitCompositeIDString(splitIDString); len(splitIDStringSplit) == 2 {
		if ownerID, err := ReadIdentityID(splitIDStringSplit[0]); err == nil {
			if ownableID, err := ReadOwnableID(splitIDStringSplit[1]); err == nil {
				return NewSplitID(ownerID, ownableID), nil
			}
		}
	}

	if splitIDString == "" {
		return PrototypeSplitID(), nil
	}

	return PrototypeSplitID(), constants.MetaDataError
}
