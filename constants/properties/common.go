// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistenceSDK/constants/ids"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

var (
	Authentication       = base.NewProperty(ids.AuthenticationProperty, base.NewListData().ZeroValue())
	Burn                 = base.NewProperty(ids.BurnProperty, base.NewHeightData(base.NewHeight(-1)))
	Expiry               = base.NewProperty(ids.ExpiryProperty, base.NewHeightData(base.NewHeight(-1)))
	Lock                 = base.NewProperty(ids.LockProperty, base.NewHeightData(base.NewHeight(-1)))
	MaintainedProperties = base.NewProperty(ids.MaintainedPropertiesProperty, base.NewListData())
	MakerOwnableSplit    = base.NewProperty(ids.MakerOwnableSplitProperty, base.NewDecData(sdkTypes.ZeroDec()))
	NubID                = base.NewProperty(ids.NubIDProperty, base.NewIDData(base.NewID("")))
	Permissions          = base.NewProperty(ids.PermissionsProperty, base.NewListData())
	TakerID              = base.NewProperty(ids.TakerIDProperty, base.NewIDData(base.NewID("")))
	Value                = base.NewProperty(ids.ValueProperty, base.NewDecData(sdkTypes.SmallestDec()))
)
