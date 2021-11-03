/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"sort"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

//var _ types.SortedDataList = (*ListData)(nil)

func (listData ListData) Search(data types.Data) int {
	index := sort.Search(
		len(listData.Value),
		func(i int) bool {
			return listData.Value[i].Compare(data) <= 0
		},
	)
	if index < len(listData.Value) {
		if listData.Value[index].Compare(data) != 0 {
			return len(listData.Value)
		}
	}
	return index
}
func (listData ListData) GetList() []types.Data {
	panic("implement me")
}

//func (listData ListData) Add(dataList ...types.Data) types.SortedDataList {
//	panic("implement me")
//}
//func (listData ListData) Remove(dataList ...types.Data) types.SortedDataList {
//	for _, data := range dataList {
//		if index := listData.Search(data); index != len(listData.Value) {
//			listData.Value = append(listData.Value[:index], listData.Value[index+1:]...)
//		}
//	}
//	return listData
//}

func (listData *ListData) Add(data ...types.Data) types.ListData {
	panic("implement me")
}

func (listData *ListData) Remove(data ...types.Data) types.ListData {
	panic("implement me")
}
