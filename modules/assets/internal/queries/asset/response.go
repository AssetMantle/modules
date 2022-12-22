// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package asset

import (
	"github.com/AssetMantle/modules/modules/assets/internal/common"
	"github.com/AssetMantle/modules/schema/helpers"
)

//type queryResponse struct {
//	Success bool               `json:"success"`
//	Error   error              `json:"error" swaggertype:"string"`
//	List    []helpers.Mappable `json:"list"`
//}

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func (queryResponse *QueryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse *QueryResponse) GetError() error {
	return queryResponse.Error
}
func (queryResponse *QueryResponse) Encode() ([]byte, error) {
	return common.LegacyAmino.MarshalJSON(queryResponse)
}
func (queryResponse *QueryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	if err := common.LegacyAmino.UnmarshalJSON(bytes, &queryResponse); err != nil {
		return nil, err
	}

	return queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(collection helpers.Collection, error error) helpers.QueryResponse {
	success := true
	if error != nil {
		success = false
	}

	return &QueryResponse{
		Success: success,
		Error:   error,
		List:    collection.GetList(),
	}
}
