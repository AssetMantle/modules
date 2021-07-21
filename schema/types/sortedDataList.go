/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type SortedDataList interface {
	Search(Data) int

	GetList() []Data

	Add(...Data) SortedDataList
	Remove(...Data) SortedDataList
}
