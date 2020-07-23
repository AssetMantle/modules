package types

type Properties interface {
	Get(ID) Property

	GetList() []Property

	Add(Property) Properties
	Remove(Property) Properties
	Mutate(Property) Properties
}
