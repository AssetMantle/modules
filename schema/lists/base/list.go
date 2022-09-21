// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"sort"

	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
)

// TODO type check that list is same type and test cases
type list []traits.Listable

var _ lists.List = (*list)(nil)

func (list list) Get() []traits.Listable {
	return list
}
func (list list) Size() int { // TODO write test
	return len(list)
}
func (list list) Search(listable traits.Listable) (int, bool) {
	index := sort.Search(
		len(list),
		func(i int) bool {
			return list[i].Compare(listable) >= 0
		},
	)

	if index < len(list) && list[index].Compare(listable) == 0 {
		return index, true
	}

	return index, false
}
func (list list) Add(listables ...traits.Listable) lists.List {
	updatedList := list

	for _, listable := range listables {
		if index, found := updatedList.Search(listable); !found {
			updatedList = append(updatedList, listable)
			copy(updatedList[index+1:], updatedList[index:])
			updatedList[index] = listable
		}
	}

	return updatedList
}
func (list list) Remove(listables ...traits.Listable) lists.List {
	updatedList := list

	for _, listable := range listables {
		if index, found := updatedList.Search(listable); found {
			updatedList = append(updatedList[:index], updatedList[index+1:]...)
		}
	}

	return updatedList
}
func (list list) Mutate(listables ...traits.Listable) lists.List {
	// TODO write test
	updatedList := list

	for _, listable := range listables {
		if index, found := updatedList.Search(listable); found {
			updatedList[index] = listable
		}
	}

	return updatedList
}

func NewList(listables ...traits.Listable) lists.List {
	list := list(listables)
	sort.Slice(list, func(i, j int) bool {
		return list[i].Compare(list[j]) <= 0
	})

	return list
}
