// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package meta

import (
	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/modules/metas/module/queries/meta"

	"github.com/AssetMantle/modules/modules/metas/module/common"
	"github.com/AssetMantle/modules/schema/helpers"
)

type queryResponse meta.QueryResponse

var _ helpers.QueryResponse = (*queryResponse)(nil)

func (queryResponse *queryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse *queryResponse) GetError() error {
	return queryResponse.Error
}
func (queryResponse *queryResponse) Encode() ([]byte, error) {
	return common.Codec.MarshalJSON(queryResponse)
}
func (queryResponse *queryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	if err := common.Codec.UnmarshalJSON(bytes, &queryResponse); err != nil {
		return nil, err
	}

	return queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return &queryResponse{}
}
func newQueryResponse(collection helpers.Collection, error error) helpers.QueryResponse {
	success := true
	if error != nil {
		success = false
	}

	return &queryResponse{
		Success: success,
		Error:   error,
		List:    collection.GetList(),
	}
}
