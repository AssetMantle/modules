package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ data.AccAddressData = (*AccAddressDataI)(nil)

func (a AccAddressDataI) GetID() ids.DataID {
	return a.Impl.(data.AccAddressData).GetID()
}

func (a AccAddressDataI) Bytes() []byte {
	return a.Impl.(data.AccAddressData).Bytes()
}

func (a AccAddressDataI) GetType() ids.StringID {
	return a.Impl.(data.AccAddressData).GetType()
}

func (a AccAddressDataI) ZeroValue() data.Data {
	return a.Impl.(data.AccAddressData).ZeroValue()
}

func (a AccAddressDataI) GenerateHashID() ids.HashID {
	return a.Impl.(data.AccAddressData).GenerateHashID()
}

func (a AccAddressDataI) Compare(listable traits.Listable) int {
	return a.Impl.(data.AccAddressData).Compare(listable)
}

func (a AccAddressDataI) Get() sdkTypes.AccAddress {
	return a.Impl.(data.AccAddressData).Get()
}
