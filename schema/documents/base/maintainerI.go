package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

var _ documents.Maintainer = (*MaintainerI)(nil)

func (x *MaintainerI) GetIdentityID() ids.IdentityID {
	return x.Impl.(documents.Maintainer).GetIdentityID()
}

func (x *MaintainerI) GetMaintainedClassificationID() ids.ClassificationID {
	return x.Impl.(documents.Maintainer).GetMaintainedClassificationID()
}

func (x *MaintainerI) GetMaintainedProperties() data.ListData {
	return x.Impl.(documents.Maintainer).GetMaintainedProperties()
}

func (x *MaintainerI) GetPermissions() data.ListData {
	return x.Impl.(documents.Maintainer).GetPermissions()
}

func (x *MaintainerI) CanMintAsset() bool {
	return x.Impl.(documents.Maintainer).CanMintAsset()
}

func (x *MaintainerI) CanBurnAsset() bool {
	return x.Impl.(documents.Maintainer).CanBurnAsset()
}

func (x *MaintainerI) CanRenumerateAsset() bool {
	return x.Impl.(documents.Maintainer).CanRenumerateAsset()
}

func (x *MaintainerI) CanAddMaintainer() bool {
	return x.Impl.(documents.Maintainer).CanAddMaintainer()
}

func (x *MaintainerI) CanRemoveMaintainer() bool {
	return x.Impl.(documents.Maintainer).CanRemoveMaintainer()
}

func (x *MaintainerI) CanMutateMaintainer() bool {
	return x.Impl.(documents.Maintainer).CanMutateMaintainer()
}

func (x *MaintainerI) MaintainsProperty(id ids.PropertyID) bool {
	return x.Impl.(documents.Maintainer).MaintainsProperty(id)
}

func (x *MaintainerI) GenerateHashID() ids.HashID {
	return x.Impl.(documents.Maintainer).GenerateHashID()
}

func (x *MaintainerI) GetClassificationID() ids.ClassificationID {
	return x.Impl.(documents.Maintainer).GetClassificationID()
}

func (x *MaintainerI) GetProperty(id ids.PropertyID) properties.Property {
	return x.Impl.(documents.Maintainer).GetProperty(id)
}

func (x *MaintainerI) GetImmutables() qualified.Immutables {
	return x.Impl.(documents.Maintainer).GetImmutables()
}
