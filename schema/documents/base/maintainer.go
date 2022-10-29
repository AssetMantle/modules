package base

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/documents/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
)

type maintainer struct {
	documents.Document
}

var _ documents.Maintainer = (*maintainer)(nil)

func (maintainer maintainer) GetIdentityID() ids.IdentityID {
	if property := maintainer.GetProperty(constantProperties.IdentityIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.IdentityID)
	}
	return constantProperties.MaintainedClassificationIDProperty.GetData().(data.IDData).Get().(ids.IdentityID)
}
func (maintainer maintainer) GetMaintainedClassificationID() ids.ClassificationID {
	if property := maintainer.GetProperty(constantProperties.MaintainedClassificationIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.ClassificationID)
	}
	return constantProperties.MaintainedClassificationIDProperty.GetData().(data.IDData).Get().(ids.ClassificationID)
}
func (maintainer maintainer) GetMaintainedProperties() data.ListData {
	if property := maintainer.GetProperty(constantProperties.MaintainedPropertiesProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.ListData)
	}

	return constantProperties.MaintainedPropertiesProperty.GetData().(data.ListData)
}
func (maintainer maintainer) GetPermissions() data.ListData {
	if property := maintainer.GetProperty(constantProperties.PermissionsProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.ListData)
	}

	return constantProperties.PermissionsProperty.GetData().(data.ListData)
}
func (maintainer maintainer) CanMintAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constants.Mint))
	return can
}
func (maintainer maintainer) CanBurnAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constants.Burn))
	return can
}
func (maintainer maintainer) CanRenumerateAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constants.Renumerate))
	return can
}
func (maintainer maintainer) CanAddMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constants.Add))
	return can
}
func (maintainer maintainer) CanRemoveMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constants.Remove))
	return can
}
func (maintainer maintainer) CanMutateMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constants.Mutate))
	return can
}
func (maintainer maintainer) MaintainsProperty(propertyID ids.PropertyID) bool {
	_, found := maintainer.GetMaintainedProperties().Search(baseData.NewIDData(propertyID))
	return found
}

func NewMaintainer(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Maintainer {
	return maintainer{
		Document: NewDocument(classificationID, immutables, mutables),
	}
}
