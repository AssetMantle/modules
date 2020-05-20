package types

type Signatures interface {
	String() string

	Signature(ID) Signature

	Add(Signature) error
	Remove(Signature) error
	Mutate(Signature) error
}
