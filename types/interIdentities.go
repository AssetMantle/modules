package types

type InterIdentities interface {
	GetID() ID

	Get(ID) InterIdentity
	GetList() []InterIdentity

	Fetch(ID) InterIdentities
	Add(InterIdentity) InterIdentities
	Remove(InterIdentity) InterIdentities
	Mutate(InterIdentity) InterIdentities
}
