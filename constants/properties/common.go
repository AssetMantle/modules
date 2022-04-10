// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/ids"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var (
	Authentication       = baseTypes.NewProperty(ids.AuthenticationProperty, baseData.NewListData().ZeroValue())
	Burn                 = baseTypes.NewProperty(ids.BurnProperty, baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Expiry               = baseTypes.NewProperty(ids.ExpiryProperty, baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Lock                 = baseTypes.NewProperty(ids.LockProperty, baseData.NewHeightData(baseTypes.NewHeight(-1)))
	MaintainedProperties = baseTypes.NewProperty(ids.MaintainedPropertiesProperty, baseData.NewListData())
	MakerOwnableSplit    = baseTypes.NewProperty(ids.MakerOwnableSplitProperty, baseData.NewDecData(sdkTypes.ZeroDec()))
	NubID                = baseTypes.NewProperty(ids.NubIDProperty, baseData.NewIDData(baseIDs.NewID("")))
	Permissions          = baseTypes.NewProperty(ids.PermissionsProperty, baseData.NewListData())
	TakerID              = baseTypes.NewProperty(ids.TakerIDProperty, baseData.NewIDData(baseIDs.NewID("")))
	Value                = baseTypes.NewProperty(ids.ValueProperty, baseData.NewDecData(sdkTypes.SmallestDec()))
)
