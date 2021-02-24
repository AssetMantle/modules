/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type Properties interface {
	Get(ID) Property

	GetList() []Property

	Add(...Property) Properties
	Remove(...Property) Properties
	Mutate(...Property) Properties
}
