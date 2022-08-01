// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

// TODO check structure
type maintainer struct {
	qualified.Document
}

var _ mappables.Maintainer = (*maintainer)(nil)

func (maintainer maintainer) GetIdentityID() ids.ID {
	return key.ReadIdentityID(maintainer.ID)
}
func (maintainer maintainer) GetMaintainedClassificationID() ids.ID {
	return key.ReadClassificationID(maintainer.ID)
}
func (maintainer maintainer) GetMaintainedPropertySet() properties.Property {
	if property := maintainer.GetProperty(constants.MaintainedPropertiesProperty); property != nil {
		return property
	}
	return constants.MaintainedProperties
}

func (maintainer maintainer) CanMintAsset() bool {
	if property := maintainer.GetProperty(constants.PermissionsProperty); property != nil {
		if property.GetID().Compare(constants.Permissions.GetID()) == 0 {
			return true
		}
	}
	return false
}

// TODO **
func (maintainer maintainer) CanBurnAsset() bool {
	if property := maintainer.GetProperty(constants.PermissionsProperty); property != nil {
		// impl
	}

	return false
}

// TODO **
func (maintainer maintainer) CanRenumerateAsset() bool {
	if property := maintainer.GetProperty(constants.PermissionsProperty); property != nil {
		// impl
	}

	return false
}

// TODO **
func (maintainer maintainer) CanAddMaintainer() bool {
	if property := maintainer.GetProperty(constants.Permissions.GetID()); property != nil {
		// TODO impl
	}

	return false
}

// TODO **
func (maintainer maintainer) CanRemoveMaintainer() bool {
	if property := maintainer.GetProperty(constants.Permissions.GetID()); property != nil {
		// TODO impl
	}

	return false
}

// TODO **
func (maintainer maintainer) CanMutateMaintainer() bool {
	if property := maintainer.GetProperty(constants.PermissionsProperty); property != nil {
		// impl
	}

	return false
}
func (maintainer maintainer) MaintainsProperty(id ids.ID) bool {
	if property := maintainer.GetProperty(constants.PermissionsProperty); property != nil {
		if property.GetID().Compare(id) == 0 {
			return true
		}
	}

	return false
}
func (maintainer maintainer) GetKey() helpers.Key {
	return key.NewKey(maintainer.ID)
}
func (maintainer) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, maintainer{})
}

// TODO add maintainer identityID in immutables
func NewMaintainer(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) mappables.Maintainer {
	return maintainer{
		Document: baseQualified.NewDocument(classificationID, immutables, mutables),
	}
}
