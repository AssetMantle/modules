package base

import (
	dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"
	"github.com/AssetMantle/modules/schema/data"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type accAddressData dataSchema.AccAddressData

var _ data.AccAddressData = &accAddressData{}

func (a accAddressData) GetID() ids.DataID {
	//TODO implement me
	panic("implement me")
}

func (a accAddressData) String() string {
	//TODO implement me
	panic("implement me")
}

func (a accAddressData) Bytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (a accAddressData) GetType() ids.StringID {
	//TODO implement me
	panic("implement me")
}

func (a accAddressData) ZeroValue() Data {
	//TODO implement me
	panic("implement me")
}

func (a accAddressData) GenerateHashID() ids.HashID {
	//TODO implement me
	panic("implement me")
}

func (a accAddressData) Compare(listable traits.Listable) int {
	//TODO implement me
	panic("implement me")
}

func (a accAddressData) Get() sdkTypes.AccAddress {
	//TODO implement me
	panic("implement me")
}
