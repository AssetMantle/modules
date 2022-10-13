package base

import (
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type classification struct {
	qualified.Document
}

var _ mappables.Classification = (*classification)(nil)

func NewClassification(immutables qualified.Immutables, mutables qualified.Mutables) mappables.Classification {
	return classification{
		Document: baseQualified.NewDocument(base.NewClassificationID(immutables, mutables), immutables, mutables),
	}
}
