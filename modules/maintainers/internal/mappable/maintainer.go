/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var _ mappables.Maintainer = (*Maintainer)(nil)

func (maintainer Maintainer) GetID() types.ID { return maintainer.ID }
func (maintainer Maintainer) GetClassificationID() types.ID {
	return key.ReadClassificationID(maintainer.ID)
}
func (maintainer Maintainer) GetIdentityID() types.ID {
	return key.ReadIdentityID(maintainer.ID)
}
func (maintainer Maintainer) GetMaintainedProperties() types.Properties {
	return maintainer.MaintainedProperties
}
func (maintainer Maintainer) CanAddMaintainer() bool    { return maintainer.AddMaintainer }
func (maintainer Maintainer) CanRemoveMaintainer() bool { return maintainer.RemoveMaintainer }
func (maintainer Maintainer) CanMutateMaintainer() bool { return maintainer.MutateMaintainer }
func (maintainer Maintainer) MaintainsProperty(id types.ID) bool {
	for _, property := range maintainer.MaintainedProperties.GetList() {
		if property.GetID().Compare(id) == 0 {
			return true
		}
	}

	return false
}
func (maintainer Maintainer) GetKey() helpers.Key {
	return key.FromID(maintainer.ID)
}

func (Maintainer) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, Maintainer{})
}
func NewMaintainer(id types.ID, maintainedProperties types.Properties, addMaintainer bool, removeMaintainer bool, mutateMaintainer bool) mappables.Maintainer {
	return &Maintainer{
		ID:                   id,
		MaintainedProperties: maintainedProperties,
		AddMaintainer:        addMaintainer,
		RemoveMaintainer:     removeMaintainer,
		MutateMaintainer:     mutateMaintainer,
	}
}
