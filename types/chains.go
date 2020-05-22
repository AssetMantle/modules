package types

type Chains interface {
	String() string

	ID() ID

	Chain(ID) Chain

	Add(Chain) error
	Remove(Chain) error
	Mutate(Chain) error
}
