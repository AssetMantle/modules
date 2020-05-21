package types

type Chains interface {
	String() string

	Chain(ID) Chain

	Add(Chain) error
	Remove(Chain) error
	Mutate(Chain) error
}
