/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type Traits interface {
	Get(ID) Trait
	GetList() []Trait

	Add(Trait) Traits
	Remove(Trait) Traits
	Mutate(Trait) Traits
}
