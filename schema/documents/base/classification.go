package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/qualified"
)

type classification struct {
	documents.Document
}

var _ documents.Classification = (*classification)(nil)

func NewClassification(immutables qualified.Immutables, mutables qualified.Mutables) documents.Classification {
	return classification{
		Document: NewDocument(base.GenerateClassificationID(immutables, mutables), immutables, mutables),
	}
}
