package types

type Trait interface {
	Name() string

	ID() ID

	IsMutable() bool
}
