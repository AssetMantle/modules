package traits

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type Exchangeable interface {
}

var _ Exchangeable = (*sdkTypes.Coin)(nil)
