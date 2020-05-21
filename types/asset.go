package types

type Asset interface {
	String() string

	ID() ID

	ClassificationID() ID
	OwnersID() Owners
	MaintainersID() ID

	Properties() Properties

	GetLock() Height
	SetLock(Height) error
	CanSend(Height) bool

	GetBurn() Height
	SetBurn(Height) error
	CanBurn(Height) bool

	IsSplittable() bool
}
