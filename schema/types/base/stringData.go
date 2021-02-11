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

func (stringData stringData) MarshalAmino() (string, error) {
	return stringData.String(), nil
}

func (stringData *stringData) UnmarshalAmino(text string) (err error) {
	stringData.Value = text
	return nil
}

func (stringData stringData) MarshalJSON() ([]byte, error) {
	//return NewCodec().MarshalJSON(stringData)
	panic("impl me")
}

func (stringData *stringData) UnmarshalJSON(bz []byte) error {
	//return NewCodec().UnmarshalJSON(bz, &stringData)
	panic("impl me")
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
	if stringData.Equal(stringData.ZeroValue()) {
		return NewID("")
	}

	return NewID(meta.Hash(stringData.Value))
}
func (stringData stringData) AsString() (string, error) {
	return stringData.Value, nil
}
func (stringData stringData) AsDec() (sdkTypes.Dec, error) {
	return sdkTypes.Dec{}, errors.EntityNotFound
}
func (stringData stringData) AsHeight() (types.Height, error) {
	return height{}, errors.EntityNotFound
}
func (stringData stringData) AsID() (types.ID, error) {
	return id{}, errors.EntityNotFound
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
