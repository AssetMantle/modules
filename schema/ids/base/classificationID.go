package base

import "C"
import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// type classificationID struct {
//	ids.HashID
// }

var _ ids.ClassificationID = (*ID_ClassificationID)(nil)

func (classificationID *ID_ClassificationID) String() string {
	return classificationID.ClassificationID.String()
}
func (classificationID *ID_ClassificationID) Bytes() []byte {
	return classificationID.ClassificationID.HashId.IdBytes
}
func (classificationID *ID_ClassificationID) IsClassificationID() {}
func (classificationID *ID_ClassificationID) Compare(listable traits.Listable) int {
	return bytes.Compare(classificationID.Bytes(), idFromInterface(listable).Bytes())
}

func NewClassificationID(hashID ids.ID) ids.ID {
	return &ID{
		Impl: &ID_ClassificationID{
			ClassificationID: &ClassificationID{HashId: hashID.(*ID).GetHashID()},
		},
	}
}
func PrototypeClassificationID() ids.ID {
	return NewClassificationID(PrototypeHashID())
}

func ReadClassificationID(classificationIDString string) (ids.ID, error) {
	if hashID, err := ReadHashID(classificationIDString); err == nil {
		return NewClassificationID(hashID), nil
	}

	if classificationIDString == "" {
		return PrototypeClassificationID(), nil
	}

	return PrototypeClassificationID(), constants.MetaDataError
}
