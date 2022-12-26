// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/lists"
	"sort"
)

var _ lists.AnyDataList = (*AnyDataList)(nil)

func (dataList *AnyDataList) Search(data data.AnyData) (int, bool) {
	index := sort.Search(
		len(dataList.DataList),
		func(i int) bool {
			return dataList.DataList[i].Compare(data) >= 0
		},
	)

	if index < len(dataList.DataList) && dataList.DataList[index].Compare(data) == 0 {
		return index, true
	}

	return index, false
}
func (dataList *AnyDataList) Add(data ...data.AnyData) lists.AnyDataList {
	updatedList := dataList
	for _, listable := range data {
		if index, found := updatedList.Search(listable); !found {
			updatedList.DataList = append(updatedList.DataList, listable.(*AnyData))
			copy(updatedList.DataList[index+1:], updatedList.DataList[index:])
			updatedList.DataList[index] = listable.(*AnyData)
		}
	}
	return updatedList
}

func (dataList *AnyDataList) GetList() []data.AnyData {
	DataList := make([]data.AnyData, len(dataList.DataList))

	for i, listable := range dataList.DataList {
		if listable != nil {
			DataList[i] = listable
			DataList[i] = listable
		}
	}
	return DataList
}

func (dataList *AnyDataList) Remove(data ...data.AnyData) lists.AnyDataList {
	updatedList := dataList

	for _, listable := range data {
		if index, found := updatedList.Search(listable); found {
			updatedList.DataList = append(updatedList.DataList[:index], updatedList.DataList[index+1:]...)
		}
	}

	return updatedList
}

func NewDataList(data ...data.Data) lists.AnyDataList {
	var dataList []*AnyData
	for _, dataVal := range data {
		dataList = append(dataList, dataVal.ToAnyData().(*AnyData))
	}
	return &AnyDataList{DataList: dataList}
}
