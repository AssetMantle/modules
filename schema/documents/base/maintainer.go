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
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

//type maintainer struct {
//	documents.Document
//}

var _ documents.Maintainer = (*MaintainerI_Maintainer)(nil)

func (maintainer *MaintainerI_Maintainer) GenerateHashID() ids.HashID {
	return maintainer.Maintainer.Document.GenerateHashID()
}

func (maintainer *MaintainerI_Maintainer) GetClassificationID() ids.ClassificationID {
	return maintainer.Maintainer.Document.GetClassificationID()
}

func (maintainer *MaintainerI_Maintainer) GetProperty(id ids.PropertyID) properties.Property {
	return maintainer.Maintainer.Document.GetProperty(id)
}

func (maintainer *MaintainerI_Maintainer) GetImmutables() qualified.Immutables {
	return maintainer.Maintainer.Document.GetImmutables()
}

func (maintainer *MaintainerI_Maintainer) GetMutables() qualified.Mutables {
	return maintainer.Maintainer.Document.GetMutables()
}

func (maintainer *MaintainerI_Maintainer) Mutate(property ...properties.Property) documents.Document {
	return maintainer.Maintainer.Document.Mutate(property...)
}

func (maintainer *MaintainerI_Maintainer) GetIdentityID() ids.IdentityID {
	if property := maintainer.Maintainer.Document.GetProperty(constantProperties.IdentityIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.IdentityID)
	}
	return constantProperties.IdentityIDProperty.GetData().(data.IDData).Get().(ids.IdentityID)
}
func (maintainer *MaintainerI_Maintainer) GetMaintainedClassificationID() ids.ClassificationID {
	if property := maintainer.Maintainer.Document.GetProperty(constantProperties.MaintainedClassificationIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.ClassificationID)
	}
	return constantProperties.MaintainedClassificationIDProperty.GetData().(data.IDData).Get().(ids.ClassificationID)
}
func (maintainer *MaintainerI_Maintainer) GetMaintainedProperties() data.ListData {
	if property := maintainer.Maintainer.Document.GetProperty(constantProperties.MaintainedPropertiesProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.ListData)
	}

	return constantProperties.MaintainedPropertiesProperty.GetData().(data.ListData)
}
func (maintainer *MaintainerI_Maintainer) GetPermissions() data.ListData {
	if property := maintainer.Maintainer.Document.GetProperty(constantProperties.PermissionsProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.ListData)
	}

	return constantProperties.PermissionsProperty.GetData().(data.ListData)
}
func (maintainer *MaintainerI_Maintainer) CanMintAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Mint))
	return can
}
func (maintainer *MaintainerI_Maintainer) CanBurnAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Burn))
	return can
}
func (maintainer *MaintainerI_Maintainer) CanRenumerateAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Renumerate))
	return can
}
func (maintainer *MaintainerI_Maintainer) CanAddMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Add))
	return can
}
func (maintainer *MaintainerI_Maintainer) CanRemoveMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Remove))
	return can
}
func (maintainer *MaintainerI_Maintainer) CanMutateMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Mutate))
	return can
}
func (maintainer *MaintainerI_Maintainer) MaintainsProperty(propertyID ids.PropertyID) bool {
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
	return &MaintainerI{
		Impl: &MaintainerI_Maintainer{
			Maintainer: &Maintainer{
				//Document: NewDocument(constansts.MaintainerClassificationID,
				//	baseQualified.NewImmutables(baseLists.NewPropertyList(
				//		baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(identityID)),
				//		baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(maintainedClassificationID)),
				//	)),
				//	baseQualified.NewMutables(baseLists.NewPropertyList(
				//		baseProperties.NewMetaProperty(constantProperties.MaintainedPropertiesProperty.GetKey(), baseData.NewListData(idListToDataList(maintainedPropertyIDList))),
				//		baseProperties.NewMetaProperty(constantProperties.PermissionsProperty.GetKey(), baseData.NewListData(idListToDataList(permissions))),
				//	)),
				//),
			}
		}
	}
}
