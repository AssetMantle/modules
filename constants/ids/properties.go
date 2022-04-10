// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

import (
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

// Note: Arranged alphabetically
var (
	AuthenticationProperty       = baseIDs.NewID("authentication")
	BurnProperty                 = baseIDs.NewID("burn")
	CreationProperty             = baseIDs.NewID("creation")
	ExchangeRateProperty         = baseIDs.NewID("exchangeRate")
	ExpiryProperty               = baseIDs.NewID("expiry")
	LockProperty                 = baseIDs.NewID("lock")
	MaintainedPropertiesProperty = baseIDs.NewID("maintainedProperties")
	MakerOwnableSplitProperty    = baseIDs.NewID("makerOwnableSplit")
	NubIDProperty                = baseIDs.NewID("nubID")
	PermissionsProperty          = baseIDs.NewID("permissions")
	TakerIDProperty              = baseIDs.NewID("takerID")
	ValueProperty                = baseIDs.NewID("value")
)
