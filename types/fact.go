package types

import "encoding/json"

type Fact interface {
	String() string
	Bytes() []byte

	Signatures() Signatures
}

var _ Fact = (*BaseFact)(nil)

type BaseFact struct {
	FactBytes      []byte
	FactSignatures Signatures
}

func (baseFact BaseFact) String() string {
	bytes, Error := json.Marshal(baseFact)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}
func (baseFact BaseFact) Bytes() []byte          { return baseFact.FactBytes }
func (baseFact BaseFact) Signatures() Signatures { return baseFact.FactSignatures }
