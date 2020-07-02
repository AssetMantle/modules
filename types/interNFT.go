package types

type InterNFT interface {
	NFT
	InterChain
	Burnable
	Lockable
	HasMutables
	HasImmutables
}
