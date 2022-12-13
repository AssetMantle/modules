package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

type maintainerID struct {
	ids.HashID
}

var _ ids.MaintainerID = (*maintainerID)(nil)

func (maintainerID maintainerID) IsMaintainerID() {}
func (maintainerID maintainerID) Compare(listable traits.Listable) int {
	return maintainerID.HashID.Compare(maintainerIDFromInterface(listable).HashID)
}
func maintainerIDFromInterface(i interface{}) maintainerID {
	switch value := i.(type) {
	case maintainerID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func NewMaintainerID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.MaintainerID {
	return maintainerID{
		HashID: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()),
	}
}

func PrototypeMaintainerID() ids.MaintainerID {
	return maintainerID{
		HashID: PrototypeHashID(),
	}
}

func ReadMaintainerID(maintainerIDString string) (ids.MaintainerID, error) {
	if hashID, err := ReadHashID(maintainerIDString); err == nil {
		return maintainerID{
			HashID: hashID,
		}, nil
	}

	if maintainerIDString == "" {
		return PrototypeMaintainerID(), nil
	}

	return maintainerID{}, errorConstants.MetaDataError
}
