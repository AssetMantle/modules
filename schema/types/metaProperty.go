/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type MetaProperty interface {
	GetMetaFact() MetaFact
	RemoveData() Property
	Property
}
