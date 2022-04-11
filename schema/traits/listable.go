// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package traits

type Listable interface {
	// Compare
	// * panic if compared with Listable of different type
	Compare(Listable) int
}
