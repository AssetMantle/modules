/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

const StringType = "S"

type stringData struct {
	Value string `json:"value"`
}

var _ types.Data = (*stringData)(nil)

func (stringData stringData) GenerateHash() string {
	return meta.Hash(stringData.Value)
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

func NewStringData(value string) types.Data {
	return stringData{
		Value: value,
	}
}

func ReadStringData(stringData string) types.Data {
	return NewStringData(stringData)
}
