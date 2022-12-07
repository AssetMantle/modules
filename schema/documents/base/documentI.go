package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ documents.Document = (*DocumentI)(nil)

func (x *DocumentI) GenerateHashID() ids.HashID {
	return x.Impl.(documents.Document).GenerateHashID()
}

func (x *DocumentI) GetClassificationID() ids.ClassificationID {
	return x.Impl.(documents.Document).GetClassificationID()
}

func (x *DocumentI) GetProperty(id ids.PropertyID) properties.Property {
	return x.Impl.(documents.Document).GetProperty(id)
}

func (x *DocumentI) GetImmutables() qualified.Immutables {
	return x.Impl.(documents.Document).GetImmutables()
}

func (x *DocumentI) GetMutables() qualified.Mutables {
	return x.Impl.(documents.Document).GetMutables()
}

func (x *DocumentI) Mutate(property ...properties.Property) documents.Document {
	return x.Impl.(documents.Document).Mutate(property...)
}
