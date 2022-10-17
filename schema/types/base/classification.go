package base

import (
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
)

type classification struct {
	qualified.Document
}

var _ types.Classification = (*classification)(nil)

func NewClassification(immutables qualified.Immutables, mutables qualified.Mutables) types.Classification {
	return classification{
		Document: baseQualified.NewDocument(base.NewClassificationID(immutables, mutables), immutables, mutables),
	}
}
