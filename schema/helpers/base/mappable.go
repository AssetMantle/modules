package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

var _ helpers.Mappable = (*Mappable)(nil)

func (m *Mappable) GetKey() helpers.Key {
	x := m.Impl
	y := x.(helpers.Mappable)
	return y.GetKey()
}

func (m *Mappable) RegisterCodec(amino *codec.LegacyAmino) {
}

func (m *Mappable) RegisterInterfaces(registry types.InterfaceRegistry) {
	m.Impl.(helpers.Mappable).RegisterInterfaces(registry)
}

func NewDataMappable(data data.Data) helpers.Mappable {
	return &Mappable{
		Impl: &Mappable_Mappable{
			Mappable: data.(*base.Data),
		},
	}
}
