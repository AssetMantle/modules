package properties

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/ids"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

var (
	Authentication       = base.NewProperty(ids.Authentication, base.NewFact(base.NewListData().ZeroValue()))
	Burn                 = base.NewProperty(ids.Burn, base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	Creation             = base.NewProperty(ids.Creation, base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	ExchangeRate         = base.NewProperty(ids.ExchangeRate, base.NewFact(base.NewDecData(sdkTypes.ZeroDec())))
	Expiry               = base.NewProperty(ids.Expiry, base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	Lock                 = base.NewProperty(ids.Lock, base.NewFact(base.NewHeightData(base.NewHeight(-1))))
	MaintainedProperties = base.NewProperty(ids.MaintainedProperties, base.NewFact(base.NewListData()))
	MakerOwnableSplit    = base.NewProperty(ids.MakerOwnableSplit, base.NewFact(base.NewDecData(sdkTypes.ZeroDec())))
	NubID                = base.NewProperty(ids.NubID, base.NewFact(base.NewIDData(base.NewID(""))))
	Permissions          = base.NewProperty(ids.Permissions, base.NewFact(base.NewListData()))
	TakerID              = base.NewProperty(ids.TakerID, base.NewFact(base.NewIDData(base.NewID(""))))
	Value                = base.NewProperty(ids.Value, base.NewFact(base.NewDecData(sdkTypes.SmallestDec())))
)
