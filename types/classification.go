package types

type Classification interface {
	String() string

	GetID() ID

	GetTraits() Traits
}
