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

var _ types.ListData = (*ListData)(nil)

// TODO: find a better impl
func (listData ListData) Compare(data types.Data) int {
	compareListData, Error := listDataFromData(data)
	if Error != nil {
		panic(Error)
	}

	var listDataString []string
	for _, data := range listData.Value {
		listDataString = append(listDataString, data.String())
	}

	var comparisonDataString []string
	for _, data := range compareListData.Value {
		comparisonDataString = append(comparisonDataString, data.String())
	}

	return strings.Compare(strings.Join(listDataString, constants.ListDataStringSeparator), strings.Join(comparisonDataString, constants.ListDataStringSeparator))
}
func (listData ListData) String() string {
	dataStringList := make([]string, len(listData.Value))

	for i, data := range listData.Value {
		dataStringList[i] = data.String()
	}

	return strings.Join(dataStringList, constants.ListDataStringSeparator)
}
func (listData ListData) GetTypeID() types.ID {
	return NewID("LD")
}
func (listData ListData) ZeroValue() types.Data {
	return NewListData([]types.Data{}...)
}
func (listData ListData) GenerateHashID() types.ID {
	if len(listData.Value) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(listData.String()))
}
func (listData ListData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := AccAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (listData ListData) AsListData() (types.ListData, error) {
	return &listData, nil
}
func (listData ListData) AsString() (string, error) {
	zeroValue, _ := StringData{}.ZeroValue().AsString()
	return zeroValue, errors.IncorrectFormat
}
func (listData ListData) AsDec() (sdkTypes.Dec, error) {
	zeroValue, _ := DecData{}.ZeroValue().AsDec()
	return zeroValue, errors.IncorrectFormat
}
func (listData ListData) AsHeight() (types.Height, error) {
	zeroValue, _ := HeightData{}.ZeroValue().AsHeight()
	return zeroValue, errors.IncorrectFormat
}
func (listData ListData) AsID() (types.ID, error) {
	zeroValue, _ := IDData{}.ZeroValue().AsID()
	return zeroValue, errors.IncorrectFormat
}
func (listData ListData) Get() interface{} {
	return listData.Value
}
func (listData ListData) Search(data types.Data) int {
	return sortedDataList(listData.Value).Search(data)
}
func (listData ListData) GetList() []types.Data {
	return listData.Value
}
func (listData ListData) Add(dataList ...types.Data) types.ListData {
	for _, data := range dataList {
		listData.Value = sortedDataList(listData.Value).Add(data).GetList()
	}

	return &listData
}
func (listData ListData) Remove(dataList ...types.Data) types.ListData {
	for _, data := range dataList {
		listData.Value = sortedDataList(listData.Value).Remove(data).GetList()
	}

	return &listData
}
func listDataFromData(data types.Data) (ListData, error) {
	switch value := data.(type) {
	case *ListData:
		return *value, nil
	default:
		return ListData{}, errors.MetaDataError
	}
}

func NewListData(value ...types.Data) types.Data {
	return ListData{}.Add(value...)
}

func ReadAccAddressListData(dataString string) (types.Data, error) {
	if dataString == "" {
		return ListData{}.ZeroValue(), nil
	}

	dataStringList := strings.Split(dataString, constants.ListDataStringSeparator)
	dataList := make([]types.Data, len(dataStringList))

	for i, accAddressString := range dataStringList {
		accAddress, Error := sdkTypes.AccAddressFromBech32(accAddressString)
		if Error != nil {
			return ListData{}.ZeroValue(), Error
		}

		dataList[i] = NewAccAddressData(accAddress)
	}

	return NewListData(dataList...), nil
}
