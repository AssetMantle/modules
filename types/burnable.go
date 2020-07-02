package types

type Burnable interface {
	CanBurn(Height) bool
	GetBurn() Height
}
