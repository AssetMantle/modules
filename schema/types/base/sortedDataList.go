/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"sort"
)

type sortedDataList []types.Data

var _ types.SortedDataList = (*sortedDataList)(nil)

func (sortedDataList sortedDataList) Insert(data types.Data) types.SortedDataList {
	if sortedDataList.Search(data) != len(sortedDataList) {
		return sortedDataList
	}

	index := sort.Search(
		len(sortedDataList),
		func(i int) bool {
			return sortedDataList[i].Compare(data) < 0
		},
	)

	newSortedList := append(sortedDataList, data)
	copy(newSortedList[index+1:], newSortedList[index:])
	newSortedList[index] = data

	return newSortedList
}
func (sortedDataList sortedDataList) Delete(data types.Data) types.SortedDataList {
	index := sortedDataList.Search(data)

	if index == len(sortedDataList) {
		return sortedDataList
	}

	return append(sortedDataList[:index], sortedDataList[index+1:]...)
}
func (sortedDataList sortedDataList) Search(data types.Data) int {
	return sort.Search(
		len(sortedDataList),
		func(i int) bool {
			return sortedDataList[i].Compare(data) == 0
		},
	)
}
