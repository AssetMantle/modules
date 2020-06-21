package types

type InterNFTs interface {
	GetID() ID
	Get(ID) InterNFT

	Add(InterNFT) InterNFTs
	Remove(InterNFT) InterNFTs
	Mutate(InterNFT) InterNFTs
}
