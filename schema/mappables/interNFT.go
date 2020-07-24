package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterNFT interface {
	types.NFT
	traits.InterChain
	traits.Burnable
	traits.Lockable
	traits.HasImmutables
	traits.HasMutables
	traits.Mappable
}
