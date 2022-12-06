package base

import (
	"bytes"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// // TODO rename to something more appropriate
// type ownableID struct {
//	ids.StringID
// }
type ownableID base.OwnableID

func (ownableID *ownableID) String() string {
	// TODO implement me
	panic("implement me")
}

var _ ids.OwnableID = (*ownableID)(nil)

func (ownableID *ownableID) Bytes() []byte {
	return ownableID.StringId.Bytes()
}
func (ownableID *ownableID) IsOwnableID() {}
func (ownableID *ownableID) Compare(listable traits.Listable) int {
	// TODO devise a better strategy to compare assetID and ownableID
	return bytes.Compare(ownableID.Bytes(), ownableIDFromInterface(listable).Bytes())
}
func ownableIDFromInterface(i interface{}) ids.OwnableID {
	switch value := i.(type) {
	case ownableID:
		return &value
	default:
		panic(constants.MetaDataError)
	}
}
func NewOwnableID(stringID ids.StringID) ids.OwnableID {
	return &ownableID{
		StringId: stringID.(*stringID),
	}
}

func PrototypeOwnableID() ids.OwnableID {
	return &ownableID{
		StringId: PrototypeStringID().(*stringID),
	}
}

func ReadOwnableID(ownableIDString string) (ids.OwnableID, error) {
	// TODO ***** never allow ownable PropertyID to be valid hash string
	if assetID, err := ReadAssetID(ownableIDString); err == nil {
		return assetID, nil
	}

	return &ownableID{
		StringId: NewStringID(ownableIDString).(*stringID),
	}, nil
}
