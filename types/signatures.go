package types

type Signatures interface {
	String() string

	Signature(ID) Signature

	Add(Signature) error
	Remove(Signature) error
	Mutate(Signature) error
}

type BaseSignatures struct {
	BaseSignatureList []BaseSignature
}

var _ Signatures = (*BaseSignatures)(nil)

func (baseSignatures BaseSignatures) String() string         {}
func (baseSignatures BaseSignatures) Signature(ID) Signature {}
func (baseSignatures BaseSignatures) Add(Signature) error    {}
func (baseSignatures BaseSignatures) Remove(Signature) error {}
func (baseSignatures BaseSignatures) Mutate(Signature) error {}
