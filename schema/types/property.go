package types

type Property interface {
	String() string
	GetID() ID
	GetFact() Fact
}
