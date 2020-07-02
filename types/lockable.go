package types

type Lockable interface {
	CanSend(Height) bool
	GetLock() Height
}
