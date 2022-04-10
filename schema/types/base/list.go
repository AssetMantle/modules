// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"sort"

	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

type list []traits.Listable

var _ types.List = (*list)(nil)

func (list list) GetList() []traits.Listable {
	return list
}
func (list list) Size() int { // TODO write test
	return len(list)
}
func (list list) Search(listable traits.Listable) (bool, int) {
	index := sort.Search(
		len(list),
		func(i int) bool {
			return list[i].Compare(listable) <= 0
		},
	)

	if list[index].Compare(listable) == 0 {
		return true, index
	}

	return false, index
}
func (list list) Add(listableList ...traits.Listable) types.List {
	updatedList := list

	for _, listable := range listableList {
		if found, index := updatedList.Search(listable); !found {
			updatedList := append(updatedList, listable)
			copy(updatedList[index+1:], updatedList[index:])
			updatedList[index] = listable
		}
	}

	return updatedList
}
func (list list) Remove(listableList ...traits.Listable) types.List {
	updatedList := list

	for _, listable := range listableList {
		if found, index := updatedList.Search(listable); found {
			updatedList = append(updatedList[:index], updatedList[index+1:]...)
		}
	}

	return updatedList
}
