package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Asset interface {
	String() string

	ID() ID

	Owner() sdkTypes.AccAddress
	IsOwner(sdkTypes.AccAddress) bool

	ClassificationID() ID
	Properties() Properties
	MaintainersID() ID

	GetLock() Height
	SetLock(Height) error
	CanSend(Height) bool

	GetBurn() Height
	SetBurn(Height) error
	CanBurn(Height) bool
}
