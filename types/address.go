package types

type Address interface {
	Bytes() []byte
	String() string
}
