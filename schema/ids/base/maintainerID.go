package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.MaintainerID = (*ID_MaintainerID)(nil)

func (maintainerID *ID_MaintainerID) IsMaintainerID() {}
func (maintainerID *ID_MaintainerID) String() string {
	return maintainerID.MaintainerID.String()
}
func (maintainerID *ID_MaintainerID) Bytes() []byte {
	return maintainerID.MaintainerID.HashID.IdBytes
}
func (maintainerID *ID_MaintainerID) Compare(listable traits.Listable) int {
	return bytes.Compare(maintainerID.Bytes(), idFromInterface(listable).Bytes())
}
func NewMaintainerID(hashID ids.ID) ids.ID {
	if hashID.(*ID).GetHashID() == nil {
		panic(errorConstants.MetaDataError)
	}
	return &ID{
		Impl: &ID_MaintainerID{
			MaintainerID: &MaintainerID{
				HashID: hashID.(*ID).GetHashID(),
			},
		},
	}
}

func PrototypeMaintainerID() ids.ID {
	return NewMaintainerID(PrototypeHashID())
}

func ReadMaintainerID(maintainerIDString string) (ids.ID, error) {
	if hashID, err := ReadHashID(maintainerIDString); err == nil {
		return NewMaintainerID(hashID), nil
	}

	if maintainerIDString == "" {
		return PrototypeMaintainerID(), nil
	}

	return PrototypeMaintainerID(), errorConstants.MetaDataError
}
