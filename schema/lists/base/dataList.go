// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/lists"
	"sort"
)

var _ lists.DataList = (*List_DataList)(nil)

func (dataList *List_DataList) GetList() []data.Data {
	DataList := make([]data.Data, len(dataList.DataList.List))

	for i, listable := range dataList.DataList.List {
		if listable != nil {
			DataList[i] = listable
		}
	}

	return DataList
}
func (dataList *List_DataList) Search(data data.Data) (int, bool) {
	index := sort.Search(
		len(dataList.DataList.List),
		func(i int) bool {
			return dataList.DataList.List[i].Compare(data) >= 0
		},
	)

	if index < len(dataList.DataList.List) && dataList.DataList.List[index].Compare(data) == 0 {
		return index, true
	}

	return index, false
}
func (dataList *List_DataList) Add(data ...data.Data) lists.List {
	updatedList := dataList
	for _, listable := range data {
		if index, found := updatedList.Search(listable); !found {
			updatedList.DataList.List = append(updatedList.DataList.List, listable.(*baseData.Data))
			copy(updatedList.DataList.List[index+1:], updatedList.DataList.List[index:])
			updatedList.DataList.List[index] = listable.(*baseData.Data)
		}
	}
	return &List{
		Impl: updatedList,
	}
}

func (dataList *List_DataList) Remove(data ...data.Data) lists.List {
	updatedList := dataList

	for _, listable := range data {
		if index, found := updatedList.Search(listable); found {
			updatedList.DataList.List = append(updatedList.DataList.List[:index], updatedList.DataList.List[index+1:]...)
		}
	}

	return &List{
		Impl: updatedList,
	}
}

func NewDataList(data ...data.Data) lists.List {
	var dataList []*baseData.Data
	for _, dataVal := range data {
		dataList = append(dataList, dataVal.(*baseData.Data))
	}
	sort.Slice(dataList, func(i, j int) bool {
		return dataList[i].Compare(dataList[j]) <= 0
	})
	return &List{
		Impl: &List_DataList{
			DataList: &DataList{
				List: dataList,
			},
		},
	}
}
