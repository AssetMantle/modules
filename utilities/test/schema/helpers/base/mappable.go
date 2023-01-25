package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/helpers"
)

// mappable struct, implements helpers.Mappable
//type *TestMappable struct {
//	ID    string
//	Value string
//}

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

func MappablePrototype() helpers.Mappable {
	return &TestMappable{}
}
