package types

type Maintainers interface {
	String() string

	ID() ID
	Traits() Traits
	Properties() Properties

	Maintainer(ID) Maintainer

	AddMaintainer(Maintainer) error
	RemoveMaintainer(Maintainer) error
	MutateMaintainer(Maintainer) error
}
