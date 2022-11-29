package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.MaintainerID = (*HashUser)(nil)

func (maintainerID HashUser) Bytes() []byte {
	return maintainerID.HashId.HashBytes
}
func (maintainerID HashUser) IsMaintainerID() {}
func (maintainerID HashUser) Compare(listable traits.Listable) int {
	return maintainerID.HashId.Compare(maintainerIDFromInterface(listable).HashId)
}
func maintainerIDFromInterface(i interface{}) *HashUser {
	switch value := i.(type) {
	case HashUser:
		return &value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func NewMaintainerID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.MaintainerID {
	return &HashUser{
		HashId: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeMaintainerID() ids.MaintainerID {
	return &HashUser{
		HashId: PrototypeHashID().(*HashID),
	}
}

func ReadMaintainerID(maintainerIDString string) (ids.MaintainerID, error) {
	if hashID, err := ReadHashID(maintainerIDString); err == nil {
		return &HashUser{
			HashId: hashID.(*HashID),
		}, nil
	}

	if maintainerIDString == "" {
		return PrototypeMaintainerID(), nil
	}

	return &HashUser{}, errorConstants.MetaDataError
}
