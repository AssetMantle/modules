package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type Signature interface {
	String() string

	Verify(sdkTypes.AccAddress, Trait) bool
}
