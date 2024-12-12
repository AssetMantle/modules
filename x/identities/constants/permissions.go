// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

var (
	CanIssueIdentityPermission = baseIDs.NewStringID("issue")
	CanQuashIdentityPermission = baseIDs.NewStringID("quash")
)
