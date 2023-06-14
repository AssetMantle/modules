package constants

import (
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
)

var (
	CanIssueIdentityPermission  = baseIDs.NewStringID("issue")
	CanMutateIdentityPermission = baseIDs.NewStringID("mutate")
	CanQuashIdentityPermission  = baseIDs.NewStringID("quash")
)
