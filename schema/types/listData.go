/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type ListData interface {
	Data

	Search(Data) int

	GetList() []Data

	Add(...Data) ListData
	Remove(...Data) ListData
}

type DummyListData interface {
	Data

	BaseSearch(Data) int

	BaseGetList() []Data

	BaseAdd(...Data) DummyListData
	BaseRemove(...Data) DummyListData
}
