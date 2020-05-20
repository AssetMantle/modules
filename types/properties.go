package types

type Properties interface {
	String() string

	Mutate(Property) error
}
