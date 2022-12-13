package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.OwnableID = (*ID_OwnableID)(nil)

func (ownableID *ID_OwnableID) String() string {
	return ownableID.OwnableID.StringId.String()
}
func (ownableID *ID_OwnableID) Bytes() []byte {
	return []byte(ownableID.OwnableID.StringId.IdString)
}
func (ownableID *ID_OwnableID) IsOwnableID() {}
func (ownableID *ID_OwnableID) Compare(listable traits.Listable) int {
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
func NewOwnableID(stringID ids.ID) ids.ID {
	if stringID.(*ID).GetStringID() == nil {
		panic(constants.MetaDataError)
	}
	return &ID{
		Impl: &ID_OwnableID{
			OwnableID: &OwnableID{
				StringId: stringID.(*ID).GetStringID(),
			},
		},
	}
}

func PrototypeOwnableID() ids.ID {
	return NewOwnableID(PrototypeStringID())
}

func ReadOwnableID(ownableIDString string) (ids.ID, error) {
	// TODO ***** never allow ownable PropertyID to be valid hash string
	if assetID, err := ReadAssetID(ownableIDString); err == nil {
		return assetID, nil
	}

	return NewOwnableID(NewStringID(ownableIDString)), nil
}
