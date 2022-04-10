// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

// TODO do as list
type MetaProperties interface {
	Get(ID) MetaProperty

	GetList() []MetaProperty

	Add(...MetaProperty) MetaProperties
	Remove(...MetaProperty) MetaProperties
	Mutate(...MetaProperty) MetaProperties

	RemoveData() Properties
}
