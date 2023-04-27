// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ownable

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func (queryResponse *QueryResponse) Encode() ([]byte, error) {
	return base.CodecPrototype().MarshalJSON(queryResponse)
}
func (queryResponse *QueryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	if err := base.CodecPrototype().UnmarshalJSON(bytes, queryResponse); err != nil {
		return nil, err
	}

	return queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(value sdkTypes.Int, error error) *QueryResponse {
	if error != nil {
		return &QueryResponse{

			Value: value.String(),
		}
	}

	return &QueryResponse{
		Error: "",
		Value: value.String(),
	}
}
