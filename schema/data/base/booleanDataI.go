package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.BooleanData = (*BooleanDataI)(nil)

func (b BooleanDataI) GetID() ids.DataID {
	return b.Impl.(data.BooleanData).GetID()
}

func (b BooleanDataI) Bytes() []byte {
	return b.Impl.(data.BooleanData).Bytes()
}

func (b BooleanDataI) GetType() ids.StringID {
	return b.Impl.(data.BooleanData).GetType()
}

func (b BooleanDataI) ZeroValue() data.Data {
	return b.Impl.(data.BooleanData).ZeroValue()
}

func (b BooleanDataI) GenerateHashID() ids.HashID {
	return b.Impl.(data.BooleanData).GenerateHashID()
}

func (b BooleanDataI) Compare(listable traits.Listable) int {
	return b.Impl.(data.BooleanData).Compare(listable)
}

func (b BooleanDataI) Get() bool {
	return b.Impl.(data.BooleanData).Get()
}
