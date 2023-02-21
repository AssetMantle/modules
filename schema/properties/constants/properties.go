// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var (
	AuthenticationProperty  = base.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData())
	BondAmountProperty      = base.NewMetaProperty(baseIDs.NewStringID("bondAmount"), baseData.NewNumberData(0))
	BondRateProperty        = base.NewMetaProperty(baseIDs.NewStringID("bondRate"), baseData.NewNumberData(0))
	BurnEnabledProperty     = base.NewMetaProperty(baseIDs.NewStringID("burnEnabled"), baseData.NewBooleanData(false))
	BurnHeightProperty      = base.NewMetaProperty(baseIDs.NewStringID("burnHeight"), baseData.NewHeightData(baseTypes.NewHeight(-1)).ZeroValue())
	DeputizeAllowedProperty = base.NewMetaProperty(baseIDs.NewStringID("deputizeAllowed"), baseData.NewBooleanData(false))
	// TODO check default value
	CreationHeightProperty = base.NewMetaProperty(baseIDs.NewStringID("creationHeight"), baseData.NewHeightData(baseTypes.NewHeight(-1)).ZeroValue())
	// TODO check default value
	ExchangeRateProperty               = base.NewMetaProperty(baseIDs.NewStringID("exchangeRate"), baseData.NewDecData(sdkTypes.SmallestDec()).ZeroValue())
	ExpiryHeightProperty               = base.NewMetaProperty(baseIDs.NewStringID("expiryHeight"), baseData.NewHeightData(baseTypes.NewHeight(-1)).ZeroValue())
	LockProperty                       = base.NewMetaProperty(baseIDs.NewStringID("lock"), baseData.NewHeightData(baseTypes.NewHeight(-1)).ZeroValue())
	IdentityIDProperty                 = base.NewMetaProperty(baseIDs.NewStringID("identityID"), baseData.NewIDData(baseIDs.PrototypeIdentityID()))
	MaintainedClassificationIDProperty = base.NewMetaProperty(baseIDs.NewStringID("maintainedClassificationID"), baseData.NewIDData(baseIDs.PrototypeClassificationID()))
	MaintainedPropertiesProperty       = base.NewMetaProperty(baseIDs.NewStringID("maintainedProperties"), baseData.NewListData())
	// TODO check default value
	MakerIDProperty = base.NewMetaProperty(baseIDs.NewStringID("makerID"), baseData.NewIDData(baseIDs.PrototypeIdentityID()))
	// TODO check default value
	MakerOwnableIDProperty           = base.NewMetaProperty(baseIDs.NewStringID("makerOwnableID"), baseData.NewIDData(baseIDs.PrototypeOwnableID()))
	MakerOwnableSplitProperty        = base.NewMetaProperty(baseIDs.NewStringID("makerOwnableSplit"), baseData.NewDecData(sdkTypes.SmallestDec()))
	MaxPropertyCountProperty         = base.NewMetaProperty(baseIDs.NewStringID("maxPropertyCount"), baseData.NewNumberData(0))
	MaxProvisionAddressCountProperty = base.NewMetaProperty(baseIDs.NewStringID("maxProvisionAddressCount"), baseData.NewNumberData(0))
	MaxOrderLifeProperty             = base.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	MintEnabledProperty              = base.NewMetaProperty(baseIDs.NewStringID("mintEnabled"), baseData.NewBooleanData(false))

	// TODO ***** rename to name
	NubIDProperty             = base.NewMetaProperty(baseIDs.NewStringID("nubID"), baseData.NewIDData(baseIDs.PrototypeStringID()))
	PermissionsProperty       = base.NewMetaProperty(baseIDs.NewStringID("permissions"), baseData.NewListData())
	RenumerateEnabledProperty = base.NewMetaProperty(baseIDs.NewStringID("renumerateEnabled"), baseData.NewBooleanData(false))
	RevealEnabledProperty     = base.NewMetaProperty(baseIDs.NewStringID("revealEnabled"), baseData.NewBooleanData(false))
	TakerIDProperty           = base.NewMetaProperty(baseIDs.NewStringID("takerID"), baseData.NewIDData(baseIDs.PrototypeStringID()))
	// TODO check default value
	TakerOwnableIDProperty   = base.NewMetaProperty(baseIDs.NewStringID("takerOwnableID"), baseData.NewIDData(baseIDs.PrototypeOwnableID()))
	SupplyProperty           = base.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.SmallestDec()))
	WrapAllowedCoinsProperty = base.NewMetaProperty(baseIDs.NewStringID("wrapAllowedCoins"), baseData.NewListData())
)
