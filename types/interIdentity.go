package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type InterIdentity interface {
	GetID() ID
	GetAddressList() []sdkTypes.AccAddress
	GetDeletedAddressList() []sdkTypes.AccAddress

	AddAddress(sdkTypes.AccAddress) InterIdentity
	DeleteAddress(sdkTypes.AccAddress) InterIdentity

	IsActive(sdkTypes.AccAddress) bool

	InterChain
	HasImmutables
	HasMutables
}
