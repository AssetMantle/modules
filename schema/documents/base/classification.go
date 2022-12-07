package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

type classification struct {
	documents.Document
}

var _ documents.Classification = (*ClassificationI_Classification)(nil)

func (c *ClassificationI_Classification) GenerateHashID() ids.HashID {
	return c.Classification.Document.GenerateHashID()
}

func (c *ClassificationI_Classification) GetClassificationID() ids.ClassificationID {
	return c.Classification.Document.GetClassificationID()
}

func (c *ClassificationI_Classification) GetProperty(id ids.PropertyID) properties.Property {
	return c.Classification.Document.GetProperty(id)
}

func (c *ClassificationI_Classification) GetImmutables() qualified.Immutables {
	return c.Classification.Document.GetImmutables()
}

func (c *ClassificationI_Classification) GetMutables() qualified.Mutables {
	return c.Classification.Document.GetMutables()
}

func (c *ClassificationI_Classification) Mutate(property ...properties.Property) documents.Document {
	return c.Classification.Document.Mutate(property...)
}

func NewClassification(immutables qualified.Immutables, mutables qualified.Mutables) documents.Classification {
	return &ClassificationI{
		Impl: &ClassificationI_Classification{
			Classification: &Classification{
				Document: NewDocument(base.GenerateClassificationID(immutables, mutables), immutables, mutables).(*DocumentI),
			},
		},
	}
}
