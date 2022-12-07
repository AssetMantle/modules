package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ documents.Asset = (*AssetI)(nil)

func (x *AssetI) GetBurn() properties.Property {
	return x.Impl.(documents.Asset).GetBurn()
}

func (x *AssetI) GetLock() properties.Property {
	return x.Impl.(documents.Asset).GetLock()
}

func (x *AssetI) GetSupply() properties.Property {
	return x.Impl.(documents.Asset).GetSupply()
}

func (x *AssetI) GenerateHashID() ids.HashID {
	return x.Impl.(documents.Asset).GenerateHashID()
}

func (x *AssetI) GetClassificationID() ids.ClassificationID {
	return x.Impl.(documents.Asset).GetClassificationID()
}

func (x *AssetI) GetProperty(id ids.PropertyID) properties.Property {
	return x.Impl.(documents.Asset).GetProperty(id)
}

func (x *AssetI) GetImmutables() qualified.Immutables {
	return x.Impl.(documents.Asset).GetImmutables()
}

func (x *AssetI) GetMutables() qualified.Mutables {
	return x.Impl.(documents.Asset).GetMutables()
}

func (x *AssetI) Mutate(property ...properties.Property) documents.Document {
	return x.Impl.(documents.Asset).Mutate(property...)
}
