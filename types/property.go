package types

type Property interface {
	String() string

	ID() ID
	Fact() Fact
}
