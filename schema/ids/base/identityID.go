package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

type identityID struct {
	ids.ClassificationID
	ids.HashID
}

func (identityID identityID) IsIdentityID() {
	// TODO implement me
	panic("implement me")
}

var _ ids.IdentityID = (*identityID)(nil)

func (identityID identityID) String() string {
	return stringUtilities.JoinIDStrings(identityID.ClassificationID.String(), identityID.HashID.String())
}
func (identityID identityID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, identityID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, identityID.HashID.Bytes()...)

	return Bytes
}
func (identityID identityID) Compare(listable traits.Listable) int {
	return bytes.Compare(identityID.Bytes(), identityIDFromInterface(listable).Bytes())
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
		ClassificationID: classificationID,
		HashID:           immutables.GenerateHashID(),
	}
}

func ReadIdentityID(identityIDString string) (ids.IdentityID, error) {
	if classificationID, err := ReadClassificationID(stringUtilities.SplitCompositeIDString(identityIDString)[0]); err == nil {
		if hashID, err := ReadHashID(stringUtilities.SplitCompositeIDString(identityIDString)[1]); err == nil {
			return identityID{
				ClassificationID: classificationID,
				HashID:           hashID,
			}, nil
		}
	}
	return identityID{}, errorConstants.MetaDataError
}
