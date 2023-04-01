package constansts

import (
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

var (
	Mint       = baseIDs.NewStringID("mint")
	Burn       = baseIDs.NewStringID("burn")
	Renumerate = baseIDs.NewStringID("renumerate")
	Add        = baseIDs.NewStringID("add")
	Remove     = baseIDs.NewStringID("remove")
	Mutate     = baseIDs.NewStringID("mutate")
)
