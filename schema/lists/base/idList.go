// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

type idList struct {
	types.List
}

var _ lists.IDList = (*idList)(nil)

func (idList idList) GetList() []types.ID {
	// TODO write test case
	returnIDList := make([]types.ID, idList.Size())

	for i, listable := range idList.GetList() {
		returnIDList[i] = listable.(types.ID)
	}

	return returnIDList
}
func (idList idList) Search(id types.ID) (index int, found bool) {
	return idList.List.Search(id)
}
func (idList idList) Add(ids ...types.ID) lists.IDList {
	idList.List = idList.List.Add(idsToListables(ids...)...)
	return idList
}
func (idList idList) Remove(ids ...types.ID) lists.IDList {
	idList.List = idList.List.Remove(idsToListables(ids...)...)
	return idList
}
func idsToListables(ids ...types.ID) []traits.Listable {
	listables := make([]traits.Listable, len(ids))
	for i, id := range ids {
		listables[i] = id
	}
	return listables
}
