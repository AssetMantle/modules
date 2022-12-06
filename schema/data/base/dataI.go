package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.Data = (*DataI)(nil)

func (x *DataI) GetID() ids.DataID {
	return x.Impl.(data.Data).GetID()
}
func (x *DataI) Bytes() []byte {
	return x.Impl.(data.Data).Bytes()
}
func (x *DataI) GetType() ids.StringID {
	return x.Impl.(data.Data).GetType()
}
func (x *DataI) ZeroValue() data.Data {
	return x.Impl.(data.Data).ZeroValue()
}
func (x *DataI) GenerateHashID() ids.HashID {
	return x.Impl.(data.Data).GenerateHashID()
}
func (x *DataI) Compare(listable traits.Listable) int {
	return x.Impl.(data.Data).Compare(listable)
}
