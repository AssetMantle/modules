package base

import (
	dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type accAddressDataI dataSchema.AccAddressData

var _ data.AccAddressData = (*accAddressDataI)(nil)

func (a accAddressDataI) String() string {
	return a.Impl.(data.AccAddressData).String()
}

func (a accAddressDataI) GetID() ids.DataID {
	return a.Impl.(data.AccAddressData).GetID()
}

func (a accAddressDataI) Bytes() []byte {
	return a.Impl.(data.AccAddressData).Bytes()
}

func (a accAddressDataI) GetType() ids.StringID {
	return a.Impl.(data.AccAddressData).GetType()
}

func (a accAddressDataI) ZeroValue() data.Data {
	return a.Impl.(data.AccAddressData).ZeroValue()
}

func (a accAddressDataI) GenerateHashID() ids.HashID {
	return a.Impl.(data.AccAddressData).GenerateHashID()
}

func (a accAddressDataI) Compare(listable traits.Listable) int {
	return a.Impl.(data.AccAddressData).Compare(listable)
}

func (a accAddressDataI) Get() sdkTypes.AccAddress {
	return a.Impl.(data.AccAddressData).Get()
}
