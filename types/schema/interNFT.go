package schema

import (
	"github.com/persistenceOne/persistenceSDK/types/trait"
)

type InterNFT interface {
	NFT
	trait.InterChain
	trait.Burnable
	trait.Lockable
	trait.HasImmutables
	trait.HasMutables
}
