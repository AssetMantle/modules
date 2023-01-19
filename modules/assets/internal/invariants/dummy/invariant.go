package dummy

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var Invariant = func(_ sdkTypes.Context) (string, bool) {
	return "", false
}
