package types

type Identities interface {
	GetID() ID

	Get(ID) Identity
	GetList() []Identity

	Fetch(ID) Identities
	Add(Identity) Identities
	Remove(Identity) Identities
	Mutate(Identity) Identities
}
