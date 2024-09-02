// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/record"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(Record helpers.Record) *QueryResponse {
	return &QueryResponse{
		Record: Record.(*record.Record),
	}
}
