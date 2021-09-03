package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Data = &Data{}

func (m Data) Compare(data types.Data) int {
	return m.Data.Compare(data)
}

func (m Data) GetTypeID() types.ID {
	return m.Data.GetTypeID()
}

func (m Data) ZeroValue() types.Data {
	return m.Data.ZeroValue()
}

func (m Data) GenerateHashID() types.ID {
	return m.Data.GetTypeID()
}

func (m Data) AsAccAddress() (sdkTypes.AccAddress, error) {
	return m.Data.AsAccAddress()
}

func (m Data) AsListData() (types.ListData, error) {
	return m.Data.AsListData()
}

func (m Data) AsString() (string, error) {
	return m.Data.AsString()
}

func (m Data) AsDec() (sdkTypes.Dec, error) {
	return m.Data.AsDec()
}

func (m Data) AsHeight() (types.Height, error) {
	return m.Data.AsHeight()
}

func (m Data) AsID() (types.ID, error) {
	return m.Data.AsID()
}

func (m Data) Get() interface{} {
	return m.Data.Get()
}

func NewData(data types.Data) *Data {
	return &Data{
		Data: data,
	}
}

//func dataFromInterface(data types.Data) (Data, error) {
//	switch value := data.(type) {
//	case *Data:
//		return *value, nil
//	default:
//		return Data{}, errors.MetaDataError
//	}
//}
