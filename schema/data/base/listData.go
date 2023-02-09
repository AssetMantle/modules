// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"fmt"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
	"sort"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.ListData = (*ListData)(nil)

func (listData *ListData) Get() []data.AnyData {
	anyDataList := make([]data.AnyData, len(listData.DataList))
	for i, anyData := range listData.DataList {
		anyDataList[i] = anyData
	}
	return anyDataList
}
func (listData *ListData) GetWeight() int {
	return dataConstants.ListDataWidth
}
func (listData *ListData) AsString() string {
	dataStrings := make([]string, len(listData.DataList))

	for i, datum := range listData.DataList {
		dataStrings[i] = datum.AsString()
	}

	return stringUtilities.JoinListStrings(dataStrings...)
}
func (listData *ListData) Search(data data.Data) (int, bool) {
	index := sort.Search(
		len(listData.DataList),
		func(i int) bool {
			return listData.DataList[i].Compare(data) >= 0
		},
	)

	if index < len(listData.DataList) && listData.DataList[index].Compare(data) == 0 {
		return index, true
	}

	return index, false
}
func (listData *ListData) Add(data ...data.Data) data.ListData {
	updatedList := listData
	for _, listable := range data {
		if index, found := updatedList.Search(listable); !found {
			updatedList.DataList = append(updatedList.DataList, listable.ToAnyData().(*AnyData))
			copy(updatedList.DataList[index+1:], updatedList.DataList[index:])
			updatedList.DataList[index] = listable.ToAnyData().(*AnyData)
		}
	}
	return updatedList
}
func (listData *ListData) Remove(data ...data.Data) data.ListData {
	updatedList := listData

	for _, listable := range data {
		if index, found := updatedList.Search(listable); found {
			updatedList.DataList = append(updatedList.DataList[:index], updatedList.DataList[index+1:]...)
		}
	}

	return updatedList
}
func (listData *ListData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(listData)
}
func (listData *ListData) Compare(listable traits.Listable) int {
	compareListData, err := listDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	// TODO check for optimization
	return bytes.Compare(listData.Bytes(), compareListData.Bytes())
}
func (listData *ListData) Bytes() []byte {
	bytesList := make([][]byte, len(listData.DataList))

	for i, datum := range listData.DataList {
		if datum != nil {
			bytesList[i] = datum.Bytes()
		}
	}
	// TODO see if separator required
	return bytes.Join(bytesList, nil)
}
func (listData *ListData) GetType() ids.StringID {
	return dataConstants.ListDataID
}
func (listData *ListData) ZeroValue() data.Data {
	return NewListData([]data.Data{}...)
}
func (listData *ListData) GenerateHashID() ids.HashID {
	if listData.Compare(listData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(listData.Bytes())
}
func (listData *ListData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_ListData{
			ListData: listData,
		},
	}
}
func listDataFromInterface(listable traits.Listable) (*ListData, error) {
	switch value := listable.(type) {
	case *ListData:
		return value, nil
	default:
		x := value
		fmt.Println(x)
		return &ListData{}, errorConstants.MetaDataError
	}
}

func ListDataPrototype() data.ListData {
	return (&ListData{}).ZeroValue().(data.ListData)
}

// NewListData
// * onus of ensuring all Data are of the same type is on DataList
func NewListData(data ...data.Data) data.ListData {
	return (&ListData{}).Add(data...)
}
