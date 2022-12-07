package base

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/constansts"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
)

//type maintainer struct {
//	documents.Document
//}

var _ documents.Maintainer = (*Maintainer)(nil)

func (maintainer *Maintainer) GenerateHashID() ids.HashID {
	return maintainer.Document.GenerateHashID()
}

func (maintainer *Maintainer) GetClassificationID() ids.ClassificationID {
	return maintainer.Document.GetClassificationID()
}

func (maintainer *Maintainer) GetProperty(id ids.PropertyID) properties.Property {
	return maintainer.Document.GetProperty(id)
}

func (maintainer *Maintainer) GetImmutables() qualified.Immutables {
	return maintainer.Document.GetImmutables()
}

func (maintainer *Maintainer) GetMutables() qualified.Mutables {
	return maintainer.Document.GetMutables()
}

func (maintainer *Maintainer) Mutate(property ...properties.Property) documents.Document {
	return maintainer.Document.Mutate(property...)
}

func (maintainer *Maintainer) GetIdentityID() ids.IdentityID {
	if property := maintainer.Document.GetProperty(constantProperties.IdentityIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.IdentityID)
	}
	return constantProperties.IdentityIDProperty.GetData().(data.IDData).Get().(ids.IdentityID)
}
func (maintainer *Maintainer) GetMaintainedClassificationID() ids.ClassificationID {
	if property := maintainer.Document.GetProperty(constantProperties.MaintainedClassificationIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.ClassificationID)
	}
	return constantProperties.MaintainedClassificationIDProperty.GetData().(data.IDData).Get().(ids.ClassificationID)
}
func (maintainer *Maintainer) GetMaintainedProperties() data.ListData {
	if property := maintainer.Document.GetProperty(constantProperties.MaintainedPropertiesProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.ListData)
	}

	return constantProperties.MaintainedPropertiesProperty.GetData().(data.ListData)
}
func (maintainer *Maintainer) GetPermissions() data.ListData {
	if property := maintainer.Document.GetProperty(constantProperties.PermissionsProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.ListData)
	}

	return constantProperties.PermissionsProperty.GetData().(data.ListData)
}
func (maintainer *Maintainer) CanMintAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Mint))
	return can
}
func (maintainer *Maintainer) CanBurnAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Burn))
	return can
}
func (maintainer *Maintainer) CanRenumerateAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Renumerate))
	return can
}
func (maintainer *Maintainer) CanAddMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Add))
	return can
}
func (maintainer *Maintainer) CanRemoveMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Remove))
	return can
}
func (maintainer *Maintainer) CanMutateMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Mutate))
	return can
}
func (maintainer *Maintainer) MaintainsProperty(propertyID ids.PropertyID) bool {
	_, found := maintainer.GetMaintainedProperties().Search(baseData.NewIDData(propertyID))
	return found
}

// TODO: Move to a common package
func idListToDataList(idList lists.IDList) lists.DataList {
	dataList := baseLists.NewDataList()
	for _, id := range idList.GetList() {
		dataList = dataList.Add(baseData.NewIDData(id))
	}
	return dataList
}

func NewMaintainer(identityID ids.IdentityID, maintainedClassificationID ids.ClassificationID, maintainedPropertyIDList lists.IDList, permissions lists.IDList) documents.Maintainer {
	return &Maintainer{
		Document: NewDocument(maintainedClassificationID, maintainedPropertyIDList),
	}
}
