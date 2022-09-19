// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/traits"
)

type listData struct {
	Value lists.DataList `json:"value"`
}

var _ data.ListData = (*listData)(nil)

func (listData listData) GetID() ids.DataID {
	return baseIDs.NewDataID(listData)
}
func (listData listData) Compare(listable traits.Listable) int {
	// TODO write test
	compareListData, Error := listDataFromInterface(listable)
	if Error != nil {
		panic(Error)
	}

	return strings.Compare(listData.GenerateHash().String(), compareListData.GenerateHash().String())
}
func (listData listData) String() string {
	dataStringList := make([]string, listData.Value.Size())

	for i, datum := range listData.Value.GetList() {
		dataStringList[i] = datum.String()
	}

	return strings.Join(dataStringList, constants.ListDataStringSeparator)
}
func (listData listData) GetType() ids.ID {
	return idsConstants.ListDataID
}
func (listData listData) ZeroValue() data.Data {
	return NewListData([]data.Data{}...)
}
func (listData listData) GenerateHash() ids.ID {
	if listData.Value.Size() == 0 {
		return baseIDs.NewID("")
	}

	hashList := make([]string, listData.Value.Size())

	for i, datum := range listData.Value.GetList() {
		if datum != nil {
			hashList[i] = datum.GenerateHash().String()
		}
	}

	hashString := strings.Join(hashList, constants.ListHashStringSeparator)

	return baseIDs.NewID(hashString)
}
func (listData listData) Get() []data.Data {
	return listData.Value.GetList()
}

func listDataFromInterface(listable traits.Listable) (listData, error) {
	switch value := listable.(type) {
	case listData:
		return value, nil
	default:
		return listData{}, errors.MetaDataError
	}
}

// NewListData
// * onus of ensuring all Data are of the same type is on DataList
func NewListData(value ...data.Data) data.Data {
	return listData{Value: baseLists.NewDataList(value...)}
}

func ReadListData(dataString string) (data.Data, error) {
	// TODO revise
	if dataString == "" {
		return listData{}.ZeroValue(), nil
	}

	dataStringList := strings.Split(dataString, constants.ListDataStringSeparator)
	dataList := make([]data.Data, len(dataStringList))

	for i, accAddressString := range dataStringList {
		accAddress, Error := sdkTypes.AccAddressFromBech32(accAddressString)
		if Error != nil {
			return listData{}.ZeroValue(), Error
		}

		dataList[i] = NewAccAddressData(accAddress)
	}

	return NewListData(dataList...), nil
}
