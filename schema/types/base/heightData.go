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
	"strconv"
)

type heightData struct {
	Value types.Height `json:"value"`
}

var _ types.Data = (*heightData)(nil)

func (HeightData heightData) String() string {
	return strconv.FormatInt(HeightData.Value.Get(), 10)
}
func (HeightData heightData) GetTypeID() types.ID {
	return NewID("H")
}
func (HeightData heightData) ZeroValue() types.Data {
	return NewHeightData(NewHeight(0))
}
func (HeightData heightData) GenerateHashID() types.ID {
	if HeightData.Equal(HeightData.ZeroValue()) {
		return NewID("")
	}
	return NewID(meta.Hash(strconv.FormatInt(HeightData.Value.Get(), 10)))
}
func (HeightData heightData) AsString() (string, error) {
	return "", errors.EntityNotFound
}
func (HeightData heightData) AsDec() (sdkTypes.Dec, error) {
	return sdkTypes.Dec{}, errors.EntityNotFound
}
func (HeightData heightData) AsHeight() (types.Height, error) {
	return HeightData.Value, nil
}
func (HeightData heightData) Get() interface{} {
	return HeightData.Value
}
func (HeightData heightData) AsID() (types.ID, error) {
	return id{}, errors.EntityNotFound
}
func (HeightData heightData) Equal(data types.Data) bool {
	switch value := data.(type) {
	case heightData:
		return value.Value.Get() == HeightData.Value.Get()
	default:
		return false
	}
}

func NewHeightData(value types.Height) types.Data {
	return heightData{
		Value: value,
	}
}

func ReadHeightData(dataString string) (types.Data, error) {
	if dataString == "" {
		return heightData{}.ZeroValue(), nil
	}
	height, Error := strconv.ParseInt(dataString, 10, 64)
	if Error != nil {
		return nil, Error
	}
	return NewHeightData(NewHeight(height)), nil
}
