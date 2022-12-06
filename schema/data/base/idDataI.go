package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.IDData = (*IdDataI)(nil)

func (i *IdDataI) GetID() ids.DataID {
	return i.Impl.(data.IDData).GetID()
}
func (i *IdDataI) Bytes() []byte {
	return i.Impl.(data.IDData).Bytes()
}
func (i *IdDataI) GetType() ids.StringID {
	return i.Impl.(data.IDData).GetType()
}
func (i *IdDataI) ZeroValue() data.Data {
	return i.Impl.(data.IDData).ZeroValue()
}
func (i *IdDataI) GenerateHashID() ids.HashID {
	return i.Impl.(data.IDData).GenerateHashID()
}
func (i *IdDataI) Compare(listable traits.Listable) int {
	return i.Impl.(data.IDData).Compare(listable)
}
func (i *IdDataI) Get() ids.ID {
	return i.Impl.(data.IDData).Get()
}
