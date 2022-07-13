// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	ids2 "github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/mappables"
	properties2 "github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

// TODO check structure
type maintainer struct {
	baseQualified.Document
}

var _ mappables.Maintainer = (*maintainer)(nil)

func (maintainer maintainer) GetIdentityID() ids2.ID {
	return key.ReadIdentityID(maintainer.ID)
}
func (maintainer maintainer) GetMaintainedClassificationID() ids2.ID {
	return key.ReadClassificationID(maintainer.ID)
}
func (maintainer maintainer) GetMaintainedPropertySet() properties2.Property {
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

// TODO
func (maintainer maintainer) CanBurnAsset() bool {
	if property := maintainer.GetProperty(constants.PermissionsProperty); property != nil {
		// impl
	}

	return false
}

// TODO
func (maintainer maintainer) CanRenumerateAsset() bool {
	if property := maintainer.GetProperty(constants.PermissionsProperty); property != nil {
		// impl
	}

	return false
}

// TODO
func (maintainer maintainer) CanAddMaintainer() bool {
	if property := maintainer.GetProperty(constants.Permissions.GetID()); property != nil {
		// TODO impl
	}

	return false
}

// TODO
func (maintainer maintainer) CanRemoveMaintainer() bool {
	if property := maintainer.GetProperty(constants.Permissions.GetID()); property != nil {
		// TODO impl
	}

	return false
}

// TODO
func (maintainer maintainer) CanMutateMaintainer() bool {
	if property := maintainer.GetProperty(constants.PermissionsProperty); property != nil {
		// impl
	}

	return false
}
func (maintainer maintainer) MaintainsProperty(id ids2.ID) bool {
	if property := maintainer.GetProperty(constants.PermissionsProperty); property != nil {
		if property.GetID().Compare(id) == 0 {
			return true
		}
	}

	return false
}
func (maintainer maintainer) GetKey() helpers.Key {
	return key.FromID(maintainer.ID)
}
func (maintainer) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, maintainer{})
}

// TODO
func NewMaintainer(id ids2.ID, immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) mappables.Maintainer {
	return maintainer{
		Document: baseQualified.Document{
			ID:         id,
			Immutables: baseQualified.Immutables{PropertyList: immutableProperties},
			Mutables:   baseQualified.Mutables{PropertyList: mutableProperties},
		},
	}
}
