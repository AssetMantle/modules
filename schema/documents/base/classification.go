package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

//
//type classification struct {
//	documents.Document
//}

var _ documents.Classification = (*Classification)(nil)

func (c *Classification) GenerateHashID() ids.HashID {
	return c.Document.GenerateHashID()
}

func (c *Classification) GetClassificationID() ids.ClassificationID {
	return c.Document.GetClassificationID()
}

func (c *Classification) GetProperty(id ids.PropertyID) properties.Property {
	return c.Document.GetProperty(id)
}

func (c *Classification) GetImmutables() qualified.Immutables {
	return c.Document.GetImmutables()
}

func (c *Classification) GetMutables() qualified.Mutables {
	return c.Document.GetMutables()
}

func (c *Classification) Mutate(property ...properties.Property) documents.Document {
	return c.Document.Mutate(property...)
}

func NewClassification(immutables qualified.Immutables, mutables qualified.Mutables) documents.Classification {
	return &Classification{
		Document: NewDocument(base.GenerateClassificationID(immutables, mutables), immutables, mutables),
	}
}
