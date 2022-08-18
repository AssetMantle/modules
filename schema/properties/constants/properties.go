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
	Authentication       = base.NewMesaProperty(AuthenticationProperty.GetKey(), baseData.NewListData(baseLists.NewDataList()).ZeroValue())
	Burn                 = base.NewMesaProperty(BurnProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Expiry               = base.NewMesaProperty(ExpiryProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Lock                 = base.NewMesaProperty(LockProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	MaintainedProperties = base.NewMesaProperty(MaintainedPropertiesProperty.GetKey(), baseData.NewListData(baseLists.NewDataList()))
	MakerOwnableSplit    = base.NewMesaProperty(MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(sdkTypes.ZeroDec()))
	// TODO ***** rename to name
	NubID       = base.NewMesaProperty(NubIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewStringID("")))
	Permissions = base.NewMesaProperty(PermissionsProperty.GetKey(), baseData.NewListData(baseLists.NewDataList()))
	TakerID     = base.NewMesaProperty(TakerIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewStringID("")))
	Supply      = base.NewMesaProperty(SupplyProperty.GetKey(), baseData.NewDecData(sdkTypes.SmallestDec()))
)
