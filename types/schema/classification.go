package schema

type Classification interface {
	String() string
	GetID() ID
	GetTraits() Traits
}
