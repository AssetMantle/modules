package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Fact = (*fact)(nil)

type fact struct {
	FactString string
	Signatures types.Signatures
}

func (fact fact) String() string                  { return fact.FactString }
func (fact fact) Bytes() []byte                   { return []byte(fact.FactString) }
func (fact fact) GetSignatures() types.Signatures { return fact.Signatures }
func NewFact(factString string, signatures types.Signatures) types.Fact {
	return fact{
		FactString: factString,
		Signatures: signatures,
	}
}
