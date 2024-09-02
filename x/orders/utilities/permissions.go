package utilities

import (
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/schema/ids"
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
