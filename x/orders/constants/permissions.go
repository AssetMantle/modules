package constants

import (
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
)

var (
	CanMakeOrderPermission   = baseIDs.NewStringID("make")
	CanModifyOrderPermission = baseIDs.NewStringID("modify")
	CanCancelOrderPermission = baseIDs.NewStringID("cancel")
)
