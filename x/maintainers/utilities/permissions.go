package utilities

import (
	"github.com/AssetMantle/schema/go/lists"
	"github.com/AssetMantle/schema/go/lists/base"

	"github.com/AssetMantle/modules/x/maintainers/constants"
)

func SetPermissions(mint bool, burn bool, renumerate bool, add bool, remove bool, mutate bool) lists.IDList {
	permissions := base.NewIDList()

	if mint {
		permissions = permissions.Add(constants.Mint)
	}
	if burn {
		permissions = permissions.Add(constants.Burn)
	}
	if renumerate {
		permissions = permissions.Add(constants.Renumerate)
	}
	if add {
		permissions = permissions.Add(constants.Add)
	}
	if remove {
		permissions = permissions.Add(constants.Remove)
	}
	if mutate {
		permissions = permissions.Add(constants.Mutate)
	}

	return permissions
}
