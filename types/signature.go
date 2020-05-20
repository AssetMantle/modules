package types

type Signature interface {
	String() string

	ID() ID

	IsValid([]byte) bool

	HasExpired(Height) bool
}
