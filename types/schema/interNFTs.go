package schema

type InterNFTs interface {
	GetID() ID

	Get(ID) InterNFT
	GetList() []InterNFT

	Fetch(ID) InterNFTs
	Add(InterNFT) InterNFTs
	Remove(InterNFT) InterNFTs
	Mutate(InterNFT) InterNFTs
}
