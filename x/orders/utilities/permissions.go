package utilities

import (
	"github.com/AssetMantle/schema/ids"

	"github.com/AssetMantle/modules/x/orders/constants"
)

func SetModulePermissions(canMakeOrder bool, canCancelOrder bool) []ids.StringID {
	var permissions []ids.StringID

	if canMakeOrder {
		permissions = append(permissions, constants.CanMakeOrderPermission)
	}
	if canCancelOrder {
		permissions = append(permissions, constants.CanCancelOrderPermission)
	}

	return permissions
}
