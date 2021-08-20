package types

type Error interface {
	Proto
	error
}