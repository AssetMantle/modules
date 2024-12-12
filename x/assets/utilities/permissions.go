// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/schema/ids"
)

func SetModulePermissions(canMintAsset bool, canRenumerateAsset bool, canBurnAsset bool) []ids.StringID {
	var permissions []ids.StringID

	if canMintAsset {
		permissions = append(permissions, constants.CanMintAssetPermission)
	}
	if canRenumerateAsset {
		permissions = append(permissions, constants.CanRenumerateAssetPermission)
	}
	if canBurnAsset {
		permissions = append(permissions, constants.CanBurnAssetPermission)
	}

	return permissions
}
