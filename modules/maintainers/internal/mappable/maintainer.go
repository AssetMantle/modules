// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/mappables"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type maintainer struct {
	baseQualified.Document
}

var _ mappables.Maintainer = (*maintainer)(nil)

func (maintainer maintainer) GetClassificationID() ids.ID {
	return key.ReadClassificationID(maintainer.ID)
}
func (maintainer maintainer) GetIdentityID() ids.ID {
	return key.ReadIdentityID(maintainer.ID)
}
func (maintainer maintainer) GetMaintainedClassificationID() ids.ID {
	return key.ReadClassificationID(maintainer.ID)
}
func (maintainer maintainer) CanMintAsset() bool {
	return true
}
func (maintainer maintainer) CanBurnAsset() bool {
	return true
}
func (maintainer maintainer) CanRenumerateAsset() bool {
	return true
}
func (maintainer maintainer) CanAddMaintainer() bool {
	return true
}
func (maintainer maintainer) CanRemoveMaintainer() bool {
	return true
}
func (maintainer maintainer) CanMutateMaintainer() bool {
	return true
}
func (maintainer maintainer) MaintainsProperty(propertyID ids.PropertyID) bool {
	if property := maintainer.GetProperty(propertyID); property != nil {
		return true
	}
	return false
}
func (maintainer maintainer) GetKey() helpers.Key {
	return key.FromID(maintainer.ID)
}
func (maintainer) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, maintainer{})
}

func NewMaintainer(id ids.ID, immutableProperties lists.PropertyList, mutableProperties lists.PropertyList) mappables.Maintainer {
	return maintainer{
		Document: baseQualified.Document{
			ID:         id,
			Immutables: baseQualified.Immutables{PropertyList: immutableProperties},
			Mutables:   baseQualified.Mutables{PropertyList: mutableProperties},
		},
	}
}
