package constants

import (
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
)

var (
	CanIssueIdentityPermission = baseIDs.NewStringID("issue")
	CanQuashIdentityPermission = baseIDs.NewStringID("quash")
)
