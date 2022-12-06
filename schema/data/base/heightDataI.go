package base

import (
	dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

type heightDataI dataSchema.HeightData

func (h heightDataI) GetID() ids.DataID {
	return h.Impl.(data.HeightData).GetID()
}

func (h heightDataI) String() string {
	return h.Impl.(data.HeightData).String()
}

func (h heightDataI) Bytes() []byte {
	return h.Impl.(data.HeightData).Bytes()
}

func (h heightDataI) GetType() ids.StringID {
	return h.Impl.(data.HeightData).GetType()
}

func (h heightDataI) ZeroValue() data.Data {
	return h.Impl.(data.HeightData).ZeroValue()
}

func (h heightDataI) GenerateHashID() ids.HashID {
	return h.Impl.(data.HeightData).GenerateHashID()
}

func (h heightDataI) Compare(listable traits.Listable) int {
	return h.Impl.(data.HeightData).Compare(listable)
}

func (h heightDataI) Get() types.Height {
	return h.Impl.(data.HeightData).Get()
}

var _ data.HeightData = (*heightDataI)(nil)
