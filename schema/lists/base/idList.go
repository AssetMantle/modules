// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"sort"
)

var _ lists.IDList = (*IDList)(nil)

func (idList *IDList) Search(id ids.ID) (index int, found bool) {
	index = sort.Search(
		len(idList.List),
		func(i int) bool {
			return idList.List[i].Compare(id) >= 0
		},
	)

	if index < len(idList.List) && idList.List[index].Compare(id) == 0 {
		return index, true
	}

	return index, false
}
func (idList *IDList) Add(ids ...ids.ID) lists.IDList {
	updatedList := idList
	for _, listable := range ids {
		if index, found := updatedList.Search(listable); !found {
			updatedList.List = append(updatedList.List, listable.(*baseIDs.ID))
			copy(updatedList.List[index+1:], updatedList.List[index:])
			updatedList.List[index] = listable.(*baseIDs.ID)
		}
	}
	return updatedList
}
func (idList *IDList) Remove(ids ...ids.ID) lists.IDList {
	updatedList := idList

	for _, listable := range ids {
		if index, found := updatedList.Search(listable); found {
			updatedList.List = append(updatedList.List[:index], updatedList.List[index+1:]...)
		}
	}

	return updatedList
}
func NewIDList(ids ...ids.ID) lists.IDList {
	var idList []*baseIDs.ID
	for _, dataVal := range ids {
		idList = append(idList, dataVal.(*baseIDs.ID))
	}
	return &IDList{List: idList}
}
