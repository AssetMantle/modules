package base

import "C"
import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
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
func GenerateClassificationID(immutables qualified.Immutables, mutables qualified.Mutables) ids.ID {
	immutableIDByteList := make([][]byte, len(immutables.GetImmutablePropertyList().GetList()))

	for i, property := range immutables.GetImmutablePropertyList().GetList() {
		immutableIDByteList[i] = property.GetID().Bytes()
	}

	mutableIDByteList := make([][]byte, len(mutables.GetMutablePropertyList().GetList()))

	for i, property := range mutables.GetMutablePropertyList().GetList() {
		mutableIDByteList[i] = property.GetID().Bytes()
	}

	defaultImmutableByteList := make([][]byte, len(immutables.GetImmutablePropertyList().GetList()))

	for i, property := range immutables.GetImmutablePropertyList().GetList() {
		// TODO test
		if hashID := property.GetDataID().GetHashID(); !(hashID.Compare(GenerateHashID()) == 0) {
			defaultImmutableByteList[i] = hashID.Bytes()
		}
	}

	return NewClassificationID(GenerateHashID(GenerateHashID(immutableIDByteList...).Bytes(), GenerateHashID(mutableIDByteList...).Bytes(), GenerateHashID(defaultImmutableByteList...).Bytes()))
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
