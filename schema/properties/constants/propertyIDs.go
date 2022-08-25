// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	"github.com/AssetMantle/modules/schema/data/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

// Note: Arranged alphabetically
var (
	AuthenticationPropertyID = baseIDs.NewPropertyID(baseIDs.NewStringID("authentication"), constants.ListDataID)
	BurnHeightPropertyID     = baseIDs.NewPropertyID(baseIDs.NewStringID("burn"), constants.HeightDataID)
	CreationHeightPropertyID = baseIDs.NewPropertyID(baseIDs.NewStringID("creationHeight"), constants.HeightDataID)
	ExchangeRatePropertyID   = baseIDs.NewPropertyID(baseIDs.NewStringID("exchangeRate"), constants.DecDataID)
	// TODO Set max expiry as order module parameter
	ExpiryHeightPropertyID               = baseIDs.NewPropertyID(baseIDs.NewStringID("expiryHeight"), constants.HeightDataID)
	IdentityIDPropertyID                 = baseIDs.NewPropertyID(baseIDs.NewStringID("identityID"), constants.IDDataID)
	LockPropertyID                       = baseIDs.NewPropertyID(baseIDs.NewStringID("lock"), constants.HeightDataID)
	MaintainedClassificationIDPropertyID = baseIDs.NewPropertyID(baseIDs.NewStringID("maintainedClassificationID"), constants.IDDataID)
	MaintainedPropertiesPropertyID       = baseIDs.NewPropertyID(baseIDs.NewStringID("maintainedProperties"), constants.ListDataID)
	MakerIDPropertyID                    = baseIDs.NewPropertyID(baseIDs.NewStringID("makerID"), constants.IDDataID)
	MakerOwnableIDPropertyID             = baseIDs.NewPropertyID(baseIDs.NewStringID("makerOwnableID"), constants.IDDataID)
	MakerOwnableSplitPropertyID          = baseIDs.NewPropertyID(baseIDs.NewStringID("makerOwnableSplit"), constants.DecDataID)
	NubIDPropertyID                      = baseIDs.NewPropertyID(baseIDs.NewStringID("nubID"), constants.IDDataID)
	PermissionsPropertyID                = baseIDs.NewPropertyID(baseIDs.NewStringID("permissions"), constants.ListDataID)
	SupplyPropertyID                     = baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), constants.DecDataID)
	TakerIDPropertyID                    = baseIDs.NewPropertyID(baseIDs.NewStringID("takerID"), constants.IDDataID)
	TakerOwnableIDPropertyID             = baseIDs.NewPropertyID(baseIDs.NewStringID("takerOwnableID"), constants.IDDataID)
)
