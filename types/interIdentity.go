package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type InterIdentity interface {
	GetID() ID
	GetProvisionedAddressList() []sdkTypes.AccAddress
	GetUnprovisionedAddressList() []sdkTypes.AccAddress

	ProvisionAddress(sdkTypes.AccAddress) InterIdentity
	UnprovisionAddress(sdkTypes.AccAddress) InterIdentity

	IsProvisioned(sdkTypes.AccAddress) bool
	IsUnprovisioned(sdkTypes.AccAddress) bool

	InterChain
	HasImmutables
	HasMutables
}
