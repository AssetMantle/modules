package types

import (
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ mappables.Meta = (*Meta)(nil)

func (meta *Meta) GetData() types.Data {
	data, ok := meta.Data.GetCachedValue().(types.Data)
	if !ok {
		return nil
	}
	return data
}

func (meta *Meta) setData(data types.Data) error {
	if data == nil {
		meta.Data = codecTypes.Any{}
		return nil
	}
	any, err := codecTypes.NewAnyWithValue(data)
	if err == nil {
		meta.Data = *any
	}
	return err
}

func (meta *Meta) GetID() types.ID {
	return &meta.ID
}

func (meta Meta) GetMetaID() MetaID {
	return meta.ID
}

func (meta *Meta) GetKey() helpers.Key {
	return meta.GetKey()
}

func (meta Meta) UnpackInterfaces(unpacker codecTypes.AnyUnpacker) error {
	var data types.Data
	return unpacker.UnpackAny(&meta.Data, &data)
}

func NewMeta(id MetaID, data types.Data) Meta {
	meta := Meta{ID: id}
	err := meta.setData(data)
	if err != nil {
		panic(err)
	}
	return meta
}
