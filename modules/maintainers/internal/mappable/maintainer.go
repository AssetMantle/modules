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
	"github.com/persistenceOne/persistenceSDK/schema/traits/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type maintainer struct {
	qualified.Document
}

var _ mappables.Maintainer = (*maintainer)(nil)

func (maintainer maintainer) GetID() types.ID { return maintainer.ID }
func (maintainer maintainer) GetClassificationID() types.ID {
	return key.ReadClassificationID(maintainer.ID)
}
func (maintainer maintainer) GetProperty(id types.ID) types.Property {
	if property := maintainer.HasImmutables.GetImmutableProperties().Get(id); property != nil {
		return property
	} else if property := maintainer.HasMutables.GetMutableProperties().Get(id); property != nil {
		return property
	} else {
		return nil
	}
}
func (maintainer maintainer) GetIdentityID() types.ID {
	return key.ReadIdentityID(maintainer.ID)
}
func (maintainer maintainer) GetMaintainedClassificationID() types.ID {
	return key.ReadClassificationID(maintainer.ID)
}
func (maintainer maintainer) GetMaintainedProperties() types.Property {
	if property := maintainer.GetProperty(ids.MaintainedPropertiesProperty); property != nil {
		return property
	}

	return properties.MaintainedProperties
}
func (maintainer maintainer) CanMintAsset() bool {
	if property := maintainer.GetProperty(ids.PermissionsProperty); property != nil {
		impl
	}

	return false
}
func (maintainer maintainer) CanBurnAsset() bool {
	if property := maintainer.GetProperty(ids.PermissionsProperty); property != nil {
		impl
	}

	return false
}
func (maintainer maintainer) CanRenumerateAsset() bool {
	if property := maintainer.GetProperty(base.NewID(properties.Permissions)); property != nil {
		impl
	}

	return false
}
func (maintainer maintainer) CanAddMaintainer() bool {
	if property := maintainer.GetProperty(base.NewID(properties.Permissions)); property != nil {
		impl
	}

	return false
}
func (maintainer maintainer) CanRemoveMaintainer() bool {
	if property := maintainer.GetProperty(base.NewID(properties.Permissions)); property != nil {
		impl
	}

	return false
}
func (maintainer maintainer) CanMutateMaintainer() bool {
	if property := maintainer.GetProperty(base.NewID(properties.Permissions)); property != nil {
		impl
	}

	return false
}
func (maintainer maintainer) MaintainsProperty(id types.ID) bool {
	if property := maintainer.GetProperty(base.NewID(properties.Permissions)); property != nil {
		impl
	}

	return false
}
func (maintainer maintainer) GetKey() helpers.Key {
	return key.FromID(maintainer.ID)
}
func (maintainer) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, maintainer{})
}

func NewMaintainer(id types.ID, maintainedProperties types.Properties, addMaintainer bool, removeMaintainer bool, mutateMaintainer bool) mappables.Maintainer {
	return maintainer{
		Document: qualified.Document{
			ID:            id,
			HasImmutables: qualified.HasImmutables{Properties: immutableProperties},
			HasMutables:   qualified.HasMutables{Properties: mutableProperties},
		},
	}
}
