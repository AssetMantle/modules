package types

type Chain interface {
	Name() string
	ID() ID
	Trust() Height
}
