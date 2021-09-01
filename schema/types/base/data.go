package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Data = &Data{}

func (m Data) Compare(data types.Data) int {
	panic("implement me")
}

func (m Data) GetTypeID() types.ID {
	panic("implement me")
}

func (m Data) ZeroValue() types.Data {
	panic("implement me")
}

func (m Data) GenerateHashID() types.ID {
	panic("implement me")
}

func (m Data) AsAccAddress() (sdkTypes.AccAddress, error) {
	panic("implement me")
}

func (m Data) AsListData() (types.ListData, error) {
	panic("implement me")
}

func (m Data) AsString() (string, error) {
	panic("implement me")
}

func (m Data) AsDec() (sdkTypes.Dec, error) {
	panic("implement me")
}

func (m Data) AsHeight() (types.Height, error) {
	panic("implement me")
}

func (m Data) AsID() (types.ID, error) {
	panic("implement me")
}

func (m Data) Get() interface{} {
	panic("implement me")
}

func NewData(data types.Data) *Data {
	return &Data{
		Data: data,
	}
}
