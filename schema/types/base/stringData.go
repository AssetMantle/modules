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

func (StringData stringData) String() string {
	return StringData.Value
}
func (StringData stringData) GetTypeID() types.ID {
	return NewID("S")
}
func (StringData stringData) ZeroValue() types.Data {
	return NewStringData("")
}
func (StringData stringData) GenerateHashID() types.ID {
	if StringData.Equal(StringData.ZeroValue()) {
		return NewID("")
	}
	return NewID(meta.Hash(StringData.Value))
}
func (StringData stringData) AsString() (string, error) {
	return StringData.Value, nil
}
func (StringData stringData) AsDec() (sdkTypes.Dec, error) {
	return sdkTypes.Dec{}, errors.EntityNotFound
}
func (StringData stringData) AsHeight() (types.Height, error) {
	return height{}, errors.EntityNotFound
}
func (StringData stringData) AsID() (types.ID, error) {
	return id{}, errors.EntityNotFound
}
func (StringData stringData) Get() interface{} {
	return StringData.Value
}
func (StringData stringData) Equal(data types.Data) bool {
	switch value := data.(type) {
	case stringData:
		return value.Value == StringData.Value
	default:
		return false
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
