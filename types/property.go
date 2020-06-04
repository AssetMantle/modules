package types

import "encoding/json"

type Property interface {
	Name() string
	ID() ID
	Fact() Fact
}

var _ Property = (*BaseProperty)(nil)

type BaseProperty struct {
	BaseID   BaseID
	BaseFact BaseFact
}

func (baseProperty BaseProperty) Name() string {
	bytes, Error := json.Marshal(baseProperty)
	if Error != nil {
		panic(Error)
	}
	return string(bytes)
}

func (baseProperty BaseProperty) ID() ID     { return baseProperty.BaseID }
func (baseProperty BaseProperty) Fact() Fact { return baseProperty.BaseFact }
