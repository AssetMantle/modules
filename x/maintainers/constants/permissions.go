package constants

import (
	"github.com/AssetMantle/schema/go/ids/base"
)

var CanAddMaintainerPermission = base.NewStringID("add")
var CanMutateMaintainerPermission = base.NewStringID("mutate")
var CanRemoveMaintainerPermission = base.NewStringID("remove")
