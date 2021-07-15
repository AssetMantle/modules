/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"sort"
)

type SortedList interface {
	sort.Interface

	Sort() SortedList

	Insert(sortable traits.Sortable) SortedList
	Delete(sortable traits.Sortable) SortedList
	Search(sortable traits.Sortable) int
}
