package types

type Properties interface {
	String() string

	Property(ID) Property

	AddProperty(Property) error
	RemoveProperty(Property) error
	MutateProperty(Property) error
}
