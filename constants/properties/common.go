// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/ids"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var (
	Authentication       = base.NewProperty(ids.AuthenticationProperty, baseData.NewListData().ZeroValue())
	Burn                 = base.NewProperty(ids.BurnProperty, baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Expiry               = base.NewProperty(ids.ExpiryProperty, baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Lock                 = base.NewProperty(ids.LockProperty, baseData.NewHeightData(baseTypes.NewHeight(-1)))
	MaintainedProperties = base.NewProperty(ids.MaintainedPropertiesProperty, baseData.NewListData())
	MakerOwnableSplit    = base.NewProperty(ids.MakerOwnableSplitProperty, baseData.NewDecData(sdkTypes.ZeroDec()))
	NubID                = base.NewProperty(ids.NubIDProperty, baseData.NewIDData(baseIDs.NewID("")))
	Permissions          = base.NewProperty(ids.PermissionsProperty, baseData.NewListData())
	TakerID              = base.NewProperty(ids.TakerIDProperty, baseData.NewIDData(baseIDs.NewID("")))
	Supply               = base.NewProperty(ids.SupplyProperty, baseData.NewDecData(sdkTypes.SmallestDec()))
)
