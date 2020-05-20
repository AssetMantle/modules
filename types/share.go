package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type Share interface {
	ID() ID
	GetOwner() sdkTypes.AccAddress
	SetOwner(sdkTypes.AccAddress)
	GetLock() bool
	SetLock(bool)
	String() string
}
