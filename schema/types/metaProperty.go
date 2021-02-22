/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type MetaProperty interface {
	GetID() ID
	GetMetaFact() MetaFact
	RemoveData() Property
}
