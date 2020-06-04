package types

type Classification interface {
	Name() string

	ID() ID

	Traits() Traits
}
