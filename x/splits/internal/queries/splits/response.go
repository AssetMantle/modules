// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package splits

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/internal/mappable"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func (queryResponse *QueryResponse) Encode() ([]byte, error) {
	return base.CodecPrototype().MarshalJSON(queryResponse)
}
func (queryResponse *QueryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	if err := base.CodecPrototype().UnmarshalJSON(bytes, queryResponse); err != nil {
		return nil, err
	}

	return queryResponse, err
}
func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(collection helpers.Collection) *QueryResponse {
	list := mappable.MappablesFromInterface(collection.GetList())

	return &QueryResponse{
		List: list,
	}
}
