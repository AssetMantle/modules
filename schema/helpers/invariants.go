package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Invariants interface {
	RegisterInvariants(sdkTypes.InvariantRegistry)
}
