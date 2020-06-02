package types

type InterNFT interface {
	NFT

	ChainID() ID
	MaintainersID() ID
	HashID() ID

	Properties() Properties

	GetLock() Height
	CanSend(Height) bool

	GetBurn() Height
	CanBurn(Height) bool
}
