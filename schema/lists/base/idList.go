// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"sort"

	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
)

var _ lists.IDList = (*IDList)(nil)

func (idList *IDList) ValidateBasic() error {
	for _, id := range idList.IDList {
		if err := id.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}
func (idList *IDList) GetList() []ids.AnyID {
	returnIDList := make([]ids.AnyID, len(idList.IDList))

	for i, listable := range idList.IDList {
		returnIDList[i] = listable
	}

	return returnIDList
}
func (idList *IDList) Search(id ids.ID) (index int, found bool) {
	index = sort.Search(
		len(idList.IDList),
		func(i int) bool {
			return idList.IDList[i].Compare(id) >= 0
		},
	)

	if index < len(idList.IDList) && idList.IDList[index].Compare(id) == 0 {
		return index, true
	}

	return index, false
}
func (idList *IDList) Add(ids ...ids.ID) lists.IDList {
	updatedList := idList
	for _, listable := range ids {
		if index, found := updatedList.Search(listable); !found {
			updatedList.IDList = append(updatedList.IDList, listable.ToAnyID().(*baseIDs.AnyID))
			copy(updatedList.IDList[index+1:], updatedList.IDList[index:])
			updatedList.IDList[index] = listable.ToAnyID().(*baseIDs.AnyID)
		}
	}
	return updatedList
}
func (idList *IDList) Remove(ids ...ids.ID) lists.IDList {
	updatedList := idList

	for _, listable := range ids {
		if index, found := updatedList.Search(listable); found {
			updatedList.IDList = append(updatedList.IDList[:index], updatedList.IDList[index+1:]...)
		}
	}

	return updatedList
}
func (idList *IDList) sort() lists.IDList {
	sort.Slice(idList.IDList, func(i, j int) bool {
		return idList.IDList[i].Compare(idList.IDList[j]) < 0
	})
	return idList
}

func NewIDList(ids ...ids.ID) lists.IDList {
	return (&IDList{}).Add(ids...)
}
