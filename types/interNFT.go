package types

type InterNFT interface {
	NFT

	ID() ID

	ChainID() ID
	ClassificationID() ID
	MaintainersID() ID
	HashID() ID

	Properties() Properties

	GetLock() Height
	CanSend(Height) bool

	GetBurn() Height
	CanBurn(Height) bool
}
