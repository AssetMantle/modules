package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type decDataI DecData

func (d decDataI) GetID() ids.DataID {
	return d.Impl.(data.DecData).GetID()
}

func (d decDataI) String() string {
	return d.Impl.(data.DecData).String()
}

func (d decDataI) Bytes() []byte {
	return d.Impl.(data.DecData).Bytes()
}

func (d decDataI) GetType() ids.StringID {
	return d.Impl.(data.DecData).GetType()
}

func (d decDataI) ZeroValue() data.Data {
	return d.Impl.(data.DecData).ZeroValue()
}

func (d decDataI) GenerateHashID() ids.HashID {
	return d.Impl.(data.DecData).GenerateHashID()
}

func (d decDataI) Compare(listable traits.Listable) int {
	return d.Impl.(data.DecData).Compare(listable)
}

func (d decDataI) Get() sdkTypes.Dec {
	return d.Impl.(data.DecData).Get()
}

var _ data.DecData = (*decDataI)(nil)
