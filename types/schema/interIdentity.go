package schema

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/trait"
)

type InterIdentity interface {
	GetID() ID
	GetProvisionedAddressList() []sdkTypes.AccAddress
	GetUnprovisionedAddressList() []sdkTypes.AccAddress

	ProvisionAddress(sdkTypes.AccAddress) InterIdentity
	UnprovisionAddress(sdkTypes.AccAddress) InterIdentity

	IsProvisioned(sdkTypes.AccAddress) bool
	IsUnprovisioned(sdkTypes.AccAddress) bool

	trait.InterChain
	trait.HasImmutables
	trait.HasMutables
}
