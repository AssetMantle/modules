// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappables

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/types"
)

// TODO check implementation
type Maintainer interface {
	GetIdentityID() types.ID
	GetMaintainedClassificationID() types.ID
	GetMaintainedPropertySet() types.Property

	CanMintAsset() bool
	CanBurnAsset() bool
	CanRenumerateAsset() bool
	CanAddMaintainer() bool
	CanRemoveMaintainer() bool
	CanMutateMaintainer() bool
	MaintainsProperty(types.ID) bool

	Document
	helpers.Mappable
}
