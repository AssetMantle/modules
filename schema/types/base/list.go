// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"sort"

	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

// TODO type check that list is same type
type list []traits.Listable

var _ types.List = (*list)(nil)

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
			return list[i].Compare(listable) <= 0
		},
	)

	if list[index].Compare(listable) == 0 {
		return index, true
	}

	return index, false
}
func (list list) Add(listableList ...traits.Listable) types.List {
	updatedList := list

	for _, listable := range listableList {
		if index, found := updatedList.Search(listable); !found {
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
		if index, found := updatedList.Search(listable); found {
			updatedList = append(updatedList[:index], updatedList[index+1:]...)
		}
	}

	return updatedList
}
func (list list) Mutate(listable ...traits.Listable) types.List {
	// TODO implement me
	panic("implement me")
}

func NewList(listables ...traits.Listable) types.List {
	// TODO write test and check type
	return list(listables)
}
