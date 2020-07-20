package schema

type Chains interface {
	GetID() ID
	Get(ID) Chain
	Add(Chain) Chains
	Remove(Chain) Chains
	Mutate(Chain) Chains
}
