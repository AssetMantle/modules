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

var _ types.Data = (*StringData)(nil)

func (stringData StringData) Compare(data types.Data) int {
	compareStringData, Error := stringDataFromInterface(data)
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
func stringDataFromInterface(data types.Data) (StringData, error) {
	switch value := data.(type) {
	case *StringData:
		return *value, nil
	default:
		return StringData{}, errors.MetaDataError
	}
}

func NewStringData(value string) types.Data {
	return &StringData{
		Value: value,
	}
}

func ReadStringData(stringData string) (types.Data, error) {
	return NewStringData(stringData), nil
}
