package constants

import (
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

var (
	CanMintAssetPermission       = baseIDs.NewStringID("mint")
	CanRenumerateAssetPermission = baseIDs.NewStringID("renumerate")
	CanBurnAssetPermission       = baseIDs.NewStringID("burn")
)
