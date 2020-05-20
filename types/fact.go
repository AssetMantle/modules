package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type Fact interface {
	String() string

	Get() []byte

	IsSigned(sdkTypes.AccAddress) bool
	AddSignature(sdkTypes.AccAddress, Signature) error
}
