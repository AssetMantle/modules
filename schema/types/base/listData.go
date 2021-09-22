/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strings"

	"github.com/persistenceOne/persistenceSDK/constants/ids"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type listData struct {
	Value sortedDataList `json:"value"`
}

var _ types.ListData = (*listData)(nil)

func (listData listData) GetID() types.ID {
	return dataID{
		TypeID: listData.GetTypeID(),
		HashID: listData.GenerateHashID(),
	}
}
func (listData listData) Compare(data types.Data) int {
	compareListData, Error := listDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return strings.Compare(listData.GenerateHashID().String(), compareListData.GenerateHashID().String())
}
func (listData listData) String() string {
	dataStringList := make([]string, len(listData.Value))

	for i, data := range listData.Value {
		dataStringList[i] = data.String()
	}

	return strings.Join(dataStringList, constants.ListDataStringSeparator)
}
func (listData listData) GetTypeID() types.ID {
	return ids.ListData
}
func (listData listData) ZeroValue() types.Data {
	return NewListData([]types.Data{}...)
}
func (listData listData) GenerateHashID() types.ID {
	if len(listData.Value) == 0 {
		return NewID("")
	}

	hashList := make([]string, len(listData.Value))

	for i, data := range listData.Value {
		hashList[i] = data.GenerateHashID().String()
	}

	hashString := strings.Join(hashList, constants.ListHashStringSeparator)

	return NewID(hashString)
}
func (listData listData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (listData listData) AsListData() (types.ListData, error) {
	return listData, nil
}
func (listData listData) AsString() (string, error) {
	zeroValue, _ := stringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (listData listData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := decData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (listData listData) AsHeight() (types.Height, error) {
	zeroValue, _ := heightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (listData listData) AsID() (types.ID, error) {
	zeroValue, _ := idData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (listData listData) Get() interface{} {
	return listData.Value
}
func (listData listData) Search(data types.Data) int {
	return listData.Value.Search(data)
}
func (listData listData) GetList() []types.Data {
	return listData.Value
}
func (listData listData) Add(dataList ...types.Data) types.ListData {
	for _, data := range dataList {
		listData.Value = listData.Value.Add(data).(sortedDataList)
	}

	return listData
}
func (listData listData) Remove(dataList ...types.Data) types.ListData {
	for _, data := range dataList {
		listData.Value = listData.Value.Remove(data).(sortedDataList)
	}

	return listData
}
func listDataFromInterface(data types.Data) (listData, error) {
	switch value := data.(type) {
	case listData:
		return value, nil
	default:
		return listData{}, errors.MetaDataError
	}
}

func NewListData(value ...types.Data) types.Data {
	return listData{}.Add(value...)
}

func ReadAccAddressListData(dataString string) (types.Data, error) {
	if dataString == "" {
		return listData{}.ZeroValue(), nil
	}

	dataStringList := strings.Split(dataString, constants.ListDataStringSeparator)
	dataList := make([]types.Data, len(dataStringList))

	for i, accAddressString := range dataStringList {
		accAddress, Error := sdkTypes.AccAddressFromBech32(accAddressString)
		if Error != nil {
			return listData{}.ZeroValue(), Error
		}

		dataList[i] = NewAccAddressData(accAddress)
	}

	return NewListData(dataList...), nil
}
