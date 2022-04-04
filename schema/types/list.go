// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
)

type List interface {
	// TODO add search and apply methods
	GetList() []traits.Listable
	Size() int

	// TODO -1 for not found
	Search(listable traits.Listable) int
	Add(listableList ...traits.Listable) List
	Remove(listableList ...traits.Listable) List
}
