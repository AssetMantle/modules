// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package metas

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/metas/record"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(collection helpers.Collection) *QueryResponse {
	list := record.RecordsFromInterface(collection.Get())

	return &QueryResponse{
		List: list,
	}
}
