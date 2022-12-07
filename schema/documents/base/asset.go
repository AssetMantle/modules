package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ documents.Asset = (*Asset)(nil)

func (asset *Asset) GenerateHashID() ids.HashID {
	return asset.Document.GenerateHashID()
}

func (asset *Asset) GetClassificationID() ids.ClassificationID {
	return asset.Document.GetClassificationID()
}

func (asset *Asset) GetProperty(id ids.PropertyID) properties.Property {
	return asset.Document.GetProperty(id)
}

func (asset *Asset) GetImmutables() qualified.Immutables {
	return asset.Document.GetImmutables()
}

func (asset *Asset) GetMutables() qualified.Mutables {
	return asset.Document.GetMutables()
}

func (asset *Asset) Mutate(property ...properties.Property) documents.Document {
	return asset.Document.Mutate(property...)
}

func (asset *Asset) GetBurn() properties.Property {
	if burn := asset.GetProperty(constants.BurnHeightProperty.GetID()); burn != nil {
		return burn
	}

	return constants.BurnHeightProperty
}
func (asset *Asset) GetLock() properties.Property {
	if lock := asset.GetProperty(constants.LockProperty.GetID()); lock != nil {
		return lock
	}

	return constants.LockProperty
}
func (asset *Asset) GetSupply() properties.Property {
	if supply := asset.GetProperty(constants.SupplyProperty.GetID()); supply != nil {
		return supply
	}

	return constants.SupplyProperty
}

func NewAsset(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Asset {
	return &Asset{
		Document: NewDocument(classificationID, immutables, mutables).(*Document),
	}
}
