// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
)

type dataList struct {
	lists.List
}

var _ lists.DataList = (*dataList)(nil)

func (dataList dataList) GetList() []data.DataI {
	DataList := make([]data.DataI, dataList.List.Size())

	for i, listable := range dataList.List.Get() {
		DataList[i] = listable.(data.DataI)
	}

	return DataList
}
func (dataList dataList) Search(data data.DataI) (int, bool) {
	return dataList.List.Search(data)
}
func (dataList dataList) Add(data ...data.DataI) lists.DataList {
	dataList.List = dataList.List.Add(dataToListables(data...)...)
	return dataList
}
func (dataList dataList) Remove(data ...data.DataI) lists.DataList {
	dataList.List = dataList.List.Remove(dataToListables(data...)...)
	return dataList
}
func dataToListables(data ...data.DataI) []traits.Listable {
	listables := make([]traits.Listable, len(data))

	for i, datum := range data {
		listables[i] = datum
	}

	return listables
}

func NewDataList(data ...data.DataI) lists.DataList {
	return dataList{List: NewList(dataToListables(data...)...)}
}
