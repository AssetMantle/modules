package base

import (
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (m *Mappable) GetKey() helpers.Key {
	return m.Impl.(helpers.Mappable).GetKey()
}

func (m *Mappable) RegisterCodec(amino *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(amino, Mappable{})
	amino.RegisterInterface((*isMappable_Impl)(nil), nil)
}

func (m *Mappable) RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface("mappable", (*helpers.Mappable)(nil), &Mappable{})
	registry.RegisterInterface("IsMappable", (*isMappable_Impl)(nil))
	//m.Impl.(helpers.Mappable).RegisterInterfaces(registry)
}
