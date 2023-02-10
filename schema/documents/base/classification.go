package base

import (
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
)

type classification struct {
	documents.Document
}

var _ documents.Classification = (*classification)(nil)

func (classification classification) GetBondAmount() types.Dec {
	if property := classification.Document.GetProperty(constants.BondAmountProperty.GetID()); property != nil && property.IsMeta() {
		return property.Get().(properties.MetaProperty).GetData().Get().(data.DecData).Get()
	}

	return constants.BondAmountProperty.GetData().Get().(data.DecData).Get()
}
func NewClassification(immutables qualified.Immutables, mutables qualified.Mutables) documents.Classification {
	return classification{
		Document: NewDocument(base.NewClassificationID(immutables, mutables), immutables, mutables),
	}
}

func NewClassificationFromDocument(document documents.Document) documents.Classification {
	return classification{
		Document: document,
	}
}
