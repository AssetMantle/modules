package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

//type maintainerID struct {
//	ids.HashID
//}

var _ ids.MaintainerID = (*MaintainerID)(nil)

func (maintainerID *MaintainerID) Bytes() []byte {
	return maintainerID.HashID.IdBytes
}
func (maintainerID *MaintainerID) IsMaintainerID() {}
func (maintainerID *MaintainerID) Compare(listable traits.Listable) int {
	return maintainerID.HashID.Compare(maintainerIDFromInterface(listable).HashID)
}
func (maintainerID *MaintainerID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_MaintainerID{
			MaintainerID: maintainerID,
		},
	}
}

func maintainerIDFromInterface(i interface{}) *MaintainerID {
	switch value := i.(type) {
	case *MaintainerID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func NewMaintainerID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.MaintainerID {
	return &MaintainerID{
		HashID: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeMaintainerID() ids.MaintainerID {
	return &MaintainerID{
		HashID: PrototypeHashID().(*HashID),
	}
}

func ReadMaintainerID(maintainerIDString string) (ids.MaintainerID, error) {
	if hashID, err := ReadHashID(maintainerIDString); err == nil {
		return &MaintainerID{
			HashID: hashID.(*HashID),
		}, nil
	}

	if maintainerIDString == "" {
		return PrototypeMaintainerID(), nil
	}

	return &MaintainerID{}, errorConstants.MetaDataError
}
