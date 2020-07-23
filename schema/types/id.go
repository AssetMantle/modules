package types

type ID interface {
	String() string
	Bytes() []byte
	Compare(ID) int
}
