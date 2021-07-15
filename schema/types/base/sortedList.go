/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"sort"
)

type sortedList []traits.Sortable

var _ types.SortedList = (*sortedList)(nil)

func (sortedList sortedList) Len() int {
	return len(sortedList)
}
func (sortedList sortedList) Less(i, j int) bool {
	return sortedList[i].Compare(sortedList[j]) < 0
}
func (sortedList sortedList) Swap(i, j int) {
	sortedList[i], sortedList[j] = sortedList[j], sortedList[i]
}
func (sortedList sortedList) Sort() types.SortedList {
	sort.Sort(sortedList)
	return sortedList
}
func (sortedList sortedList) Insert(sortable traits.Sortable) types.SortedList {
	if sortedList.Search(sortable) != sortedList.Len() {
		return sortedList
	}

	index := sort.Search(
		sortedList.Len(),
		func(i int) bool {
			return sortedList[i].Compare(sortable) < 0
		},
	)

	newSortedList := append(sortedList, sortable)
	copy(newSortedList[index+1:], newSortedList[index:])
	newSortedList[index] = sortable

	return newSortedList
}
func (sortedList sortedList) Delete(sortable traits.Sortable) types.SortedList {
	index := sortedList.Search(sortable)

	if index == sortedList.Len() {
		return sortedList
	}

	return append(sortedList[:index], sortedList[index+1:]...)
}
func (sortedList sortedList) Search(sortable traits.Sortable) int {
	return sort.Search(
		sortedList.Len(),
		func(i int) bool {
			return sortedList[i].Compare(sortable) == 0
		},
	)
}
