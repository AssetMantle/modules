package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

type identityID struct {
	ids.HashID
}

var _ ids.IdentityID = (*identityID)(nil)

// TODO deprecate
func (identityID identityID) IsIdentityID() {}
func (identityID identityID) String() string {
	return identityID.HashID.String()
}
func (identityID identityID) Bytes() []byte {
	return identityID.HashID.Bytes()
}
func (identityID identityID) Compare(listable traits.Listable) int {
	return identityID.HashID.Compare(identityIDFromInterface(listable).HashID)
}
func (identityID identityID) GetHashID() ids.HashID {
	return identityID.HashID
}
func identityIDFromInterface(i interface{}) identityID {
	switch value := i.(type) {
	case identityID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewIdentityID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.IdentityID {
	return identityID{
		HashID: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()),
	}
}

func PrototypeIdentityID() ids.IdentityID {
	return identityID{
		HashID: PrototypeHashID(),
	}
}

func ReadIdentityID(identityIDString string) (ids.IdentityID, error) {

	if hashID, err := ReadHashID(identityIDString); err == nil {
		return identityID{
			HashID: hashID,
		}, nil
	}

	return identityID{}, errorConstants.IncorrectFormat
}
