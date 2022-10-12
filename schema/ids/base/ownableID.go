package base

import (
	"github.com/AssetMantle/modules/schema/ids"
)

type ownableID struct {
	ids.StringID
}

func (o ownableID) IsOwnableID() {
	// TODO implement me
	panic("implement me")
}

var _ ids.OwnableID = (*ownableID)(nil)

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
	return ownableID{
		StringID: NewStringID(ownableIDString),
	}, nil
}
