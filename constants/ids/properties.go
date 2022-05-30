// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

import (
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

// Note: Arranged alphabetically
var (
	AuthenticationProperty = baseIDs.NewPropertyID(baseIDs.NewID("authentication"), constants.ListDataID)
	BurnProperty           = baseIDs.NewPropertyID(baseIDs.NewID("burn"), constants.HeightDataID)
	CreationProperty       = baseIDs.NewPropertyID(baseIDs.NewID("creation"), constants.HeightDataID)
	ExchangeRateProperty   = baseIDs.NewPropertyID(baseIDs.NewID("exchangeRate"), constants.DecDataID)
	// TODO Set max expirey as order module parameter
	ExpiryProperty               = baseIDs.NewPropertyID(baseIDs.NewID("expiry"), constants.HeightDataID)
	LockProperty                 = baseIDs.NewPropertyID(baseIDs.NewID("lock"), constants.HeightDataID)
	MaintainedPropertiesProperty = baseIDs.NewPropertyID(baseIDs.NewID("maintainedProperties"), constants.ListDataID)
	MakerOwnableSplitProperty    = baseIDs.NewPropertyID(baseIDs.NewID("makerOwnableSplit"), constants.DecDataID)
	NubIDProperty                = baseIDs.NewPropertyID(baseIDs.NewID("nubID"), constants.IDDataID)
	PermissionsProperty          = baseIDs.NewPropertyID(baseIDs.NewID("permissions"), constants.ListDataID)
	SupplyProperty               = baseIDs.NewPropertyID(baseIDs.NewID("supply"), constants.DecDataID)
	TakerIDProperty              = baseIDs.NewPropertyID(baseIDs.NewID("takerID"), constants.IDDataID)
)
