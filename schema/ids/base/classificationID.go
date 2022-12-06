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

var _ ids.ClassificationID = (*ClassificationIDI_ClassificationID)(nil)

func (classificationID *ClassificationIDI_ClassificationID) String() string {
	return classificationID.ClassificationID.String()
}
func (classificationID *ClassificationIDI_ClassificationID) Bytes() []byte {
	return classificationID.ClassificationID.HashId.Bytes()
}
func (classificationID *ClassificationIDI_ClassificationID) IsClassificationID() {}
func (classificationID *ClassificationIDI_ClassificationID) Compare(listable traits.Listable) int {
	return bytes.Compare(classificationID.Bytes(), classificationIDFromInterface(listable).Bytes())
}

func classificationIDFromInterface(i interface{}) *ClassificationIDI {
	switch value := i.(type) {
	case *ClassificationIDI:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func GenerateClassificationID(immutables qualified.Immutables, mutables qualified.Mutables) ids.ClassificationID {
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

func NewClassificationID(hashID ids.HashID) ids.ClassificationID {
	return &ClassificationIDI{
		Impl: &ClassificationIDI_ClassificationID{
			ClassificationID: &ClassificationID{HashId: hashID.(*HashIDI)},
		},
	}
}
func PrototypeClassificationID() ids.ClassificationID {
	return NewClassificationID(PrototypeHashID())
}

func ReadClassificationID(classificationIDString string) (ids.ClassificationID, error) {
	if hashID, err := ReadHashID(classificationIDString); err == nil {
		return NewClassificationID(hashID), nil
	}

	if classificationIDString == "" {
		return PrototypeClassificationID(), nil
	}

	return PrototypeClassificationID(), constants.MetaDataError
}
