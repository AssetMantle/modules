package base

import (
	dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type booleanDataI dataSchema.BooleanData

func (b booleanDataI) GetID() ids.DataID {
	return b.Impl.(data.BooleanData).GetID()
}

func (b booleanDataI) String() string {
	return b.Impl.(data.BooleanData).String()
}

func (b booleanDataI) Bytes() []byte {
	return b.Impl.(data.BooleanData).Bytes()
}

func (b booleanDataI) GetType() ids.StringID {
	return b.Impl.(data.BooleanData).GetType()
}

func (b booleanDataI) ZeroValue() data.Data {
	return b.Impl.(data.BooleanData).ZeroValue()
}

func (b booleanDataI) GenerateHashID() ids.HashID {
	return b.Impl.(data.BooleanData).GenerateHashID()
}

func (b booleanDataI) Compare(listable traits.Listable) int {
	return b.Impl.(data.BooleanData).Compare(listable)
}

func (b booleanDataI) Get() bool {
	return b.Impl.(data.BooleanData).Get()
}

var _ data.BooleanData = &booleanDataI{}
