/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type MetaProperties interface {
	GetMetaProperty(ID) MetaProperty

	GetMetaPropertyList() []MetaProperty

	AddMetaProperty(MetaProperty) MetaProperties
	RemoveMetaProperty(MetaProperty) MetaProperties
	MutateMetaProperty(MetaProperty) MetaProperties

	RemoveData() Properties

	Properties
}
