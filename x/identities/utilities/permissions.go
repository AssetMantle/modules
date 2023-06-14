package utilities

import (
	"github.com/AssetMantle/schema/go/ids"

	"github.com/AssetMantle/modules/x/identities/constants"
)

func SetModulePermissions(canIssueIdentity bool, canMutateIdentity bool, canQuashIdentity bool) []ids.StringID {
	var permissions []ids.StringID

	if canIssueIdentity {
		permissions = append(permissions, constants.CanIssueIdentityPermission)
	}
	if canMutateIdentity {
		permissions = append(permissions, constants.CanMutateIdentityPermission)
	}
	if canQuashIdentity {
		permissions = append(permissions, constants.CanQuashIdentityPermission)
	}

	return permissions
}
