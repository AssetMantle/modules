package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.IdentityID = (*ID_IdentityID)(nil)

func (identityID *ID_IdentityID) String() string {
	return identityID.IdentityID.HashId.String()
}

// TODO deprecate
func (identityID *ID_IdentityID) IsIdentityID() {}

func (identityID *ID_IdentityID) Bytes() []byte {
	return identityID.IdentityID.HashId.IdBytes
}
func (identityID *ID_IdentityID) Compare(listable traits.Listable) int {
	return bytes.Compare(identityID.Bytes(), idFromInterface(listable).Bytes())
}
func (identityID *ID_IdentityID) GetHashID() ids.ID {
	return &ID{Impl: &ID_HashID{HashID: identityID.IdentityID.HashId}}
}

func GenerateIdentityID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.ID {
	return NewIdentityID(GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()))
}

func NewIdentityID(hashID ids.ID) ids.ID {
	if hashID.(*ID).GetHashID() == nil {
		panic(errorConstants.MetaDataError)
	}
	return &ID{
		Impl: &ID_IdentityID{
			IdentityID: &IdentityID{
				HashId: hashID.(*ID).GetHashID(),
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
