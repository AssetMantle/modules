package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type Maintainer interface {
	Sting() string

	GetAddress() sdkTypes.AccAddress
	GetID() ID

	CanMutateMaintainersProperty(ID) bool

	CanAddMaintainer() bool
	CanRemoveMaintainer() bool
	CanMutateMaintainer() bool

	CanMutateLock() bool
	CanMutateBurn() bool
	CanMutateTrait(ID) bool
}
