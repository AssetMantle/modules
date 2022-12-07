package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ documents.Classification = (*ClassificationI)(nil)

func (x *ClassificationI) GenerateHashID() ids.HashID {
	return x.Impl.(documents.Classification).GenerateHashID()
}

func (x *ClassificationI) GetClassificationID() ids.ClassificationID {
	return x.Impl.(documents.Classification).GetClassificationID()
}

func (x *ClassificationI) GetProperty(id ids.PropertyID) properties.Property {
	return x.Impl.(documents.Classification).GetProperty(id)
}

func (x *ClassificationI) GetImmutables() qualified.Immutables {
	return x.Impl.(documents.Classification).GetImmutables()
}

func (x *ClassificationI) GetMutables() qualified.Mutables {
	return x.Impl.(documents.Classification).GetMutables()
}

func (x *ClassificationI) Mutate(property ...properties.Property) documents.Document {
	return x.Impl.(documents.Classification).Mutate(property...)
}
