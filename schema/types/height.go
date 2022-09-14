// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

// Height type to represent blockchain's block height
// 0 height === chain genesis block height
// -1 height === infinite block height
type Height interface {
	Get() int64

	Compare(Height) int
}
