// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/constants/properties"
	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/modules/maintainers/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	qualifiedMappables "github.com/AssetMantle/modules/schema/mappables/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

// TODO check structure
type maintainer struct {
	qualifiedMappables.Document
}

var _ mappables.Maintainer = (*maintainer)(nil)

func (maintainer maintainer) GetIdentityID() types.ID {
	return key.ReadIdentityID(maintainer.ID)
}
func (maintainer maintainer) GetMaintainedClassificationID() types.ID {
	return key.ReadClassificationID(maintainer.ID)
}
func (maintainer maintainer) GetMaintainedPropertySet() types.Property {
	if property := maintainer.GetProperty(ids.MaintainedPropertiesProperty); property != nil {
		return property
	}
	return properties.MaintainedProperties
}

func (maintainer maintainer) CanMintAsset() bool {
	if property := maintainer.GetProperty(ids.PermissionsProperty); property != nil {
		if property.GetID().Compare(properties.Permissions.GetID()) == 0 {
			return true
		}
	}
	return false
}

// TODO
func (maintainer maintainer) CanBurnAsset() bool {
	if property := maintainer.GetProperty(ids.PermissionsProperty); property != nil {
		// impl
	}

	return false
}

// TODO
func (maintainer maintainer) CanRenumerateAsset() bool {
	if property := maintainer.GetProperty(ids.PermissionsProperty); property != nil {
		// impl
	}

	return false
}

// TODO
func (maintainer maintainer) CanAddMaintainer() bool {
	if property := maintainer.GetProperty(baseTypes.NewID(properties.Permissions.GetID().String())); property != nil {
		// impl
	}

	return false
}

// TODO
func (maintainer maintainer) CanRemoveMaintainer() bool {
	if property := maintainer.GetProperty(baseTypes.NewID(properties.Permissions.GetID().String())); property != nil {
		// impl
	}

	return false
}

// TODO
func (maintainer maintainer) CanMutateMaintainer() bool {
	if property := maintainer.GetProperty(ids.PermissionsProperty); property != nil {
		// impl
	}

	return false
}
func (maintainer maintainer) MaintainsProperty(id types.ID) bool {
	if property := maintainer.GetProperty(ids.PermissionsProperty); property != nil {
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
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, maintainer{})
}

// TODO
func NewMaintainer(id types.ID, immutableProperties types.Properties, mutableProperties types.Properties) mappables.Maintainer {
	return maintainer{
		Document: qualifiedMappables.Document{
			ID:            id,
			HasImmutables: baseQualified.HasImmutables{Properties: immutableProperties},
			HasMutables:   baseQualified.HasMutables{Properties: mutableProperties},
		},
	}
}
