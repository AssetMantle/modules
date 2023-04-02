package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
	"strings"
)

var _ ids.ClassificationID = (*ClassificationID)(nil)

func (classificationID *ClassificationID) ValidateBasic() error {
	return classificationID.HashID.ValidateBasic()
}
func (classificationID *ClassificationID) GetTypeID() ids.StringID {
	return NewStringID(constants.ClassificationIDType)
}
func (classificationID *ClassificationID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeClassificationID(), nil
	}

	if hashID, err := PrototypeHashID().FromString(idString); err != nil {
		return PrototypeClassificationID(), err
	} else {
		return &ClassificationID{
			HashID: hashID.(*HashID),
		}, nil
	}
}
func (classificationID *ClassificationID) AsString() string {
	return classificationID.HashID.AsString()
}
func (classificationID *ClassificationID) Bytes() []byte {
	return classificationID.HashID.IDBytes
}
func (classificationID *ClassificationID) IsClassificationID() {}
func (classificationID *ClassificationID) Compare(listable traits.Listable) int {
	return classificationID.HashID.Compare(classificationIDFromInterface(listable).HashID)
}
func (classificationID *ClassificationID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_ClassificationID{
			ClassificationID: classificationID,
		},
	}
}

func classificationIDFromInterface(i interface{}) *ClassificationID {
	switch value := i.(type) {
	case *ClassificationID:
		return value
	default:
		panic(errorConstants.IncorrectFormat.Wrapf("Expected ClassificationID, got %T", i))
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

	return &ClassificationID{HashID: GenerateHashID(GenerateHashID(immutableIDByteList...).Bytes(), GenerateHashID(mutableIDByteList...).Bytes(), immutables.GenerateHashID().Bytes()).(*HashID)}
}

func PrototypeClassificationID() ids.ClassificationID {
	return &ClassificationID{
		HashID: PrototypeHashID().(*HashID),
	}
}

func ReadClassificationID(classificationIDString string) (ids.ClassificationID, error) {
	if hashID, err := ReadHashID(classificationIDString); err == nil {
		return &ClassificationID{HashID: hashID.(*HashID)}, nil
	}

	if classificationIDString == "" {
		return PrototypeClassificationID(), nil
	}

	return &ClassificationID{}, errorConstants.IncorrectFormat.Wrapf("Invalid ClassificationID: %s", classificationIDString)
}
