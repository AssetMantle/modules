// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/lists"
	"sort"
)

var _ lists.DataList = (*AnyDataList)(nil)

func (dataList *AnyDataList) Search(data data.Data) (int, bool) {
	index := sort.Search(
		len(dataList.List),
		func(i int) bool {
			return dataList.List[i].Compare(data) >= 0
		},
	)

	if index < len(dataList.List) && dataList.List[index].Compare(data) == 0 {
		return index, true
	}

	return index, false
}
func (dataList *AnyDataList) Add(data ...data.Data) lists.DataList {
	updatedList := dataList
	for _, listable := range data {
		if index, found := updatedList.Search(listable); !found {
			updatedList.List = append(updatedList.List, listable.(*baseData.AnyData))
			copy(updatedList.List[index+1:], updatedList.List[index:])
			updatedList.List[index] = listable.(*baseData.AnyData)
		}
	}
	return updatedList
}

func (dataList *AnyDataList) Remove(data ...data.Data) lists.DataList {
	updatedList := dataList

	for _, listable := range data {
		if index, found := updatedList.Search(listable); found {
			updatedList.List = append(updatedList.List[:index], updatedList.List[index+1:]...)
		}
	}

	return updatedList
}

func NewDataList(data ...data.Data) lists.DataList {
	var dataList []*baseData.AnyData
	for _, dataVal := range data {
		dataList = append(dataList, dataVal.(*baseData.AnyData))
	}
	return &AnyDataList{List: dataList}
}
