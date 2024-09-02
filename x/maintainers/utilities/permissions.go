package utilities

import (
	"github.com/AssetMantle/modules/x/maintainers/constants"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/lists/base"
)

func SetModulePermissions(canAddMaintainer bool, canMutateMaintainer bool, canRemoveMaintainer bool) lists.IDList {
	permissions := base.NewIDList()

	if canAddMaintainer {
		permissions = permissions.Add(constants.CanAddMaintainerPermission)
	}
	if canMutateMaintainer {
		permissions = permissions.Add(constants.CanMutateMaintainerPermission)
	}
	if canRemoveMaintainer {
		permissions = permissions.Add(constants.CanRemoveMaintainerPermission)
	}

	return permissions
}
