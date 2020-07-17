package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Splittable interface {
	GetSplit() sdkTypes.Dec
}
