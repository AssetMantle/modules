// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	"github.com/AssetMantle/modules/schema/data/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

// Note: Arranged alphabetically
var (
	// TODO append propertyID
	AuthenticationProperty = baseIDs.NewPropertyID(baseIDs.NewStringID("authentication"), constants.ListDataID)
	BurnProperty           = baseIDs.NewPropertyID(baseIDs.NewStringID("burn"), constants.HeightDataID)
	CreationProperty       = baseIDs.NewPropertyID(baseIDs.NewStringID("creation"), constants.HeightDataID)
	ExchangeRateProperty   = baseIDs.NewPropertyID(baseIDs.NewStringID("exchangeRate"), constants.DecDataID)
	// TODO Set max expiry as order module parameter
	ExpiryProperty               = baseIDs.NewPropertyID(baseIDs.NewStringID("expiry"), constants.HeightDataID)
	LockProperty                 = baseIDs.NewPropertyID(baseIDs.NewStringID("lock"), constants.HeightDataID)
	MaintainedPropertiesProperty = baseIDs.NewPropertyID(baseIDs.NewStringID("maintainedProperties"), constants.ListDataID)
	MakerOwnableSplitProperty    = baseIDs.NewPropertyID(baseIDs.NewStringID("makerOwnableSplit"), constants.DecDataID)
	NubIDProperty                = baseIDs.NewPropertyID(baseIDs.NewStringID("nubID"), constants.IDDataID)
	PermissionsProperty          = baseIDs.NewPropertyID(baseIDs.NewStringID("permissions"), constants.ListDataID)
	SupplyProperty               = baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), constants.DecDataID)
	TakerIDProperty              = baseIDs.NewPropertyID(baseIDs.NewStringID("takerID"), constants.IDDataID)
)
