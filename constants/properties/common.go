// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/ids"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/types/base"
)

var (
	Authentication       = base.NewProperty(ids.AuthenticationProperty, baseData.NewListData().ZeroValue())
	Burn                 = base.NewProperty(ids.BurnProperty, baseData.NewHeightData(base.NewHeight(-1)))
	Expiry               = base.NewProperty(ids.ExpiryProperty, baseData.NewHeightData(base.NewHeight(-1)))
	Lock                 = base.NewProperty(ids.LockProperty, baseData.NewHeightData(base.NewHeight(-1)))
	MaintainedProperties = base.NewProperty(ids.MaintainedPropertiesProperty, baseData.NewListData())
	MakerOwnableSplit    = base.NewProperty(ids.MakerOwnableSplitProperty, baseData.NewDecData(sdkTypes.ZeroDec()))
	NubID                = base.NewProperty(ids.NubIDProperty, baseData.NewIDData(base.NewID("")))
	Permissions          = base.NewProperty(ids.PermissionsProperty, baseData.NewListData())
	TakerID              = base.NewProperty(ids.TakerIDProperty, baseData.NewIDData(base.NewID("")))
	Value                = base.NewProperty(ids.ValueProperty, baseData.NewDecData(sdkTypes.SmallestDec()))
)
