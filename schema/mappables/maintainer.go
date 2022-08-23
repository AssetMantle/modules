// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappables

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
)

type Maintainer interface {
	GetIdentityID() ids.IdentityID
	GetMaintainedClassificationID() ids.ClassificationID
	GetMaintainedProperties() data.ListData
	GetPermissions() data.ListData

	CanMintAsset() bool
	CanBurnAsset() bool
	CanRenumerateAsset() bool
	CanAddMaintainer() bool
	CanRemoveMaintainer() bool
	CanMutateMaintainer() bool
	MaintainsProperty(ids.ID) bool

	qualified.Document
	helpers.Mappable
}
