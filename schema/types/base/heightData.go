/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type heightData struct {
	Value types.Height `json:"value"`
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
func (heightData heightData) AsAccAddressData() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddressData()
	return zeroValue, errors.IncorrectFormat
}
func (heightData heightData) AsString() (string, error) {
	zeroValue, _ := stringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (heightData heightData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := decData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (heightData heightData) AsHeight() (types.Height, error) {
	return heightData.Value, nil
}
func (heightData heightData) AsID() (types.ID, error) {
	zeroValue, _ := idData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (heightData heightData) Get() interface{} {
	return heightData.Value
}
func (heightData heightData) Equal(data types.Data) bool {
	compareHeightData, Error := heightDataFromInterface(data)
	if Error != nil {
		return false
	}

	return heightData.Value.Get() == compareHeightData.Value.Get()
}
func heightDataFromInterface(data types.Data) (heightData, error) {
	switch value := data.(type) {
	case heightData:
		return value, nil
	default:
		return heightData{}, errors.MetaDataError
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
