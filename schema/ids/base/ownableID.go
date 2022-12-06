package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.OwnableID = (*OwnableIDI_OwnableID)(nil)

func (ownableID *OwnableIDI_OwnableID) String() string {
	return ownableID.OwnableID.StringId.String()
}
func (ownableID *OwnableIDI_OwnableID) Bytes() []byte {
	return ownableID.OwnableID.StringId.Bytes()
}
func (ownableID *OwnableIDI_OwnableID) IsOwnableID() {}
func (ownableID *OwnableIDI_OwnableID) Compare(listable traits.Listable) int {
	// TODO devise a better strategy to compare assetID and ownableID
	return bytes.Compare(ownableID.Bytes(), ownableIDFromInterface(listable).Bytes())
}
func ownableIDFromInterface(i interface{}) ids.OwnableID {
	switch value := i.(type) {
	case ids.OwnableID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}
func NewOwnableID(stringID ids.StringID) ids.OwnableID {
	return &OwnableIDI{
		Impl: &OwnableIDI_OwnableID{
			OwnableID: &OwnableID{
				StringId: stringID.(*StringIDI),
			},
		},
	}
}

func PrototypeOwnableID() ids.OwnableID {
	return NewOwnableID(PrototypeStringID())
}

func ReadOwnableID(ownableIDString string) (ids.OwnableID, error) {
	// TODO ***** never allow ownable PropertyID to be valid hash string
	if assetID, err := ReadAssetID(ownableIDString); err == nil {
		return assetID, nil
	}

	return NewOwnableID(NewStringID(ownableIDString)), nil
}
