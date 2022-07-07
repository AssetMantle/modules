// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/data/utlities"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/traits"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

type listData struct {
	Value lists.DataList `json:"value"`
}

var _ data.ListData = (*listData)(nil)

func (listData listData) GetID() ids.DataID {
	return baseIDs.NewDataID(listData)
}
func (listData listData) Compare(listable traits.Listable) int {
	compareListData, err := listDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return strings.Compare(listData.GenerateHash().String(), compareListData.GenerateHash().String())
}
func (listData listData) String() string {
	dataStrings := make([]string, listData.Value.Size())

	for i, datum := range listData.Value.GetList() {
		dataStrings[i] = datum.String()
	}

	return stringUtilities.JoinListStrings(dataStrings...)
}
func (listData listData) GetType() ids.ID {
	return dataConstants.ListDataID
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
		hashList[i] = datum.GenerateHash().String()
	}

	return baseIDs.NewID(stringUtilities.Hash(hashList...))
}
func (listData listData) Get() lists.DataList {
	return listData.Value
}

func listDataFromInterface(listable traits.Listable) (listData, error) {
	switch value := listable.(type) {
	case listData:
		return value, nil
	default:
		return listData{}, errorConstants.MetaDataError
	}
}

// NewListData
// * onus of ensuring all Data are of the same type is on DataList
func NewListData(value ...data.Data) data.Data {
	return listData{Value: baseLists.NewDataList(value...)}
}

func ReadListData(dataString string) (data.Data, error) {
	if dataString == "" {
		return listData{}.ZeroValue(), nil
	}

	dataStringList := stringUtilities.SplitListString(dataString)
	dataList := make([]data.Data, len(dataStringList))

	for i, datumString := range dataStringList {
		data, err := utlities.ReadData(datumString)
		if err != nil {
			return listData{}.ZeroValue(), err
		}

		dataList[i] = data
	}

	return NewListData(dataList...), nil
}
