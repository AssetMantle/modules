// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"sort"
)

var _ lists.IDList = (*List_IdList)(nil)

func (idList *List_IdList) GetList() []ids.ID {
	returnIDList := make([]ids.ID, len(idList.IdList.List))

	for i, listable := range idList.IdList.List {
		returnIDList[i] = listable
	}

	return returnIDList
}
func (idList *List_IdList) Search(id ids.ID) (index int, found bool) {
	index = sort.Search(
		len(idList.IdList.List),
		func(i int) bool {
			return idList.IdList.List[i].Compare(id) >= 0
		},
	)

	if index < len(idList.IdList.List) && idList.IdList.List[index].Compare(id) == 0 {
		return index, true
	}

	return index, false
}
func (idList *List_IdList) Add(ids ...ids.ID) lists.List {
	updatedList := idList
	for _, listable := range ids {
		if index, found := updatedList.Search(listable); !found {
			updatedList.IdList.List = append(updatedList.IdList.List, listable.(*baseIDs.ID))
			copy(updatedList.IdList.List[index+1:], updatedList.IdList.List[index:])
			updatedList.IdList.List[index] = listable.(*baseIDs.ID)
		}
	}
	return &List{
		Impl: updatedList,
	}
}
func (idList *List_IdList) Remove(ids ...ids.ID) lists.List {
	updatedList := idList

	for _, listable := range ids {
		if index, found := updatedList.Search(listable); found {
			updatedList.IdList.List = append(updatedList.IdList.List[:index], updatedList.IdList.List[index+1:]...)
		}
	}

	return &List{
		Impl: updatedList,
	}
}
func NewIDList(ids ...ids.ID) lists.List {
	var idList []*baseIDs.ID
	for _, dataVal := range ids {
		idList = append(idList, dataVal.(*baseIDs.ID))
	}
	return &List{
		Impl: &List_IdList{
			IdList: &IDList{
				List: idList,
			},
		},
	}
}
