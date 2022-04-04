// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"sort"

	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

// TODO Revise the code
type list []traits.Listable

func (list list) Search(f func()) int {
	// TODO implement me
	panic("implement me")
}

func (list list) Apply(f func()) types.List {
	// TODO implement me
	panic("implement me")
}

func (list list) Mutate(listable ...traits.Listable) types.List {
	// TODO implement me
	panic("implement me")
}

var _ types.List = (*list)(nil)

func (list list) SearchListable(listable traits.Listable) int {
	index := sort.Search(
		len(list),
		func(i int) bool {
			return list[i].Compare(listable) <= 0
		},
	)

	if index < len(list) {
		if list[index].Compare(listable) != 0 {
			return len(list)
		}
	}

	return index
}

func (list list) GetList() []traits.Listable {
	return list
}

func (list list) Add(listableList ...traits.Listable) types.List {
	for _, listable := range listableList {
		if list.SearchListable(listable) != len(list) {
			return list
		}

		index := sort.Search(
			len(list),
			func(i int) bool {
				return list[i].Compare(listable) < 0 // nolint
			},
		)

		list = append(list, listable)
		copy(list[index+1:], list[index:])
		list[index] = listable
	}

	return list
}

func (list list) Remove(listableList ...traits.Listable) types.List {
	// TODO check if the return is properly assigned to
	for _, listable := range listableList {
		if index := list.SearchListable(listable); index != len(list) {
			list = append(list[:index], list[index+1:]...)
		}
	}

	return list
}
