package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

// TODO rename to something more appropriate
type ownableID struct {
	ids.StringID
}

var _ ids.OwnableID = (*ownableID)(nil)

func (ownableID ownableID) IsOwnableID() {}
func (ownableID ownableID) Compare(listable traits.Listable) int {
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
	return ownableID{
		StringID: stringID,
	}
}

func PrototypeOwnableID() ids.OwnableID {
	return ownableID{
		StringID: PrototypeStringID(),
	}
}

func ReadOwnableID(ownableIDString string) (ids.OwnableID, error) {
	// TODO ***** never allow ownable ID to be valid hash string
	if assetID, err := ReadAssetID(ownableIDString); err == nil {
		return assetID, nil
	}

	return ownableID{
		StringID: NewStringID(ownableIDString),
	}, nil
}
