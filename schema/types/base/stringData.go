// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/lists"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type stringData struct {
	Value string `json:"value"`
}

var _ types.Data = (*stringData)(nil)

func (stringData stringData) GetID() types.ID {
	return dataID{
		TypeID: stringData.GetTypeID(),
		HashID: stringData.GenerateHashID(),
	}
}
func (stringData stringData) Compare(data types.Data) int {
	compareStringData, err := stringDataFromInterface(data)
	if err != nil {
		panic(err)
	}

	return strings.Compare(stringData.Value, compareStringData.Value)
}
func (stringData stringData) String() string {
	return stringData.Value
}
func (stringData stringData) GetTypeID() types.ID {
	return stringDataID
}
func (stringData stringData) ZeroValue() types.Data {
	return NewStringData("")
}
func (stringData stringData) GenerateHashID() types.ID {
	return NewID(meta.Hash(stringData.Value))
}
func (stringData stringData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.EntityNotFound
}
func (stringData stringData) AsDataList() (lists.DataList, error) {
	zeroValue, _ := listData{}.ZeroValue().AsDataList()
	return zeroValue, errors.IncorrectFormat
}
func (stringData stringData) AsString() (string, error) {
	return stringData.Value, nil
}
func (stringData stringData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := decData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (stringData stringData) AsHeight() (types.Height, error) {
	zeroValue, _ := heightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (stringData stringData) AsID() (types.ID, error) {
	zeroValue, _ := idData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (stringData stringData) Get() interface{} {
	return stringData.Value
}
func stringDataFromInterface(data types.Data) (stringData, error) {
	switch value := data.(type) {
	case stringData:
		return value, nil
	default:
		return stringData{}, errors.MetaDataError
	}
}

func NewStringData(value string) types.Data {
	return stringData{
		Value: value,
	}
}

func ReadStringData(stringData string) (types.Data, error) {
	return NewStringData(stringData), nil
}
