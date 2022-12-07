package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.Data = (*Data)(nil)

func (x *Data) GetID() ids.DataID {
	return x.Impl.(data.Data).GetID()
}
func (x *Data) Bytes() []byte {
	return x.Impl.(data.Data).Bytes()
}
func (x *Data) GetType() ids.StringID {
	return x.Impl.(data.Data).GetType()
}
func (x *Data) ZeroValue() data.Data {
	return x.Impl.(data.Data).ZeroValue()
}
func (x *Data) GenerateHashID() ids.HashID {
	return x.Impl.(data.Data).GenerateHashID()
}
func (x *Data) Compare(listable traits.Listable) int {
	return x.Impl.(data.Data).Compare(listable)
}
