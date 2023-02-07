package base

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

var _ helpers.Mappable = (*Mappable_DataMappable)(nil)

func (m *Mappable_DataMappable) GetKey() helpers.Key {
	return NewKey(base.GenerateDataID(m.DataMappable.Data))
}

func (m *Mappable_DataMappable) RegisterCodec(amino *codec.LegacyAmino) {

}

func (m *Mappable_DataMappable) RegisterInterfaces(registry types.InterfaceRegistry) {
}

func NewDataMappable(data data.Data) helpers.Mappable {
	return &Mappable{
		Impl: &Mappable_DataMappable{
			DataMappable: &DataMappable{
				Data: data.(*baseData.Data),
			},
		},
	}
}

func DataMappablePrototype() helpers.Mappable {
	return &Mappable{}
}
