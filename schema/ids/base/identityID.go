package base

import (
	"bytes"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// type identityID struct {
//	ids.HashID
// }

type identityID base.IdentityID

var _ ids.IdentityID = (*identityID)(nil)

func (identityID *identityID) String() string {
	return identityID.HashId.String()
}

// TODO deprecate
func (identityID *identityID) IsIdentityID() {}

func (identityID *identityID) Bytes() []byte {
	return identityID.HashId.Bytes()
}
func (identityID *identityID) Compare(listable traits.Listable) int {
	return bytes.Compare(identityID.Bytes(), identityIDFromInterface(listable).HashId.Bytes())
}
func (identityID *identityID) GetHashID() ids.HashID {
	return identityID.HashId
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
	return &identityID{
		HashId: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*hashID),
	}
}

func PrototypeIdentityID() ids.IdentityID {
	return &identityID{
		HashId: PrototypeHashID().(*hashID),
	}
}

func ReadIdentityID(identityIDString string) (ids.IdentityID, error) {

	if hashID, err := ReadHashID(identityIDString); err == nil {
		return &identityID{
			HashId: hashID.(*hashID),
		}, nil
	}

	return &identityID{}, errorConstants.IncorrectFormat
}
