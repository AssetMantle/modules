package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.MaintainerID = (*MaintainerIDI_MaintainerID)(nil)

func (maintainerID *MaintainerIDI_MaintainerID) IsMaintainerID() {}
func (maintainerID *MaintainerIDI_MaintainerID) String() string {
	return maintainerID.MaintainerID.String()
}
func (maintainerID *MaintainerIDI_MaintainerID) Bytes() []byte {
	return maintainerID.MaintainerID.HashId.Bytes()
}
func (maintainerID *MaintainerIDI_MaintainerID) Compare(listable traits.Listable) int {
	return bytes.Compare(maintainerID.Bytes(), maintainerIDFromInterface(listable).Bytes())
}

func maintainerIDFromInterface(i interface{}) *MaintainerIDI {
	switch value := i.(type) {
	case *MaintainerIDI:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func GenerateMaintainerID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.MaintainerID {
	return NewMaintainerID(GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()))
}
func NewMaintainerID(hashID ids.HashID) ids.MaintainerID {
	return &MaintainerIDI{
		Impl: &MaintainerIDI_MaintainerID{
			MaintainerID: &MaintainerID{
				HashId: hashID.(*HashIDI),
			},
		},
	}
}

func PrototypeMaintainerID() ids.MaintainerID {
	return NewMaintainerID(PrototypeHashID())
}

func ReadMaintainerID(maintainerIDString string) (ids.MaintainerID, error) {
	if hashID, err := ReadHashID(maintainerIDString); err == nil {
		return NewMaintainerID(hashID), nil
	}

	if maintainerIDString == "" {
		return PrototypeMaintainerID(), nil
	}

	return PrototypeMaintainerID(), errorConstants.MetaDataError
}
