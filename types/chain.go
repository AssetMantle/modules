package types

type Chain interface {
	String() string

	ID() ID
	Trust() Height
}
