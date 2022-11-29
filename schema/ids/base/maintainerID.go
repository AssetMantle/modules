package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ ids.MaintainerID = (*HashID)(nil)

func (maintainerID HashID) IsMaintainerID() {}

func maintainerIDFromInterface(i interface{}) *HashID {
	switch value := i.(type) {
	case HashID:
		return &value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func NewMaintainerID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.MaintainerID {
	return &HashID{
		HashBytes: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).Bytes(),
	}
}

func PrototypeMaintainerID() ids.MaintainerID {
	return &HashID{
		HashBytes: PrototypeHashID().Bytes(),
	}
}

func ReadMaintainerID(maintainerIDString string) (ids.MaintainerID, error) {
	if hashID, err := ReadHashID(maintainerIDString); err == nil {
		return &HashID{
			HashBytes: hashID.Bytes(),
		}, nil
	}

	if maintainerIDString == "" {
		return PrototypeMaintainerID(), nil
	}

	return &HashID{}, errorConstants.MetaDataError
}
