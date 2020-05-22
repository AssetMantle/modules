package types

type ID interface {
	String() string

	Bytes() []byte

	IsEqualTo(ID) bool
}
