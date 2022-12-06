package base

import (
	"bytes"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

type maintainerID base.MaintainerID

var _ ids.MaintainerID = (*maintainerID)(nil)

func (maintainerID *maintainerID) IsMaintainerID() {}
func (maintainerID *maintainerID) String() string {
	return maintainerID.HashId.String()
}
func (maintainerID *maintainerID) Bytes() []byte {
	return maintainerID.HashId.Bytes()
}
func (maintainerID *maintainerID) Compare(listable traits.Listable) int {
	return bytes.Compare(maintainerID.Bytes(), maintainerIDFromInterface(listable).Bytes())
}

func maintainerIDFromInterface(i interface{}) *hashID {
	switch value := i.(type) {
	case hashID:
		return &value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func NewMaintainerID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.MaintainerID {
	return &maintainerID{
		HashId: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*hashID),
	}
}

func PrototypeMaintainerID() ids.MaintainerID {
	return &maintainerID{
		HashId: PrototypeHashID().(*hashID),
	}
}

func ReadMaintainerID(maintainerIDString string) (ids.MaintainerID, error) {
	if hashID, err := ReadHashID(maintainerIDString); err == nil {
		return &maintainerID{
			HashId: hashID.(*hashID),
		}, nil
	}

	if maintainerIDString == "" {
		return PrototypeMaintainerID(), nil
	}

	return &maintainerID{}, errorConstants.MetaDataError
}
