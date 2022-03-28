/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"sort"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.SortedDataList = (*ListData)(nil)

func (sortedDataList ListData) Search(data types.Data) int {
	index := sort.Search(
		len(sortedDataList.Value),
		func(i int) bool {
			return sortedDataList.Value[i].Compare(data) <= 0
		},
	)
	if index < len(sortedDataList.Value) {
		if sortedDataList.Value[index].Compare(data) != 0 {
			return len(sortedDataList.Value)
		}
	}
	return index
}
func (sortedDataList ListData) GetList() []types.Data {
	newList := make([]types.Data, len(sortedDataList.Value))
	for i, _ := range sortedDataList.Value {
		newList[i] = &sortedDataList.Value[i]
	}
	return newList
}
func (sortedDataList ListData) Add(dataList ...types.Data) types.SortedDataList {
	for _, data := range dataList {
		if sortedDataList.Search(data) != len(sortedDataList.Value) {
			return sortedDataList
		}

		index := sort.Search(
			len(sortedDataList.Value),
			func(i int) bool {
				return sortedDataList.Value[i].Compare(data) < 0
			},
		)
		newData := *NewData(data)
		sortedDataList.Value = append(sortedDataList.Value, newData)
		copy(sortedDataList.Value[index+1:], sortedDataList.Value[index:])
		sortedDataList.Value[index] = *NewData(data)
	}
	return sortedDataList
}
func (sortedDataList ListData) Remove(dataList ...types.Data) types.SortedDataList {
	for _, data := range dataList {
		if index := sortedDataList.Search(data); index != len(sortedDataList.Value) {
			sortedDataList.Value = append(sortedDataList.Value[:index], sortedDataList.Value[index+1:]...)
		}
	}
	return sortedDataList
}
