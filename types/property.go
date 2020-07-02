package types

import "encoding/json"

type Property interface {
	String() string
	GetID() ID
	GetFact() Fact
}

var _ Property = (*property)(nil)

type property struct {
	ID   ID
	Fact Fact
}

func (property property) String() string {
	bytes, Error := json.Marshal(property)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (property property) GetID() ID     { return property.ID }
func (property property) GetFact() Fact { return property.Fact }
func NewProperty(id ID, fact Fact) Property {
	return property{
		ID:   id,
		Fact: fact,
	}
}
