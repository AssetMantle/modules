package types

type InterNFT interface {
	NFT

	GetChainID() ID
	GetMaintainersID() ID
	GetHashID() ID

	GetProperties() Properties

	GetLock() Height
	CanSend(Height) bool

	GetBurn() Height
	CanBurn(Height) bool
}
