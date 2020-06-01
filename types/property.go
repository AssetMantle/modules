package types

import "encoding/json"

type Property interface {
	String() string

	ID() ID
	Fact() Fact
}

var _ Property = (*BaseProperty)(nil)

type BaseProperty struct {
	BaseID   BaseID
	BaseFact BaseFact
}

func (baseProperty BaseProperty) String() string {
	bytes, Error := json.Marshal(baseProperty)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (baseProperty BaseProperty) ID() ID     { return baseProperty.BaseID }
func (baseProperty BaseProperty) Fact() Fact { return baseProperty.BaseFact }
