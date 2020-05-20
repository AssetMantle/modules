package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Asset interface {
	String() string

	ID() ID

	Owner() sdkTypes.AccAddress
	MutateOwner(sdkTypes.AccAddress) error
	IsOwner(sdkTypes.AccAddress) bool

	ClassificationID() ID
	Properties() Properties
	MaintainersID() ID

	GetLock() int
	SetLock(int)
	CanSend() bool

	GetBurn() int
	SetBurn(int)
	CanBurn() bool
}
