package types

type Fact interface {
	String() string
	Bytes() []byte

	GetSignatures() Signatures
}

var _ Fact = (*fact)(nil)

type fact struct {
	FactString string
	Signatures Signatures
}

func (fact fact) String() string            { return fact.FactString }
func (fact fact) Bytes() []byte             { return []byte(fact.FactString) }
func (fact fact) GetSignatures() Signatures { return fact.Signatures }
func NewFact(factString string, signatures Signatures) Fact {
	return fact{
		FactString: factString,
		Signatures: signatures,
	}
}
