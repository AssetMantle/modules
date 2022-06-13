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

func (dataList dataList) GetList() []data.Data {
	DataList := make([]data.Data, dataList.List.Size())

	for i, listable := range dataList.List.Get() {
		DataList[i] = listable.(data.Data)
	}

	return DataList
}
func (dataList dataList) Search(data data.Data) (int, bool) {
	return dataList.List.Search(data)
}
func dataToListables(data ...data.Data) []traits.Listable {
	listables := make([]traits.Listable, len(data))

	for i, datum := range data {
		listables[i] = datum
	}

	return listables
}

func NewDataList(data ...data.Data) lists.DataList {
	return dataList{List: NewList(dataToListables(data...)...)}
}
