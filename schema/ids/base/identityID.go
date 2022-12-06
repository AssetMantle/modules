package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// type identityID struct {
//	ids.HashID
// }

var _ ids.IdentityID = (*IdentityIDI_IdentityID)(nil)

func (identityID *IdentityIDI_IdentityID) String() string {
	return identityID.IdentityID.HashId.String()
}

// TODO deprecate
func (identityID *IdentityIDI_IdentityID) IsIdentityID() {}

func (identityID *IdentityIDI_IdentityID) Bytes() []byte {
	return identityID.IdentityID.HashId.Bytes()
}
func (identityID *IdentityIDI_IdentityID) Compare(listable traits.Listable) int {
	return bytes.Compare(identityID.Bytes(), identityIDFromInterface(listable).Bytes())
}
func (identityID *IdentityIDI_IdentityID) GetHashID() ids.HashID {
	return identityID.IdentityID.HashId
}
func identityIDFromInterface(i interface{}) *IdentityIDI {
	switch value := i.(type) {
	case *IdentityIDI:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func GenerateIdentityID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.IdentityID {
	return NewIdentityID(GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()))
}

func NewIdentityID(hashID ids.HashID) ids.IdentityID {
	return &IdentityIDI{
		Impl: &IdentityIDI_IdentityID{
			IdentityID: &IdentityID{
				HashId: hashID.(*HashIDI),
			},
		},
	}
}

func PrototypeIdentityID() ids.IdentityID {
	return NewIdentityID(PrototypeHashID())
}

func ReadIdentityID(identityIDString string) (ids.IdentityID, error) {

	if hashID, err := ReadHashID(identityIDString); err == nil {
		return NewIdentityID(hashID), nil
	}

	return PrototypeIdentityID(), errorConstants.IncorrectFormat
}
