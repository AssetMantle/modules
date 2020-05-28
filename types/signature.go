package types

type Signature interface {
	String() string

	ID() ID

	IsValid([]byte) bool
	HasExpired(Height) bool
}

type BaseSignature struct {
	BaseBytes []byte
}

var _ Signature = (*BaseSignature)(nil)

func (baseSignature BaseSignature) String() string         {}
func (baseSignature BaseSignature) ID() ID                 {}
func (baseSignature BaseSignature) IsValid([]byte) bool    {}
func (baseSignature BaseSignature) HasExpired(Height) bool {}
