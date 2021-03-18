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

type stringData struct {
	Value string `json:"value"`
}

var _ types.Data = (*stringData)(nil)

func (stringData stringData) String() string {
	return stringData.Value
}
func (stringData stringData) GetTypeID() types.ID {
	return NewID("S")
}
func (stringData stringData) ZeroValue() types.Data {
	return NewStringData("")
}
func (stringData stringData) GenerateHashID() types.ID {
	return NewID(meta.Hash(stringData.Value))
}
func (stringData stringData) AsAccAddressData() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddressData()
	return zeroValue, errors.EntityNotFound
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
func (stringData stringData) Equal(data types.Data) bool {
	compareStringData, Error := stringDataFromInterface(data)
	if Error != nil {
		return false
	}

	return stringData.Value == compareStringData.Value
}
func stringDataFromInterface(data types.Data) (*stringData, error) {
	switch value := data.(type) {
	case *stringData:
		return value, nil
	default:
		return &stringData{}, errors.MetaDataError
	}
}

func NewStringData(value string) types.Data {
	return &stringData{
		Value: value,
	}
}

func ReadStringData(stringData string) (types.Data, error) {
	return NewStringData(stringData), nil
}
