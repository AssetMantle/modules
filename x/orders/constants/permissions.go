package constants

import (
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
)

var (
	CanMakeOrderPermission   = baseIDs.NewStringID("make")
	CanCancelOrderPermission = baseIDs.NewStringID("cancel")
)
