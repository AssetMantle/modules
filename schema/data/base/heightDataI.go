package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

var _ data.HeightData = (*HeightDataI)(nil)

func (h *HeightDataI) GetID() ids.DataID {
	return h.Impl.(data.HeightData).GetID()
}
func (h *HeightDataI) Bytes() []byte {
	return h.Impl.(data.HeightData).Bytes()
}
func (h *HeightDataI) GetType() ids.StringID {
	return h.Impl.(data.HeightData).GetType()
}
func (h *HeightDataI) ZeroValue() data.Data {
	return h.Impl.(data.HeightData).ZeroValue()
}
func (h *HeightDataI) GenerateHashID() ids.HashID {
	return h.Impl.(data.HeightData).GenerateHashID()
}
func (h *HeightDataI) Compare(listable traits.Listable) int {
	return h.Impl.(data.HeightData).Compare(listable)
}
func (h *HeightDataI) Get() types.Height {
	return h.Impl.(data.HeightData).Get()
}
