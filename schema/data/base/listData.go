// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

type listData struct {
	Value lists.DataList `json:"value"`
}

var _ types.Data = (*listData)(nil)

func (listData listData) GetID() types.ID {
	return base.NewDataID(listData)
}
func (listData listData) Compare(data types.Data) int {
	// TODO write test and see if correct
	compareListData, Error := listDataFromInterface(data)
	if Error != nil {
		panic(Error)
	}

	return strings.Compare(listData.GenerateHashID().String(), compareListData.GenerateHashID().String())
}
func (listData listData) String() string {
	dataStringList := make([]string, listData.Value.Size())

	for i, data := range listData.Value.GetList() {
		dataStringList[i] = data.String()
	}

	return strings.Join(dataStringList, constants.ListDataStringSeparator)
}
func (listData listData) GetTypeID() types.ID {
	return ListDataID
}
func (listData listData) ZeroValue() types.Data {
	return NewListData([]types.Data{}...)
}
func (listData listData) GenerateHashID() types.ID {
	if listData.Value.Size() == 0 {
		return base.NewID("")
	}

	hashList := make([]string, listData.Value.Size())

	for i, data := range listData.Value.GetList() {
		hashList[i] = data.GenerateHashID().String()
	}

	hashString := strings.Join(hashList, constants.ListHashStringSeparator)

	return base.NewID(hashString)
}
func (listData listData) AsAccAddress() (sdkTypes.AccAddress, error) {
	zeroValue, _ := accAddressData{}.ZeroValue().AsAccAddress()
	return zeroValue, errors.IncorrectFormat
}
func (listData listData) AsDataList() (lists.DataList, error) {
	return listData.Value, nil
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
func listDataFromInterface(data types.Data) (listData, error) {
	switch value := data.(type) {
	case listData:
		return value, nil
	default:
		return listData{}, errors.MetaDataError
	}
}

func NewListData(value ...types.Data) types.Data {
	// TODO Implement
	return nil
}

func ReadAccAddressListData(dataString string) (types.Data, error) {
	// TODO revise
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
