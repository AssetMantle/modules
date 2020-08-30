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
	"strconv"
)

const HeightType = "H"

type heightData struct {
	Value types.Height `json:"value"`
}

var _ types.Data = (*heightData)(nil)

func (heightData heightData) GenerateHash() string {
	if heightData.Value.Get() == -1 {
		return ""
	}
	return meta.Hash(strconv.FormatInt(heightData.Value.Get(), 10))
}

func (heightData heightData) AsString() (string, error) {
	return "", errors.EntityNotFound
}

func (heightData heightData) AsDec() (sdkTypes.Dec, error) {
	return sdkTypes.Dec{}, errors.EntityNotFound
}

func (heightData heightData) AsHeight() (types.Height, error) {
	return heightData.Value, nil
}

func (heightData heightData) Get() interface{} {
	return heightData.Value
}

func (heightData heightData) AsID() (types.ID, error) {
	return id{}, errors.EntityNotFound
}

func NewHeightData(value types.Height) types.Data {
	return heightData{
		Value: value,
	}
}

func ReadHeightData(dataString string) (types.Data, error) {
	if dataString == "" {
		return NewHeightData(NewHeight(-1)), nil
	}
	height, Error := strconv.ParseInt(dataString, 10, 64)
	if Error != nil {
		return nil, Error
	}
	return NewHeightData(NewHeight(height)), nil
}
