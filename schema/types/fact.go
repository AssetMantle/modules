package types

type Fact interface {
	String() string
	Bytes() []byte

	GetSignatures() Signatures
}
