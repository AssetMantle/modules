package constants

import (
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

var (
	CanIssueIdentityPermission = baseIDs.NewStringID("issue")
	CanQuashIdentityPermission = baseIDs.NewStringID("quash")
)
