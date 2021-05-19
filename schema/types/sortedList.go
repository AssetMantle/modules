/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import "sort"

type SortedList interface {
	sort.Interface

	Sort() SortedList

	Insert(interface{}) SortedList
	Delete(interface{}) SortedList
	Search(interface{}) int
}
