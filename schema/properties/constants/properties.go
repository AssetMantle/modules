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
	Authentication       = base.NewProperty(AuthenticationProperty.GetKey(), baseData.NewListData().ZeroValue())
	Burn                 = base.NewProperty(BurnProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Expiry               = base.NewProperty(ExpiryProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Lock                 = base.NewProperty(LockProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	MaintainedProperties = base.NewProperty(MaintainedPropertiesProperty.GetKey(), baseData.NewListData())
	MakerOwnableSplit    = base.NewProperty(MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(sdkTypes.ZeroDec()))
	NubID                = base.NewProperty(NubIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewID("")))
	TakerID              = base.NewProperty(TakerIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewID("")))
	Supply               = base.NewProperty(SupplyProperty.GetKey(), baseData.NewDecData(sdkTypes.SmallestDec()))
)
