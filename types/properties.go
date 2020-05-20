package types

type Properties interface {
	String() string

	AddProperty(Property) error
	RemoveProperty(Property) error
	MutateProperty(Property) error
}
