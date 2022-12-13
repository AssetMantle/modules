package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
)

type asset struct {
	documents.Document
}

var _ documents.Asset = (*asset)(nil)

func (asset asset) GetBurn() properties.Property {
	if burn := asset.GetProperty(constants.BurnHeightProperty.GetID()); burn != nil {
		return burn
	}

	return constants.BurnHeightProperty
}
func (asset asset) GetLock() properties.Property {
	if lock := asset.GetProperty(constants.LockProperty.GetID()); lock != nil {
		return lock
	}

	return constants.LockProperty
}
func (asset asset) GetSupply() properties.Property {
	if supply := asset.GetProperty(constants.SupplyProperty.GetID()); supply != nil {
		return supply
	}

	return constants.SupplyProperty
}

func NewAsset(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Asset {
	return asset{
		Document: NewDocument(classificationID, immutables, mutables),
	}
}
