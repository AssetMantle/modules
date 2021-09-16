/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

var _ types.ListData = (*Data_ListData)(nil)

// TODO: find a better impl
func (listData Data_ListData) Compare(data types.Data) int {
	compareListData, Error := listDataFromData(data)
	if Error != nil {
		panic(Error)
	}

	var listDataString []string
	for _, data := range listData.ListData.Value {
		listDataString = append(listDataString, data.String())
	}

	var comparisonDataString []string
	for _, data := range compareListData.ListData.Value {
		comparisonDataString = append(comparisonDataString, data.String())
	}

	return strings.Compare(strings.Join(listDataString, constants.ListDataStringSeparator), strings.Join(comparisonDataString, constants.ListDataStringSeparator))
}
func (listData Data_ListData) String() string {
	dataStringList := make([]string, len(listData.ListData.Value))

	for i, data := range listData.ListData.Value {
		dataStringList[i] = data.String()
	}

	return strings.Join(dataStringList, constants.ListDataStringSeparator)
}
func (listData Data_ListData) GetTypeID() types.ID {
	return NewID("LD")
}
func (listData Data_ListData) ZeroValue() types.Data {
	return NewListData([]types.Data{}...)
}
func (listData Data_ListData) GenerateHashID() types.ID {
	if len(listData.ListData.Value) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(listData.String()))
}
func (listData Data_ListData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := Data_AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (listData Data_ListData) AsListData() (types.ListData, error) {
	return &listData, nil
}
func (listData Data_ListData) AsString() (string, error) {
	zeroValue, _ := Data_StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (listData Data_ListData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := Data_DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (listData Data_ListData) AsHeight() (types.Height, error) {
	zeroValue, _ := Data_HeightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (listData Data_ListData) AsID() (types.ID, error) {
	zeroValue, _ := Data_IdData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (listData Data_ListData) Get() interface{} {
	return listData.ListData.Value
}
func (listData Data_ListData) Search(data types.Data) int {
	newList := make([]types.Data, len(listData.ListData.Value))
	for i, element := range listData.ListData.Value {
		newList[i] = element.Data
	}
	return sortedDataList(newList).Search(data)
}
func (listData Data_ListData) GetList() []types.Data {
	newList := make([]types.Data, len(listData.ListData.Value))
	for i, element := range listData.ListData.Value {
		newList[i] = element.Data
	}
	return newList
}
func (listData Data_ListData) Add(dataList ...types.Data) types.ListData {
	newList := make([]types.Data, len(listData.ListData.Value))
	for i, element := range listData.ListData.Value {
		newList[i] = element.Data
	}
	for _, data := range dataList {
		dataList := sortedDataList(newList).Add(data).GetList()
		newDataList := make([]Data, len(dataList))
		for i, element := range dataList {
			newDataList[i] = *NewData(element)
		}
		listData.ListData.Value = newDataList
	}

	return &listData
}
func (listData Data_ListData) Remove(dataList ...types.Data) types.ListData {
	newList := make([]types.Data, len(listData.ListData.Value))
	for i, element := range listData.ListData.Value {
		newList[i] = element.Data
	}
	for _, data := range dataList {
		dataList := sortedDataList(newList).Add(data).GetList()
		newDataList := make([]Data, len(dataList))
		for i, element := range dataList {
			newDataList[i] = *NewData(element)
		}
		listData.ListData.Value = newDataList
	}

	return &listData
}
func (listData Data_ListData) Unmarshal(dAtA []byte) error {
	return listData.ListData.Unmarshal(dAtA)
}
func (listData *Data_ListData) Reset() { *listData = Data_ListData{} }
func (*Data_ListData) ProtoMessage()   {}

func listDataFromData(data types.Data) (Data_ListData, error) {
	switch value := data.(type) {
	case *Data_ListData:
		return *value, nil
	default:
		return Data_ListData{}, errors.MetaDataError
	}
}

func NewListData(value ...types.Data) *Data_ListData {
	newValue := make([]Data, len(value))
	for i, element := range value {
		newValue[i] = *NewData(element)
	}
	return &Data_ListData{
		ListData: ListData{
			Value: newValue,
		},
	}
}

func ReadAccAddressListData(dataString string) (types.Data, error) {
	if dataString == "" {
		return Data_ListData{}.ZeroValue(), nil
	}

	dataStringList := strings.Split(dataString, constants.ListDataStringSeparator)
	dataList := make([]types.Data, len(dataStringList))

	for i, accAddressString := range dataStringList {
		accAddress, Error := sdkTypes.AccAddressFromBech32(accAddressString)
		if Error != nil {
			return Data_ListData{}.ZeroValue(), Error
		}

		dataList[i] = NewAccAddressData(accAddress)
	}

	return NewListData(dataList...), nil
}
