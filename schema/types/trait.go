package types

type Trait interface {
	String() string

	GetID() ID

	GetProperty() Property

	IsMutable() bool
}
