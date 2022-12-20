package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.IdentityID = (*ID_IdentityID)(nil)

func (identityID *ID_IdentityID) String() string {
	return identityID.IdentityID.HashID.String()
}

// TODO deprecate
func (identityID *ID_IdentityID) IsIdentityID() {}

func (identityID *ID_IdentityID) Bytes() []byte {
	return identityID.IdentityID.HashID.IdBytes
}
func (identityID *ID_IdentityID) Compare(listable traits.Listable) int {
	return bytes.Compare(identityID.Bytes(), idFromInterface(listable).Bytes())
}
func (identityID *ID_IdentityID) GetHashID() ids.ID {
	return &ID{Impl: &ID_HashID{HashID: identityID.IdentityID.HashID}}
}

func NewIdentityID(hashID ids.ID) ids.ID {
	if hashID.(*ID).GetHashID() == nil {
		panic(errorConstants.MetaDataError)
	}
	return &ID{
		Impl: &ID_IdentityID{
			IdentityID: &IdentityID{
				HashID: hashID.(*ID).GetHashID(),
			},
		},
	}
}

func PrototypeIdentityID() ids.ID {
	return NewIdentityID(PrototypeHashID())
}

func ReadIdentityID(identityIDString string) (ids.ID, error) {

	if hashID, err := ReadHashID(identityIDString); err == nil {
		return NewIdentityID(hashID), nil
	}

	return PrototypeIdentityID(), errorConstants.IncorrectFormat
}
