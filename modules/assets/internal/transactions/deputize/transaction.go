// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var Transaction = baseHelpers.NewTransaction(
	"deputize",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	constants.FromID,
	constants.ToID,
	constants.ClassificationID,
	constants.MaintainedProperties,
	constants.CanMintAsset,
	constants.CanBurnAsset,
	constants.CanRenumerateAsset,
	constants.CanAddMaintainer,
	constants.CanRemoveMaintainer,
	constants.CanMutateMaintainer,
)
