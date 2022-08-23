// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var (
	Authentication             = base.NewMetaProperty(AuthenticationProperty.GetKey(), baseData.NewListData(baseLists.NewDataList()).ZeroValue())
	Burn                       = base.NewMetaProperty(BurnProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Expiry                     = base.NewMetaProperty(ExpiryProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	Lock                       = base.NewMetaProperty(LockProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	MaintainedClassificationID = base.NewMetaProperty(MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))))
	MaintainedProperties       = base.NewMetaProperty(MaintainedPropertiesProperty.GetKey(), baseData.NewListData(baseLists.NewDataList()))
	MakerOwnableSplit          = base.NewMetaProperty(MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(sdkTypes.ZeroDec()))
	// TODO ***** rename to name
	NubID       = base.NewMetaProperty(NubIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewStringID("")))
	Permissions = base.NewMetaProperty(PermissionsProperty.GetKey(), baseData.NewListData(baseLists.NewDataList()))
	TakerID     = base.NewMetaProperty(TakerIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewStringID("")))
	Supply      = base.NewMetaProperty(SupplyProperty.GetKey(), baseData.NewDecData(sdkTypes.SmallestDec()))
)
