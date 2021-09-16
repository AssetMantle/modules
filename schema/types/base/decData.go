/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

func decDataFromInterface(data types.Data) (Data_DecData, error) {
	switch value := data.(type) {
	case *Data_DecData:
		return *value, nil
	default:
		return Data_DecData{}, errors.MetaDataError
	}
}

func NewDecData(value sdkTypes.Dec) *Data_DecData {
	return &Data_DecData{
		DecData: &DecData{
			Value: value,
		},
	}
}

func ReadDecData(dataString string) (types.Data, error) {
	if dataString == "" {
		return Data_DecData{}.ZeroValue(), nil
	}

	dec, Error := sdkTypes.NewDecFromStr(dataString)
	if Error != nil {
		return Data_DecData{}.ZeroValue(), Error
	}

	return NewDecData(dec), nil
}

var _ types.Data = (*Data_DecData)(nil)

func (decData Data_DecData) Compare(data types.Data) int {
	compareDecData, Error := decDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	if decData.DecData.Value.GT(compareDecData.DecData.Value) {
		return 1
	} else if decData.DecData.Value.LT(compareDecData.DecData.Value) {
		return -1
	}

	return 0
}
func (decData Data_DecData) String() string {
	return decData.DecData.Value.String()
}
func (decData Data_DecData) GetTypeID() types.ID {
	return NewID("D")
}
func (decData Data_DecData) ZeroValue() types.Data {
	return NewDecData(sdkTypes.ZeroDec())
}
func (decData Data_DecData) GenerateHashID() types.ID {
	if decData.Compare(decData.ZeroValue()) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(decData.DecData.Value.String()))
}
func (decData Data_DecData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := Data_AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (decData Data_DecData) AsListData() (types.ListData, error) {
	zeroValue, _ := Data_ListData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (decData Data_DecData) AsString() (string, error) {
	zeroValue, _ := Data_StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (decData Data_DecData) AsDec() (sdkTypes.Dec, error) {
	return decData.DecData.Value, nil
}
func (decData Data_DecData) AsHeight() (types.Height, error) {
	zeroValue, _ := Data_HeightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (decData Data_DecData) AsID() (types.ID, error) {
	zeroValue, _ := Data_DecData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (decData Data_DecData) Get() interface{} {
	return decData.DecData.Value
}
func (decData Data_DecData) Unmarshal(dAtA []byte) error {
	return decData.DecData.Unmarshal(dAtA)
}
func (decData *Data_DecData) Reset() { *decData = Data_DecData{} }
func (*Data_DecData) ProtoMessage()  {}
