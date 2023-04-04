package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/schema/x/helpers"
)

// mappable struct, implements helpers.Mappable
// type *TestMappable struct {
//	ID    string
//	Value string
// }

var _ helpers.Mappable = (*TestMappable)(nil)

func (t *TestMappable) GetKey() helpers.Key {
	return NewKey(t.ID)
}

func (t *TestMappable) RegisterLegacyAminoCodec(c *codec.LegacyAmino) {
	c.RegisterConcrete(&TestMappable{}, "test/*TestMappable", nil)
}

func NewMappable(id string, value string) helpers.Mappable {
	return &TestMappable{ID: id, Value: value}
}

func MappablesFromInterface(mappables []helpers.Mappable) []*TestMappable {
	Mappables := make([]*TestMappable, len(mappables))
	for index, mappable := range mappables {
		Mappables[index] = mappable.(*TestMappable)
	}
	return Mappables
}

func MappablesToInterface(mappables []*TestMappable) []helpers.Mappable {
	Mappables := make([]helpers.Mappable, len(mappables))
	for index, mappable := range mappables {
		Mappables[index] = mappable
	}
	return Mappables
}

func MappablePrototype() helpers.Mappable {
	return &TestMappable{}
}
