package base

import (
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids2.ClassificationID = (*ClassificationIDI)(nil)

func (classificationIDI *ClassificationIDI) Compare(listable traits.Listable) int {
	return classificationIDI.Impl.(ids2.ClassificationID).Compare(listable)
}

func (classificationIDI *ClassificationIDI) Bytes() []byte {
	return classificationIDI.Impl.(ids2.ClassificationID).Bytes()
}

func (classificationIDI *ClassificationIDI) IsClassificationID() {}
