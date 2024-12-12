// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/x/identities/constants"
	"github.com/AssetMantle/schema/ids"
)

func SetModulePermissions(canIssueIdentity bool, canQuashIdentity bool) []ids.StringID {
	var permissions []ids.StringID

	if canIssueIdentity {
		permissions = append(permissions, constants.CanIssueIdentityPermission)
	}
	if canQuashIdentity {
		permissions = append(permissions, constants.CanQuashIdentityPermission)
	}

	return permissions
}
