// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
	baseList "github.com/AssetMantle/modules/schema/types/base"
)

type dataList struct {
	types.List
}

var _ lists.DataList = (*dataList)(nil)

func (dataList dataList) GetList() []types.Data {
	data := make([]types.Data, dataList.Size())

	for i, listable := range dataList.Get() {
		data[i] = listable.(types.Data)
	}

	return data
}
func (dataList dataList) Search(data types.Data) (int, bool) {
	return dataList.List.Search(data)
}
func dataToListables(data ...types.Data) []capabilities.Listable {
	listables := make([]capabilities.Listable, len(data))

	for i, datum := range data {
		listables[i] = datum
	}

	return listables
}

func NewDataList(data ...types.Data) lists.DataList {
	return dataList{List: baseList.NewList(dataToListables(data...)...)}
}
