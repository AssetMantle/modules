// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappables

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
)

// TODO check implementation
type Maintainer interface {
	GetIdentityID() ids.ID
	GetMaintainedClassificationID() ids.ID

	CanMintAsset() bool
	CanBurnAsset() bool
	CanRenumerateAsset() bool
	CanAddMaintainer() bool
	CanRemoveMaintainer() bool
	CanMutateMaintainer() bool
	MaintainsProperty(ids.PropertyID) bool

	qualified.Document
	helpers.Mappable
}
