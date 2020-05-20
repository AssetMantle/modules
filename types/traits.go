package types

type Traits interface {
	String() string

	Trait(ID) Trait
}
