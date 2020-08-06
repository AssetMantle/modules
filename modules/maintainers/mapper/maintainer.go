/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type maintainer struct {
	ID                  types.ID      `json:"id" valid:"required field id missing"`
	AddMaintainer       bool          `json:"addMaintainer" valid:"required field addMaintainer missing"`
	RemoveMaintainer    bool          `json:"removeMaintainer" valid:"required field removeMaintainer missing"`
	MutateMaintainer    bool          `json:"mutateMaintainer" valid:"required field mutateMaintainer missing"`
	MaintainedTraitList []types.Trait `json:"maintainedTraitList" valid:"required field maintainedTraitList missing"`
}

var _ mappables.Maintainer = (*maintainer)(nil)

func (maintainer maintainer) GetID() types.ID { return maintainer.ID }
func (maintainer maintainer) GetIdentityID() types.ID {
	return maintainerIDFromInterface(maintainer.ID).IdentityID
}
func (maintainer maintainer) GetMaintainedID() types.ID {
	return maintainerIDFromInterface(maintainer.ID).MaintainedID
}
func (maintainer maintainer) CanAddMaintainer() bool    { return maintainer.AddMaintainer }
func (maintainer maintainer) CanRemoveMaintainer() bool { return maintainer.RemoveMaintainer }
func (maintainer maintainer) CanMutateMaintainer() bool { return maintainer.MutateMaintainer }
func (maintainer maintainer) MaintainsTrait(id types.ID) bool {
	for _, trait := range maintainer.MaintainedTraitList {
		if trait.GetID().Compare(id) == 0 {
			return true
		}
	}
	return false
}
func (maintainer maintainer) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(maintainer)
}
func (maintainer maintainer) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &maintainer)
	return maintainer
}
func maintainerPrototype() traits.Mappable {
	return maintainer{}
}
func NewMaintainer(ID types.ID, addMaintainer bool, removeMaintainer bool, mutateMaintainer bool, maintainedTraitList []types.Trait) mappables.Maintainer {
	return maintainer{
		ID:                  ID,
		AddMaintainer:       addMaintainer,
		RemoveMaintainer:    removeMaintainer,
		MutateMaintainer:    mutateMaintainer,
		MaintainedTraitList: maintainedTraitList,
	}
}
