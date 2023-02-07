package utilities

import (
	"github.com/AssetMantle/modules/schema/ids/constansts"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
)

func SetPermissions(mint bool, burn bool, renumerate bool, add bool, remove bool, mutate bool) lists.IDList {
	permissions := base.NewIDList()

	if mint {
		permissions.Add(constansts.Mint)
	}
	if burn {
		permissions.Add(constansts.Burn)
	}
	if renumerate {
		permissions.Add(constansts.Renumerate)
	}
	if add {
		permissions.Add(constansts.Add)
	}
	if remove {
		permissions.Add(constansts.Remove)
	}
	if mutate {
		permissions.Add(constansts.Mutate)
	}

	return permissions
}
