package types

type Owners interface {
	String() string

	ID() ID
	Traits() Traits
	Properties() Properties

	Owner(ID) Owner

	Add(Owner) error
	Remove(Owner) error
	Mutate(Owner) error
}
