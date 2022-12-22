package base

import (
	"github.com/AssetMantle/modules/schema/data"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.AnyData = (*AnyData)(nil)

func (x *AnyData) GetID() ids.DataID {
	return x.Impl.(data.Data).GetID()
}
func (x *AnyData) Bytes() []byte {
	return x.Impl.(data.Data).Bytes()
}
func (x *AnyData) GetType() ids.StringID {
	return x.Impl.(data.Data).GetType()
}
func (x *AnyData) ZeroValue() data.Data {
	return x.Impl.(data.Data).ZeroValue()
}
func (x *AnyData) GenerateHashID() ids.HashID {
	return x.Impl.(data.Data).GenerateHashID()
}
func (x *AnyData) ToAnyData() data.AnyData {
	return x.Impl.(data.Data).ToAnyData()
}
func (x *AnyData) Compare(listable traits.Listable) int {
	return x.Impl.(data.Data).Compare(listable)
}

func dataFromInterface(listable traits.Listable) (data.Data, error) {
	switch value := listable.(type) {
	case data.Data:
		return value, nil
	default:
		panic(errorConstants.MetaDataError)
	}
}
