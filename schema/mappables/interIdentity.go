package mappables

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterIdentity interface {
	GetID() types.ID
	GetProvisionedAddressList() []sdkTypes.AccAddress
	GetUnprovisionedAddressList() []sdkTypes.AccAddress

	ProvisionAddress(sdkTypes.AccAddress) InterIdentity
	UnprovisionAddress(sdkTypes.AccAddress) InterIdentity

	IsProvisioned(sdkTypes.AccAddress) bool
	IsUnprovisioned(sdkTypes.AccAddress) bool

	traits.InterChain
	traits.HasImmutables
	traits.HasMutables
	traits.Mappable
}
