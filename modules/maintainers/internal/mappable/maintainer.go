// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/ids/constansts"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type maintainer struct {
	qualified.Document
}

var _ mappables.Maintainer = (*maintainer)(nil)

func (maintainer maintainer) GetIdentityID() ids.IdentityID {
	if property := maintainer.GetProperty(constants.IdentityIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.IdentityID)
	}
	return constants.MaintainedClassificationIDProperty.GetData().(data.IDData).Get().(ids.IdentityID)
}
func (maintainer maintainer) GetMaintainedClassificationID() ids.ClassificationID {
	if property := maintainer.GetProperty(constants.MaintainedClassificationIDProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.IDData).Get().(ids.ClassificationID)
	}
	return constants.MaintainedClassificationIDProperty.GetData().(data.IDData).Get().(ids.ClassificationID)
}
func (maintainer maintainer) GetMaintainedProperties() data.ListData {
	if property := maintainer.GetProperty(constants.MaintainedPropertiesProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.ListData)
	}

	return constants.MaintainedPropertiesProperty.GetData().(data.ListData)
}
func (maintainer maintainer) GetPermissions() data.ListData {
	if property := maintainer.GetProperty(constants.PermissionsProperty.GetID()); property != nil && property.IsMeta() {
		return property.(properties.MetaProperty).GetData().(data.ListData)
	}

	return constants.PermissionsProperty.GetData().(data.ListData)
}
func (maintainer maintainer) CanMintAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Mint))
	return can
}
func (maintainer maintainer) CanBurnAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Burn))
	return can
}
func (maintainer maintainer) CanRenumerateAsset() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Renumerate))
	return can
}
func (maintainer maintainer) CanAddMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Add))
	return can
}
func (maintainer maintainer) CanRemoveMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Remove))
	return can
}
func (maintainer maintainer) CanMutateMaintainer() bool {
	_, can := maintainer.GetPermissions().Search(baseData.NewIDData(constansts.Mutate))
	return can
}
func (maintainer maintainer) MaintainsProperty(id ids.ID) bool {
	if property := maintainer.GetProperty(constants.PermissionsProperty.GetID()); property != nil {
		if property.GetID().Compare(id) == 0 {
			return true
		}
	}

	return false
}
func (maintainer maintainer) GetKey() helpers.Key {
	return key.NewKey(base.NewMaintainerID(maintainer.GetMaintainedClassificationID(), maintainer.GetIdentityID()))
}
func (maintainer) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, maintainer{})
}

func NewMaintainer(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) mappables.Maintainer {
	return maintainer{
		Document: baseQualified.NewDocument(classificationID, immutables, mutables),
	}
}

func Prototype() helpers.Mappable {
	return maintainer{}
}
