/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

var _, _ types.Data = (*Data_StringData)(nil), (*StringData)(nil)

func (stringData Data_StringData) Compare(data types.Data) int {
	compareStringData, Error := stringDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return strings.Compare(stringData.StringData.Value, compareStringData.StringData.Value)
}
func (stringData Data_StringData) String() string {
	return stringData.StringData.Value
}
func (stringData Data_StringData) GetTypeID() types.ID {
	return NewID("S")
}
func (stringData Data_StringData) ZeroValue() types.Data {
	return NewStringData("")
}
func (stringData Data_StringData) GenerateHashID() types.ID {
	return NewID(meta.Hash(stringData.StringData.Value))
}
func (stringData Data_StringData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := Data_AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.EntityNotFound
}
func (stringData Data_StringData) AsListData() (types.ListData, error) {
	zeroValue, _ := Data_ListData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (stringData Data_StringData) AsString() (string, error) {
	return stringData.StringData.Value, nil
}
func (stringData Data_StringData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := Data_DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (stringData Data_StringData) AsHeight() (types.Height, error) {
	zeroValue, _ := Data_HeightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (stringData Data_StringData) AsID() (types.ID, error) {
	zeroValue, _ := Data_IdData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (stringData Data_StringData) Get() interface{} {
	return stringData.StringData.Value
}
func (stringData Data_StringData) Unmarshal(dAtA []byte) error {
	return stringData.StringData.Unmarshal(dAtA)
}
func (stringData *Data_StringData) Reset() { *stringData = Data_StringData{} }
func (*Data_StringData) ProtoMessage()     {}

func stringDataFromInterface(data types.Data) (Data_StringData, error) {
	switch value := data.(type) {
	case *Data_StringData:
		return *value, nil
	default:
		return Data_StringData{}, errors.MetaDataError
	}
}

func NewStringData(value string) types.Data {
	return &Data_StringData{
		StringData: &StringData{
			Value: value,
		},
	}
}

func ReadStringData(stringData string) (types.Data, error) {
	return NewStringData(stringData), nil
}

func (stringData StringData) Compare(data types.Data) int {
	compareStringData, Error := dummyStringDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return strings.Compare(stringData.Value, compareStringData.Value)
}
func (stringData StringData) String() string {
	return stringData.Value
}
func (stringData StringData) GetTypeID() types.ID {
	return NewID("S")
}
func (stringData StringData) ZeroValue() types.Data {
	return NewStringData("")
}
func (stringData StringData) GenerateHashID() types.ID {
	return NewID(meta.Hash(stringData.Value))
}
func (stringData StringData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.EntityNotFound
}
func (stringData StringData) AsListData() (types.ListData, error) {
	zeroValue, _ := ListData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (stringData StringData) AsString() (string, error) {
	return stringData.Value, nil
}
func (stringData StringData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (stringData StringData) AsHeight() (types.Height, error) {
	zeroValue, _ := HeightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (stringData StringData) AsID() (types.ID, error) {
	zeroValue, _ := IDData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (stringData StringData) Get() interface{} {
	return stringData.Value
}
func dummyStringDataFromInterface(data types.Data) (StringData, error) {
	switch value := data.(type) {
	case *StringData:
		return *value, nil
	default:
		return StringData{}, errors.MetaDataError
	}
}
