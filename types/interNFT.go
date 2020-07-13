package types

type InterNFT interface {
	NFT
	InterChain
	Burnable
	Lockable
	HasImmutables
	HasMutables
}
