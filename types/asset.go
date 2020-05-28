package types

type Asset interface {
	String() string

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

	MutateProperties(Properties) error
	MutateLock(Height) error
	MutateBurn(Height) error
}
