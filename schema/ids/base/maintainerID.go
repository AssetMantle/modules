package base

import (
	"bytes"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.MaintainerID = (*MaintainerID)(nil)

func (maintainerID MaintainerID) IsMaintainerID() {}

func (maintainerID MaintainerID) Bytes() []byte {
	return *maintainerID.HashId
}

func (maintainerID MaintainerID) Compare(listable traits.Listable) int {
	return bytes.Compare(maintainerID.Bytes(), maintainerIDFromInterface(listable).Bytes())
}

func maintainerIDFromInterface(i interface{}) *HashID {
	switch value := i.(type) {
	case HashID:
		return &value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func NewMaintainerID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.MaintainerID {
	return &MaintainerID{
		HashId: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeMaintainerID() ids.MaintainerID {
	return &MaintainerID{
		HashId: PrototypeHashID().(*HashID),
	}
}

func ReadMaintainerID(maintainerIDString string) (ids.MaintainerID, error) {
	if hashID, err := ReadHashID(maintainerIDString); err == nil {
		return &MaintainerID{
			HashId: hashID.(*HashID),
		}, nil
	}

	if maintainerIDString == "" {
		return PrototypeMaintainerID(), nil
	}

	return &MaintainerID{}, errorConstants.MetaDataError
}
