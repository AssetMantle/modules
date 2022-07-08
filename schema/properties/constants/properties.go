// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var (
	Authentication       = base.NewProperty(AuthenticationProperty, baseData.NewListData(baseLists.NewDataList()).ZeroValue())
	Burn                 = base.NewProperty(BurnProperty, baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Expiry               = base.NewProperty(ExpiryProperty, baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Lock                 = base.NewProperty(LockProperty, baseData.NewHeightData(baseTypes.NewHeight(-1)))
	MaintainedProperties = base.NewProperty(MaintainedPropertiesProperty, baseData.NewListData(baseLists.NewDataList()))
	MakerOwnableSplit    = base.NewProperty(MakerOwnableSplitProperty, baseData.NewDecData(sdkTypes.ZeroDec()))
	NubID                = base.NewProperty(NubIDProperty, baseData.NewIDData(baseIDs.NewID("")))
	Permissions          = base.NewProperty(PermissionsProperty, baseData.NewListData(baseLists.NewDataList()))
	TakerID              = base.NewProperty(TakerIDProperty, baseData.NewIDData(baseIDs.NewID("")))
	Supply               = base.NewProperty(SupplyProperty, baseData.NewDecData(sdkTypes.SmallestDec()))
)
