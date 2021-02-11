/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type heightData struct {
	Value types.Height `json:"value"`
}

func (heightData heightData) MarshalAmino() (string, error) {
	return heightData.String(), nil
}

func (heightData *heightData) UnmarshalAmino(text string) (err error) {
	height, Error := strconv.ParseInt(text, 10, 64)
	heightData.Value = NewHeight(height)
	return Error
}

func (heightData heightData) MarshalJSON() ([]byte, error) {
	return json.Marshal(heightData)
}

func (heightData *heightData) UnmarshalJSON(bz []byte) error {
	return json.Unmarshal(bz, &heightData)
}

var _ types.Data = (*heightData)(nil)

func (heightData heightData) String() string {
	return strconv.FormatInt(heightData.Value.Get(), 10)
}
func (heightData heightData) GetTypeID() types.ID {
	return NewID("H")
}
func (heightData heightData) ZeroValue() types.Data {
	return NewHeightData(NewHeight(0))
}
func (heightData heightData) GenerateHashID() types.ID {
	if heightData.Equal(heightData.ZeroValue()) {
		return NewID("")
	}

	return NewID(meta.Hash(strconv.FormatInt(heightData.Value.Get(), 10)))
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
func (heightData heightData) Equal(data types.Data) bool {
	compareHeightData, Error := heightDataFromInterface(data)
	if Error != nil {
		return false
	}

	return heightData.Value.Get() == compareHeightData.Value.Get()
}
func heightDataFromInterface(data types.Data) (*heightData, error) {
	switch value := data.(type) {
	case *heightData:
		return value, nil
	default:
		return &heightData{}, errors.MetaDataError
	}
}

func NewHeightData(value types.Height) types.Data {
	return &heightData{
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
