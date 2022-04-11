// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

import (
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

// Note: Arranged alphabetically
var (
	AuthenticationProperty = baseIDs.NewPropertyID(baseIDs.NewID("authentication"), ListDataID)
	BurnProperty           = baseIDs.NewPropertyID(baseIDs.NewID("burn"), HeightDataID)
	CreationProperty       = baseIDs.NewPropertyID(baseIDs.NewID("creation"), HeightDataID)
	ExchangeRateProperty   = baseIDs.NewPropertyID(baseIDs.NewID("exchangeRate"), DecDataID)
	// TODO Set max expirey as order module parameter
	ExpiryProperty               = baseIDs.NewPropertyID(baseIDs.NewID("expiry"), HeightDataID)
	LockProperty                 = baseIDs.NewPropertyID(baseIDs.NewID("lock"), HeightDataID)
	MaintainedPropertiesProperty = baseIDs.NewPropertyID(baseIDs.NewID("maintainedProperties"), ListDataID)
	MakerOwnableSplitProperty    = baseIDs.NewPropertyID(baseIDs.NewID("makerOwnableSplit"), DecDataID)
	NubIDProperty                = baseIDs.NewPropertyID(baseIDs.NewID("nubID"), IDDataID)
	PermissionsProperty          = baseIDs.NewPropertyID(baseIDs.NewID("permissions"), ListDataID)
	SupplyProperty               = baseIDs.NewPropertyID(baseIDs.NewID("supply"), DecDataID)
	TakerIDProperty              = baseIDs.NewPropertyID(baseIDs.NewID("takerID"), IDDataID)
)
