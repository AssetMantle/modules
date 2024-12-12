// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supply

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(supply math.Int) *QueryResponse {
	return &QueryResponse{Supply: supply.String()}
}
