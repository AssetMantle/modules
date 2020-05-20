package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Asset interface {
	String() string

	Address() Address

	Issuer() sdkTypes.AccAddress
	MutateIssuer(sdkTypes.AccAddress) error
	IsIssuer(sdkTypes.AccAddress) bool

	Owner() sdkTypes.AccAddress
	MutateOwner(sdkTypes.AccAddress) error
	IsOwner(sdkTypes.AccAddress) bool

	Properties() Properties

	GetLock() int
	SetLock(int)
	CanSend() bool

	GetBurn() int
	SetBurn(int)
	CanBurn() bool
}
