package base

import (
	"github.com/AssetMantle/modules/schema/data"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type getter interface {
	get() data.Data
}

func (x *AnyData_AccAddressData) get() data.Data {
	return x.AccAddressData
}
func (x *AnyData_BooleanData) get() data.Data {
	return x.BooleanData
}
func (x *AnyData_DecData) get() data.Data {
	return x.DecData
}
func (x *AnyData_HeightData) get() data.Data {
	return x.HeightData
}
func (x *AnyData_IDData) get() data.Data {
	return x.IDData
}
func (x *AnyData_StringData) get() data.Data {
	return x.StringData
}
func (x *AnyData_ListData) get() data.Data {
	return &x.ListData
}

var _ data.AnyData = (*AnyData)(nil)

func (x *AnyData) GetID() ids.DataID {
	return x.Impl.(getter).get().GetID()
}
func (x *AnyData) Bytes() []byte {
	return x.Impl.(getter).get().Bytes()
}
func (x *AnyData) GetType() ids.StringID {
	return x.Impl.(getter).get().GetType()
}
func (x *AnyData) ZeroValue() data.Data {
	return x.Impl.(getter).get().ZeroValue()
}
func (x *AnyData) GenerateHashID() ids.HashID {
	return x.Impl.(getter).get().GenerateHashID()
}
func (x *AnyData) ToAnyData() data.AnyData {
	return x.Impl.(getter).get().ToAnyData()
}
func (x *AnyData) Compare(listable traits.Listable) int {
	return x.Impl.(getter).get().Compare(listable)
}

func dataFromInterface(listable traits.Listable) (data.Data, error) {
	switch value := listable.(type) {
	case data.Data:
		return value, nil
	default:
		panic(errorConstants.MetaDataError)
	}
}
