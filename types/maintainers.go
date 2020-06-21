package types

type Maintainers interface {
	ID() ID

	Get(ID) Maintainer

	Add(Maintainer) Maintainers
	Remove(Maintainer) Maintainers
	Mutate(Maintainer) Maintainers
}
