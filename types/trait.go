package types

type Trait interface {
	String() string

	ID() ID

	IsMutable() bool
}
