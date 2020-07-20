package schema

type Trait interface {
	String() string

	GetID() ID

	IsMutable() bool
}
