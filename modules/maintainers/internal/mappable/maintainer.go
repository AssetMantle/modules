/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/persistenceOne/persistenceSDK/constants/ids"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	qualifiedMappables "github.com/persistenceOne/persistenceSDK/schema/mappables/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/qualified/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
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
	if property := maintainer.GetProperty(base.NewID(properties.Permissions.GetID().String())); property != nil {
		// impl
	}

	return false
}

// TODO
func (maintainer maintainer) CanRemoveMaintainer() bool {
	if property := maintainer.GetProperty(base.NewID(properties.Permissions.GetID().String())); property != nil {
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
			HasImmutables: base.HasImmutables{Properties: immutableProperties},
			HasMutables:   base.HasMutables{Properties: mutableProperties},
		},
	}
}
