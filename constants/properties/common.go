package properties

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/ids"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

var (
	Authentication       = base.NewProperty(ids.AuthenticationProperty, base.NewFact(base.NewListData().ZeroValue()))
	Burn                 = base.NewProperty(ids.BurnProperty, base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	Creation             = base.NewProperty(ids.CreationProperty, base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	ExchangeRate         = base.NewProperty(ids.ExchangeRateProperty, base.NewFact(base.NewDecData(sdkTypes.ZeroDec())))
	Expiry               = base.NewProperty(ids.ExpiryProperty, base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	Lock                 = base.NewProperty(ids.LockProperty, base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	MaintainedProperties = base.NewProperty(ids.MaintainedPropertiesProperty, base.NewFact(base.NewListData()))
	MakerOwnableSplit    = base.NewProperty(ids.MakerOwnableSplitProperty, base.NewFact(base.NewDecData(sdkTypes.ZeroDec())))
	NubID                = base.NewProperty(ids.NubIDProperty, base.NewFact(base.NewIDData(base.NewID(""))))
	Permissions          = base.NewProperty(ids.PermissionsProperty, base.NewFact(base.NewListData()))
	TakerID              = base.NewProperty(ids.TakerIDProperty, base.NewFact(base.NewIDData(base.NewID(""))))
	Value                = base.NewProperty(ids.ValueProperty, base.NewFact(base.NewDecData(sdkTypes.SmallestDec())))
)
