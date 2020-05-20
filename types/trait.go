package types

type Trait interface {
	String() string

	GetID() TraitID
	IsMutable() bool
}

type TraitID []byte
