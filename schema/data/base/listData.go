// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type listData struct {
	Value lists.DataList `json:"value"`
}

var _ data.ListData = (*listData)(nil)

func (listData listData) GetID() types.ID {
	return baseTypes.NewDataID(listData)
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

	for i, datum := range listData.Value.GetList() {
		dataStringList[i] = datum.String()
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
		return baseTypes.NewID("")
	}

	hashList := make([]string, listData.Value.Size())

	for i, datum := range listData.Value.GetList() {
		hashList[i] = datum.GenerateHashID().String()
	}

	hashString := strings.Join(hashList, constants.ListHashStringSeparator)

	return baseTypes.NewID(hashString)
}
func (listData listData) Get() lists.DataList {
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

// NewListData
// * onus of ensuring all Data are of the same type is on DataList
func NewListData(value ...types.Data) types.Data {
	return listData{Value: baseLists.NewDataList(value...)}
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
