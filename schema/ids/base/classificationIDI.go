package base

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids"

	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type classificationIDI ids.ClassificationID

var _ ids2.ClassificationID = (*classificationID)(nil)

func (classificationIDI *classificationIDI) Compare(listable traits.Listable) int {
	return classificationIDI.Impl.(ids2.ClassificationID).Compare(listable)
}

func (classificationIDI *classificationIDI) String() string {
	return classificationIDI.Impl.(ids2.ClassificationID).String()
}

func (classificationIDI *classificationIDI) Bytes() []byte {
	return classificationIDI.Impl.(ids2.ClassificationID).Bytes()
}

func (classificationIDI *classificationIDI) IsClassificationID() {}
