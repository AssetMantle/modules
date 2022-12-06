package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

var _ ids.SplitID = (*SplitIDI_SplitID)(nil)

func (splitID *SplitIDI_SplitID) String() string {
	return splitID.SplitID.String()
}
func (splitID *SplitIDI_SplitID) IsSplitID() {}
func (splitID *SplitIDI_SplitID) Bytes() []byte {
	return append(
		splitID.SplitID.OwnerId.Bytes(),
		splitID.SplitID.OwnableId.Bytes()...)
}
func (splitID *SplitIDI_SplitID) Compare(listable traits.Listable) int {
	return bytes.Compare(splitID.Bytes(), splitIDFromInterface(listable).Bytes())
}
func (splitID *SplitIDI_SplitID) GetOwnableID() ids.ID {
	return splitID.SplitID.OwnableId
}
func splitIDFromInterface(i interface{}) *SplitIDI_SplitID {
	switch value := i.(type) {
	case *SplitIDI_SplitID:
		return value
	default:
		panic(i)
	}
}

func GenerateSplitID(split types.Split) ids.SplitID {
	return NewSplitID(split.GetOwnerID(), split.GetOwnableID())
}
func NewSplitID(ownerID ids.IdentityID, ownableID ids.OwnableID) ids.SplitID {
	return &SplitIDI{
		Impl: &SplitIDI_SplitID{
			SplitID: &SplitID{
				OwnerId:   ownerID.(*IdentityIDI),
				OwnableId: ownableID.(*OwnableIDI),
			},
		},
	}
}
func PrototypeSplitID() ids.SplitID {
	return NewSplitID(PrototypeIdentityID(), PrototypeOwnableID())
}

func ReadSplitID(splitIDString string) (ids.SplitID, error) {
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
