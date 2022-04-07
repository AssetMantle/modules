// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/AssetMantle/modules/schema/traits"
)

type List interface {
	// TODO add search and apply methods
	GetList() []traits.Listable
	Size() int

	Search(listable traits.Listable) (bool, int)
	Add(listableList ...traits.Listable) List
	Remove(listableList ...traits.Listable) List
}
