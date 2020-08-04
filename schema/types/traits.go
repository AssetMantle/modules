package types

type Traits interface {
	Get(ID) Trait
	GetList() []Trait

	Add(Trait) Traits
	Remove(Trait) Traits
	Mutate(Trait) Traits
}
