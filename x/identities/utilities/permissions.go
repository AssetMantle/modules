package utilities

import (
	"github.com/AssetMantle/schema/ids"

	"github.com/AssetMantle/modules/x/identities/constants"
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
