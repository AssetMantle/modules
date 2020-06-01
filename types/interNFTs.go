package types

type InterNFTs interface {
	ID() ID
	Get(ID) InterNFT

	Add(InterNFT) error
	Remove(InterNFT) error
	Mutate(InterNFT) error
}
