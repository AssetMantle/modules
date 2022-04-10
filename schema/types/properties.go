// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

// TODO Do as list
type Properties interface {
	Get(ID) Property

	GetList() []Property

	Add(...Property) Properties
	Remove(...Property) Properties
	Mutate(...Property) Properties
}
