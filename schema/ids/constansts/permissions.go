package constansts

import (
	"github.com/AssetMantle/modules/schema/ids/base"
)

// TODO impl through enums
var (
	Mint       = base.NewStringID("mint")
	Burn       = base.NewStringID("burn")
	Renumerate = base.NewStringID("renumerate")
	Add        = base.NewStringID("add")
	Remove     = base.NewStringID("remove")
	Mutate     = base.NewStringID("mutate")
)
