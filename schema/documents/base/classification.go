package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/qualified"
)

type classification struct {
	qualified.Document
}

var _ documents.Classification = (*classification)(nil)

func NewClassification(immutables qualified.Immutables, mutables qualified.Mutables) documents.Classification {
	return classification{
		Document: NewDocument(base.NewClassificationID(immutables, mutables), immutables, mutables),
	}
}
