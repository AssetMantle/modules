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
	AuthenticationProperty = base.NewMetaProperty(AuthenticationPropertyID.GetKey(), baseData.NewListData(baseLists.NewDataList()).ZeroValue())
	BurnHeightProperty     = base.NewMetaProperty(BurnHeightPropertyID.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	// TODO check default value
	CreationHeightProperty = base.NewMetaProperty(CreationHeightPropertyID.GetType(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	// TODO check default value
	ExchangeRateProperty               = base.NewMetaProperty(ExchangeRatePropertyID.GetKey(), baseData.NewDecData(sdkTypes.NewDec(0)))
	ExpiryHeightProperty               = base.NewMetaProperty(ExpiryHeightPropertyID.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	LockProperty                       = base.NewMetaProperty(LockPropertyID.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(-1)))
	MaintainedClassificationIDProperty = base.NewMetaProperty(MaintainedClassificationIDPropertyID.GetKey(), baseData.NewIDData(baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))))
	MaintainedPropertiesProperty       = base.NewMetaProperty(MaintainedPropertiesPropertyID.GetKey(), baseData.NewListData(baseLists.NewDataList()))
	// TODO check default value
	MakerIDProperty = base.NewMetaProperty(MakerIDPropertyID.GetKey(), baseData.NewIDData(baseIDs.NewIdentityID(nil, nil)))
	// TODO check default value
	MakerOwnableIDProperty    = base.NewMetaProperty(MakerOwnableIDPropertyID.GetKey(), baseData.NewIDData(baseIDs.NewOwnableID(baseIDs.NewStringID(""))))
	MakerOwnableSplitProperty = base.NewMetaProperty(MakerOwnableSplitPropertyID.GetKey(), baseData.NewDecData(sdkTypes.ZeroDec()))
	// TODO ***** rename to name
	NubIDProperty       = base.NewMetaProperty(NubIDPropertyID.GetKey(), baseData.NewIDData(baseIDs.NewStringID("")))
	PermissionsProperty = base.NewMetaProperty(PermissionsPropertyID.GetKey(), baseData.NewListData(baseLists.NewDataList()))
	TakerIDProperty     = base.NewMetaProperty(TakerIDPropertyID.GetKey(), baseData.NewIDData(baseIDs.NewStringID("")))
	// TODO check default value
	TakerOwnableIDProperty = base.NewMetaProperty(TakerOwnableIDPropertyID.GetKey(), baseData.NewIDData(baseIDs.NewOwnableID(baseIDs.NewStringID(""))))
	SupplyProperty         = base.NewMetaProperty(SupplyPropertyID.GetKey(), baseData.NewDecData(sdkTypes.SmallestDec()))
)
