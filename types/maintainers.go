package types

type Maintainers interface {
	ID() ID
	Traits() Traits
	Properties() Properties

	Maintainer(ID) Maintainer

	Add(Maintainer) error
	Remove(Maintainer) error
	Mutate(Maintainer) error
}
