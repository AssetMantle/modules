package utilities

import (
	"github.com/AssetMantle/schema/go/lists"
	"github.com/AssetMantle/schema/go/lists/base"

	"github.com/AssetMantle/modules/x/maintainers/internal/module"
)

func SetPermissions(mint bool, burn bool, renumerate bool, add bool, remove bool, mutate bool) lists.IDList {
	permissions := base.NewIDList()

	if mint {
		permissions = permissions.Add(module.Mint)
	}
	if burn {
		permissions = permissions.Add(module.Burn)
	}
	if renumerate {
		permissions = permissions.Add(module.Renumerate)
	}
	if add {
		permissions = permissions.Add(module.Add)
	}
	if remove {
		permissions = permissions.Add(module.Remove)
	}
	if mutate {
		permissions = permissions.Add(module.Mutate)
	}

	return permissions
}
