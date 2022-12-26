package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

// TODO rename to something more appropriate
//type ownableID struct {
//	ids.StringID
//}

var _ ids.OwnableID = (*OwnableID)(nil)

//TODO: Verify
func (ownableID *OwnableID) Bytes() []byte {
	return []byte(ownableID.StringID.IdString)
}
func (ownableID *OwnableID) IsOwnableID() {}
func (ownableID *OwnableID) Compare(listable traits.Listable) int {
	// TODO devise a better strategy to compare ownableID and ownableID
	return bytes.Compare(ownableID.Bytes(), ownableIDFromInterface(listable).Bytes())
}
func (ownableID *OwnableID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_OwnableID{
			OwnableID: ownableID,
		},
	}
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
	return &OwnableID{
		StringID: stringID.(*StringID),
	}
}

func PrototypeOwnableID() ids.OwnableID {
	return &OwnableID{
		StringID: PrototypeStringID().(*StringID),
	}
}

func ReadOwnableID(ownableIDString string) (ids.OwnableID, error) {
	// TODO ***** never allow ownable ID to be valid hash string
	if ownableID, err := ReadOwnableID(ownableIDString); err == nil {
		return ownableID, nil
	}

	return &OwnableID{
		StringID: NewStringID(ownableIDString).(*StringID),
	}, nil
}
