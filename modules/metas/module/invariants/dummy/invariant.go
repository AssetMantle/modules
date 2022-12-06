package dummy

import (
	"github.com/cosmos/cosmos-sdk/types"
)

var Invariant = func(_ types.Context) (string, bool) {
	return "", true
}
