// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"github.com/cosmos/cosmos-sdk/types"
)

var Invariant = func(_ types.Context) (string, bool) {
	return "", false
}
