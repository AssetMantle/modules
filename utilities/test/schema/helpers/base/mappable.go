package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/helpers"
)

// mappable struct, implements helpers.Mappable
type testMappable struct {
	ID    string
	Value string
}

var _ helpers.Mappable = (*testMappable)(nil)

func (t testMappable) GetKey() helpers.Key {
	return NewKey(t.ID)
}

func (t testMappable) RegisterLegacyAminoCodec(c *codec.LegacyAmino) {
	c.RegisterConcrete(testMappable{}, "test/testMappable", nil)
}

func NewMappable(id string, value string) helpers.Mappable {
	return testMappable{ID: id, Value: value}
}

func MappablePrototype() helpers.Mappable {
	return testMappable{}
}
