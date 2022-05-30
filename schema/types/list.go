// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/AssetMantle/modules/schema/traits"
)

// List
// * all elements are sorted
// * all methods are search and insertion complexity optimized
type List interface {
	// TODO add search and apply methods
	Get() []traits.Listable
	// Size
	// * returns the number of elements in the list
	Size() int

	// Search
	// * returns true and index of element if element is found
	// * return false and index of insertion if element is not found
	Search(traits.Listable) (index int, found bool) // TODO prevent compare panic
	Add(...traits.Listable) List                    // TODO prevent compare panic
	Remove(...traits.Listable) List                 // TODO prevent compare panic
	Mutate(...traits.Listable) List                 // TODO prevent compare panic
}
