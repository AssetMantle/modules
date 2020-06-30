package types

type InterNFTs interface {
	GetID() ID

	Get(ID) InterNFT
	GetList() []InterNFT

	Read(ID) InterNFTs
	Add(InterNFT) InterNFTs
	Remove(InterNFT) InterNFTs
	Mutate(InterNFT) InterNFTs
}
