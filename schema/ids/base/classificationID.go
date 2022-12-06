package base

import "C"
import (
	"bytes"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// type classificationID struct {
//	ids.HashID
// }
type classificationID base.ClassificationID

func (classificationID *classificationID) String() string {
	// TODO implement me
	panic("implement me")
}

var _ ids.ClassificationID = (*classificationID)(nil)

func (classificationID *classificationID) Bytes() []byte {
	return classificationID.HashId.Bytes()
}

func (classificationID *classificationID) IsClassificationID() {}

func (classificationID *classificationID) Compare(listable traits.Listable) int {
	return bytes.Compare(classificationID.Bytes(), classificationIDFromInterface(listable).Bytes())
}

func classificationIDFromInterface(i interface{}) classificationID {
	switch value := i.(type) {
	case classificationID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewClassificationID(immutables qualified.Immutables, mutables qualified.Mutables) ids.ClassificationID {
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

	return &classificationID{HashId: GenerateHashID(GenerateHashID(immutableIDByteList...).Bytes(), GenerateHashID(mutableIDByteList...).Bytes(), GenerateHashID(defaultImmutableByteList...).Bytes()).(*hashID)}
}

func PrototypeClassificationID() ids.ClassificationID {
	return &classificationID{
		HashId: PrototypeHashID().(*hashID),
	}
}

func ReadClassificationID(classificationIDString string) (ids.ClassificationID, error) {
	if hashID, err := ReadHashID(classificationIDString); err == nil {
		return &classificationID{HashId: hashID.(*hashID)}, nil
	}

	if classificationIDString == "" {
		return PrototypeClassificationID(), nil
	}

	return &classificationID{}, constants.MetaDataError
}
