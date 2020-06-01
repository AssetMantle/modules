package types

type Chains interface {
	ID() ID
	Get(ID) Chain
	Add(Chain) error
	Remove(Chain) error
	Mutate(Chain) error
}
