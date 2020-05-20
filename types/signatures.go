package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type Signatures interface {
	String() string

	Get() []Signature
	Add(sdkTypes.AccAddress, []byte, Signature) error

	IsSigned(sdkTypes.AccAddress, []byte) bool
}
