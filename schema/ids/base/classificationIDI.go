package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.ClassificationID = (*ClassificationIDI)(nil)

func (classificationIDI *ClassificationIDI) Compare(listable traits.Listable) int {
	return classificationIDI.Impl.(ids.ClassificationID).Compare(listable)
}

func (classificationIDI *ClassificationIDI) Bytes() []byte {
	return classificationIDI.Impl.(ids.ClassificationID).Bytes()
}

func (classificationIDI *ClassificationIDI) IsClassificationID() {}
