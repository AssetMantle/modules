package base

import (
	dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type stringDataI dataSchema.StringData

func (s stringDataI) GetID() ids.DataID {
	return s.Impl.(data.StringData).GetID()
}

func (s stringDataI) String() string {
	return s.Impl.(data.StringData).String()
}

func (s stringDataI) Bytes() []byte {
	return s.Impl.(data.StringData).Bytes()
}

func (s stringDataI) GetType() ids.StringID {
	return s.Impl.(data.StringData).GetType()
}

func (s stringDataI) ZeroValue() data.Data {
	return s.Impl.(data.StringData).ZeroValue()
}

func (s stringDataI) GenerateHashID() ids.HashID {
	return s.Impl.(data.StringData).GenerateHashID()
}

func (s stringDataI) Compare(listable traits.Listable) int {
	return s.Impl.(data.StringData).Compare(listable)
}

func (s stringDataI) Get() string {
	return s.Impl.(data.StringData).Get()
}

var _ data.StringData = (*stringDataI)(nil)
