package base

import (
	dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type idDataI dataSchema.IdData

func (i idDataI) GetID() ids.DataID {
	return i.Impl.(data.IDData).GetID()
}

func (i idDataI) String() string {
	return i.Impl.(data.IDData).String()
}

func (i idDataI) Bytes() []byte {
	return i.Impl.(data.IDData).Bytes()
}

func (i idDataI) GetType() ids.StringID {
	return i.Impl.(data.IDData).GetType()
}

func (i idDataI) ZeroValue() data.Data {
	return i.Impl.(data.IDData).ZeroValue()
}

func (i idDataI) GenerateHashID() ids.HashID {
	return i.Impl.(data.IDData).GenerateHashID()
}

func (i idDataI) Compare(listable traits.Listable) int {
	return i.Impl.(data.IDData).Compare(listable)
}

func (i idDataI) Get() ids.ID {
	return i.Impl.(data.IDData).Get()
}

var _ data.IDData = (*idDataI)(nil)
