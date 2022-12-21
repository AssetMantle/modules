package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

//type identityID struct {
//	ids.HashID
//}

var _ ids.IdentityID = (*IdentityID)(nil)

// TODO deprecate
func (identityID *IdentityID) IsIdentityID() {}
func (identityID *IdentityID) IDString() string {
	return identityID.HashId.EncodedString()
}
func (identityID *IdentityID) GetHashID() ids.HashID {
	return identityID.HashId
}
func (identityID *IdentityID) Bytes() []byte {
	return identityID.HashId.Bytes()
}
func (identityID *IdentityID) Compare(listable traits.Listable) int {
	return identityID.HashId.Compare(identityIDFromInterface(listable).HashId)
}
func (identityID *IdentityID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_IdentityId{
			IdentityId: identityID,
		},
	}
}

func identityIDFromInterface(i interface{}) *IdentityID {
	switch value := i.(type) {
	case *IdentityID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewIdentityID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.IdentityID {
	return &IdentityID{
		HashId: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeIdentityID() ids.IdentityID {
	return &IdentityID{
		HashId: PrototypeHashID().(*HashID),
	}
}

func ReadIdentityID(identityIDString string) (ids.IdentityID, error) {

	if hashID, err := ReadHashID(identityIDString); err == nil {
		return &IdentityID{
			HashId: hashID.(*HashID),
		}, nil
	}

	return &IdentityID{}, errorConstants.IncorrectFormat
}
