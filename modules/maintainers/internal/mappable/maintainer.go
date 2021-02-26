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

type maintainer struct {
	ID                   types.ID         `json:"id" valid:"required field key missing"`
	MaintainedProperties types.Properties `json:"maintainedProperties" valid:"required field maintainedProperties missing"`
	AddMaintainer        bool             `json:"addMaintainer" valid:"required field addMaintainer missing"`
	RemoveMaintainer     bool             `json:"removeMaintainer" valid:"required field removeMaintainer missing"`
	MutateMaintainer     bool             `json:"mutateMaintainer" valid:"required field mutateMaintainer missing"`
}

var _ mappables.Maintainer = (*maintainer)(nil)

func (maintainer maintainer) GetID() types.ID { return maintainer.ID }
func (maintainer maintainer) GetClassificationID() types.ID {
	return key.ReadClassificationID(maintainer.ID)
}
func (maintainer maintainer) GetIdentityID() types.ID {
	return key.ReadIdentityID(maintainer.ID)
}
func (maintainer maintainer) GetMaintainedProperties() types.Properties {
	return maintainer.MaintainedProperties
}
func (maintainer maintainer) CanAddMaintainer() bool    { return maintainer.AddMaintainer }
func (maintainer maintainer) CanRemoveMaintainer() bool { return maintainer.RemoveMaintainer }
func (maintainer maintainer) CanMutateMaintainer() bool { return maintainer.MutateMaintainer }
func (maintainer maintainer) MaintainsProperty(id types.ID) bool {
	for _, property := range maintainer.MaintainedProperties.GetList() {
		if property.GetID().Equals(id) {
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
func NewMaintainer(id types.ID, maintainedProperties types.Properties, addMaintainer bool, removeMaintainer bool, mutateMaintainer bool) mappables.Maintainer {
	return maintainer{
		ID:                   id,
		MaintainedProperties: maintainedProperties,
		AddMaintainer:        addMaintainer,
		RemoveMaintainer:     removeMaintainer,
		MutateMaintainer:     mutateMaintainer,
	}
}
