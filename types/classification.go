package types

type Classification interface {
	Name() string

	ID() ID

	IsBurnMutable() bool
	IsLockMutable() bool

	IsSplittable() bool

	Traits() Traits
}
