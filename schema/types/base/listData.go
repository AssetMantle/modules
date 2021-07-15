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

type listData struct {
	Value []types.Data `json:"value"`
}

var _ types.ListData = (*listData)(nil)

func (listData listData) String() string {
	dataStringList := make([]string, len(listData.Value))

	for i, data := range listData.Value {
		dataStringList[i] = data.String()
	}

	return strings.Join(dataStringList, constants.ListDataStringSeparator)
}
func (listData listData) GetTypeID() types.ID {
	return NewID("LD")
}
func (listData listData) ZeroValue() types.Data {
	return NewListData([]types.Data{}...)
}
func (listData listData) GenerateHashID() types.ID {
	if len(listData.Value) == 0 {
		return NewID("")
	}

	return NewID(meta.Hash(listData.String()))
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
func (listData listData) Equal(data types.Data) bool {
	compareAccAddressListData, Error := listDataFromInterface(data)
	if Error != nil {
		return false
	}

	if len(listData.Value) != len(compareAccAddressListData.Value) {
		return false
	}

	return listData.GenerateHashID().Equals(compareAccAddressListData.GenerateHashID())
}
func (listData listData) Add(dataList ...types.Data) types.ListData {
	for _, data := range dataList {
		accAddressData, Error := accAddressDataFromInterface(data)
		if Error != nil {
			panic(Error)
		}

		listData.Value = listData.Value.Insert(accAddressData.Value).(sortedList)
	}

	return listData
}
func (listData listData) Remove(dataList ...types.Data) types.ListData {
	for _, data := range dataList {
		accAddressData, Error := accAddressDataFromInterface(data)
		if Error != nil {
			panic(Error)
		}

		listData.Value = listData.Value.Delete(accAddressData.Value).(sortedList)
	}

	return listData
}
func (listData listData) IsPresent(data types.Data) bool {
	accAddressData, Error := accAddressDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return listData.Value.Search(accAddressData.Value) != listData.Value.Len()
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
	return listData{
		Value: value,
	}
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
