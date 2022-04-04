// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/lists"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type heightData struct {
	Value types.Height `json:"value"`
}

var _ types.Data = (*heightData)(nil)

func (heightData heightData) GetID() types.ID {
	return dataID{
		TypeID: heightData.GetTypeID(),
		HashID: heightData.GenerateHashID(),
	}
}
func (heightData heightData) Compare(data types.Data) int {
	compareHeightData, err := heightDataFromInterface(data)
	if err != nil {
		panic(err)
	}

	return heightData.Value.Compare(compareHeightData.Value)
}
func (heightData heightData) String() string {
	return strconv.FormatInt(heightData.Value.Get(), 10)
}
func (heightData heightData) GetTypeID() types.ID {
	return heightDataID
}
func (heightData heightData) ZeroValue() types.Data {
	return NewHeightData(NewHeight(0))
}
func (heightData heightData) GenerateHashID() types.ID {
	if heightData.Compare(heightData.ZeroValue()) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(strconv.FormatInt(heightData.Value.Get(), 10)))
}
func (heightData heightData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (heightData heightData) AsDataList() (lists.DataList, error) {
	zeroValue, _ := listData{}.ZeroValue().AsDataList()
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

	height, err := strconv.ParseInt(dataString, 10, 64)
	if err != nil {
		return nil, err
	}

	return NewHeightData(NewHeight(height)), nil
}
