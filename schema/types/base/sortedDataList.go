/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"sort"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type sortedDataList []types.Data

var _ types.SortedDataList = (*sortedDataList)(nil)

func (sortedDataList sortedDataList) Search(data types.Data) int {
	index := sort.Search(
		len(sortedDataList),
		func(i int) bool {
			return sortedDataList[i].Compare(data) <= 0
		},
	)

	if index < len(sortedDataList) {
		if sortedDataList[index].Compare(data) != 0 {
			return len(sortedDataList)
		}
	}

	return index
}

func (sortedDataList sortedDataList) GetList() []types.Data {
	return sortedDataList
}

func (sortedDataList sortedDataList) Add(dataList ...types.Data) types.SortedDataList {
	for _, data := range dataList {
		if sortedDataList.Search(data) != len(sortedDataList) {
			return sortedDataList
		}

		index := sort.Search(
			len(sortedDataList),
			func(i int) bool {
				return sortedDataList[i].Compare(data) < 0 // nolint
			},
		)

		//goland:noinspection GoAssignmentToReceiver
		sortedDataList = append(sortedDataList, data)
		copy(sortedDataList[index+1:], sortedDataList[index:])
		sortedDataList[index] = data
	}

	return sortedDataList
}

func (sortedDataList sortedDataList) Remove(dataList ...types.Data) types.SortedDataList {
	for _, data := range dataList {
		if index := sortedDataList.Search(data); index != len(sortedDataList) {
			sortedDataList = append(sortedDataList[:index], sortedDataList[index+1:]...)
		}
	}

	return sortedDataList
}
