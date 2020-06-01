package types

type Property interface {
	ID() ID
	Fact() Fact
}

var _ Property = (*BaseProperty)(nil)

type BaseProperty struct {
	BaseID   BaseID
	BaseFact BaseFact
}

func (baseProperty BaseProperty) ID() ID     { return baseProperty.BaseID }
func (baseProperty BaseProperty) Fact() Fact { return baseProperty.BaseFact }
