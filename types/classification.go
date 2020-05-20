package types

type Classification interface {
	String() string

	ID() ID

	IsBurnMutable() bool
	IsLockMutable() bool

	Traits() Traits
}
