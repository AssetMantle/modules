/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
	"strconv"
)

type heightData struct {
	Value types.Height `json:"value"`
}

var _ types.Data = (*heightData)(nil)

func (heightData heightData) GenerateHash() string {
	return meta.Hash(strconv.FormatInt(heightData.Value.Get(), 10))
}

func (heightData heightData) AsString() (string, error) {
	return "", constants.EntityNotFound
}

func (heightData heightData) AsDec() (sdkTypes.Dec, error) {
	return sdkTypes.Dec{}, constants.EntityNotFound
}

func (heightData heightData) AsHeight() (types.Height, error) {
	return heightData.Value, nil
}

func (heightData heightData) Get() interface{} {
	return heightData.Value
}

func (heightData heightData) AsID() (types.ID, error) {
	return id{}, constants.EntityNotFound
}

func NewHeightData(value types.Height) types.Data {
	return heightData{
		Value: value,
	}
}

func ReadHeightData(dataString string) types.Data {
	if height, Error := strconv.ParseInt(dataString, 10, 64); Error != nil {
		return NewHeightData(NewHeight(height))
	}
	return nil
}
