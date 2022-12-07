package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ documents.Asset = (*AssetI_Asset)(nil)

func (asset *AssetI_Asset) GenerateHashID() ids.HashID {
	return asset.Asset.Document.GenerateHashID()
}

func (asset *AssetI_Asset) GetClassificationID() ids.ClassificationID {
	return asset.Asset.Document.GetClassificationID()
}

func (asset *AssetI_Asset) GetProperty(id ids.PropertyID) properties.Property {
	return asset.Asset.Document.GetProperty(id)
}

func (asset *AssetI_Asset) GetImmutables() qualified.Immutables {
	return asset.Asset.Document.GetImmutables()
}

func (asset *AssetI_Asset) GetMutables() qualified.Mutables {
	return asset.Asset.Document.GetMutables()
}

func (asset *AssetI_Asset) Mutate(property ...properties.Property) documents.Document {
	return asset.Asset.Document.Mutate(property...)
}

func (asset *AssetI_Asset) GetBurn() properties.Property {
	if burn := asset.GetProperty(constants.BurnHeightProperty.GetID()); burn != nil {
		return burn
	}

	return constants.BurnHeightProperty
}
func (asset *AssetI_Asset) GetLock() properties.Property {
	if lock := asset.GetProperty(constants.LockProperty.GetID()); lock != nil {
		return lock
	}

	return constants.LockProperty
}
func (asset *AssetI_Asset) GetSupply() properties.Property {
	if supply := asset.GetProperty(constants.SupplyProperty.GetID()); supply != nil {
		return supply
	}

	return constants.SupplyProperty
}

func NewAsset(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Asset {
	return &AssetI{
		Impl: &AssetI_Asset{
			Asset: &Asset{
				Document: NewDocument(classificationID, immutables, mutables).(*DocumentI),
			},
		},
	}
}
