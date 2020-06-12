package types

type Fact interface {
	String() string
	Bytes() []byte

	Signatures() Signatures
}

var _ Fact = (*BaseFact)(nil)

type BaseFact struct {
	BaseString     string
	BaseSignatures BaseSignatures
}

func (baseFact BaseFact) String() string         { return baseFact.BaseString }
func (baseFact BaseFact) Bytes() []byte          { return []byte(baseFact.BaseString) }
func (baseFact BaseFact) Signatures() Signatures { return &baseFact.BaseSignatures }
