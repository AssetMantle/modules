package utilities

import (
	"github.com/AssetMantle/schema/go/ids"

	"github.com/AssetMantle/modules/x/orders/constants"
)

func SetModulePermissions(canMakeOrder bool, canModifyOrder bool, canCancelOrder bool) []ids.StringID {
	var permissions []ids.StringID

	if canMakeOrder {
		permissions = append(permissions, constants.CanMakeOrderPermission)
	}
	if canModifyOrder {
		permissions = append(permissions, constants.CanModifyOrderPermission)
	}
	if canCancelOrder {
		permissions = append(permissions, constants.CanCancelOrderPermission)
	}

	return permissions
}
