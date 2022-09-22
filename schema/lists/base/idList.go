// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
)

type idList struct {
	lists.List
}

var _ lists.IDList = (*idList)(nil)

func (idList idList) GetList() []ids.ID {
	returnIDList := make([]ids.ID, idList.List.Size())

	for i, listable := range idList.List.Get() {
		returnIDList[i] = listable.(ids.ID)
	}

	return returnIDList
}
func (idList idList) Search(id ids.ID) (index int, found bool) {
	return idList.List.Search(id)
}
func (idList idList) Add(ids ...ids.ID) lists.IDList {
	idList.List = idList.List.Add(idsToListables(ids...)...)
	return idList
}
func (idList idList) Remove(ids ...ids.ID) lists.IDList {
	idList.List = idList.List.Remove(idsToListables(ids...)...)
	return idList
}
func idsToListables(ids ...ids.ID) []traits.Listable {
	listables := make([]traits.Listable, len(ids))
	for i, id := range ids {
		listables[i] = id
	}
	return listables
}
