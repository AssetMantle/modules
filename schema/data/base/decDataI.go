package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.DecData = (*DecDataI)(nil)

func (d *DecDataI) GetID() ids.DataID {
	return d.Impl.(data.DecData).GetID()
}
func (d *DecDataI) Bytes() []byte {
	return d.Impl.(data.DecData).Bytes()
}
func (d *DecDataI) GetType() ids.StringID {
	return d.Impl.(data.DecData).GetType()
}
func (d *DecDataI) ZeroValue() data.Data {
	return d.Impl.(data.DecData).ZeroValue()
}
func (d *DecDataI) GenerateHashID() ids.HashID {
	return d.Impl.(data.DecData).GenerateHashID()
}
func (d *DecDataI) Compare(listable traits.Listable) int {
	return d.Impl.(data.DecData).Compare(listable)
}
func (d *DecDataI) Get() sdkTypes.Dec {
	return d.Impl.(data.DecData).Get()
}
