// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package traits

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/traits"
)

type Listable interface {
	// Compare
	// * panic if compared with Listable of different type
	// ***** TODO remove panic on compare with different type
	Compare(Listable) int
}

type listable traits.Listable

var _ Listable = (*listable)(nil)

func (l listable) Compare(l2 Listable) int {
	return l.Impl.(Listable).Compare(l2)
}
