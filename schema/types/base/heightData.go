/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

var _, _ types.Data = (*Data_HeightData)(nil), (*HeightData)(nil)

func (heightData Data_HeightData) Compare(data types.Data) int {
	compareHeightData, Error := heightDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return heightData.HeightData.Value.Compare(&compareHeightData.HeightData.Value)
}
func (heightData Data_HeightData) String() string {
	return strconv.FormatInt(heightData.HeightData.Value.Get(), 10)
}
func (heightData Data_HeightData) GetTypeID() types.ID {
	return NewID("H")
}
func (heightData Data_HeightData) ZeroValue() types.Data {
	return NewHeightData(NewHeight(0))
}
func (heightData Data_HeightData) GenerateHashID() types.ID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(strconv.FormatInt(heightData.HeightData.Value.Get(), 10)))
}
func (heightData Data_HeightData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := Data_AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (heightData Data_HeightData) AsListData() (types.ListData, error) {
	zeroValue, _ := Data_ListData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (heightData Data_HeightData) AsString() (string, error) {
	zeroValue, _ := Data_StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (heightData Data_HeightData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := Data_DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (heightData Data_HeightData) AsHeight() (types.Height, error) {
	return &heightData.HeightData.Value, nil
}
func (heightData Data_HeightData) AsID() (types.ID, error) {
	zeroValue, _ := Data_IdData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (heightData Data_HeightData) Get() interface{} {
	return heightData.HeightData.Value
}
func (heightData Data_HeightData) Unmarshal(dAtA []byte) error {
	return heightData.HeightData.Unmarshal(dAtA)
}
func (heightData *Data_HeightData) Reset() { *heightData = Data_HeightData{} }
func (*Data_HeightData) ProtoMessage()     {}
func heightDataFromInterface(data types.Data) (Data_HeightData, error) {
	switch value := data.(type) {
	case *Data_HeightData:
		return *value, nil
	default:
		return Data_HeightData{}, errors.MetaDataError
	}
}

func NewHeightData(value types.Height) *Data_HeightData {
	height := *NewHeight(value.Get())
	return &Data_HeightData{
		HeightData: &HeightData{
			Value: height,
		},
	}
}

func ReadHeightData(dataString string) (types.Data, error) {
	if dataString == "" {
		return Data_HeightData{}.ZeroValue(), nil
	}

	height, Error := strconv.ParseInt(dataString, 10, 64)
	if Error != nil {
		return nil, Error
	}

	return NewHeightData(NewHeight(height)), nil
}

func (heightData HeightData) Compare(data types.Data) int {
	compareHeightData, Error := dummyHeightDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return heightData.Value.Compare(&compareHeightData.Value)
}
func (heightData HeightData) String() string {
	return strconv.FormatInt(heightData.Value.Get(), 10)
}
func (heightData HeightData) GetTypeID() types.ID {
	return NewID("H")
}
func (heightData HeightData) ZeroValue() types.Data {
	return NewHeightData(NewHeight(0))
}
func (heightData HeightData) GenerateHashID() types.ID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(strconv.FormatInt(heightData.Value.Get(), 10)))
}
func (heightData HeightData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (heightData HeightData) AsListData() (types.ListData, error) {
	zeroValue, _ := ListData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (heightData HeightData) AsString() (string, error) {
	zeroValue, _ := StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (heightData HeightData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (heightData HeightData) AsHeight() (types.Height, error) {
	return &heightData.Value, nil
}
func (heightData HeightData) AsID() (types.ID, error) {
	zeroValue, _ := IDData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (heightData HeightData) Get() interface{} {
	return heightData.Value
}
func dummyHeightDataFromInterface(data types.Data) (HeightData, error) {
	switch value := data.(type) {
	case *HeightData:
		return *value, nil
	default:
		return HeightData{}, errors.MetaDataError
	}
}
