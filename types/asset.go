package types

type Asset interface {
	String() string

	ID() ID

	ChainID() ID
	ClassificationID() ID
	HashID() ID

	OwnersID() ID
	MaintainersID() ID

	Properties() Properties

	GetLock() Height
	CanSend(Height) bool

	GetBurn() Height
	CanBurn(Height) bool

	MutateProperties(Properties) error
	MutateLock(Height) error
	MutateBurn(Height) error
}
