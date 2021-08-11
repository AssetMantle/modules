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

var _ types.Data = (*HeightData)(nil)

func (heightData HeightData) Compare(data types.Data) int {
	compareHeightData, Error := heightDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return heightData.Value.Compare(compareHeightData.Value)
}
func (heightData HeightData) String() string {
	return strconv.FormatInt(heightData.Value.Get(), 10)
}
func (heightData HeightData) GetTypeID() types.ID {
	return NewID("H")
}
func (heightData HeightData) ZeroValue() types.Data {
	return NewHeightData(NewHeight(0))
}
func (heightData HeightData) GenerateHashID() types.ID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(strconv.FormatInt(heightData.Value.Get(), 10)))
}
func (heightData HeightData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (heightData HeightData) AsListData() (types.ListData, error) {
	zeroValue, _ := ListData{}.ZeroValue().AsListData()
	return zeroValue, errors.IncorrectFormat
}
func (heightData HeightData) AsString() (string, error) {
	zeroValue, _ := StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (heightData HeightData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (heightData HeightData) AsHeight() (types.Height, error) {
	return heightData.Value, nil
}
func (heightData HeightData) AsID() (types.ID, error) {
	zeroValue, _ := IDData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (heightData HeightData) Get() interface{} {
	return heightData.Value
}
func heightDataFromInterface(data types.Data) (HeightData, error) {
	switch value := data.(type) {
	case HeightData:
		return value, nil
	default:
		return HeightData{}, errors.MetaDataError
	}
}

func NewHeightData(value types.Height) types.Data {
	return HeightData{
		Value: value,
	}
}

func ReadHeightData(dataString string) (types.Data, error) {
	if dataString == "" {
		return HeightData{}.ZeroValue(), nil
	}

	height, Error := strconv.ParseInt(dataString, 10, 64)
	if Error != nil {
		return nil, Error
	}

	return NewHeightData(NewHeight(height)), nil
}
