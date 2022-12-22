// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package meta

import (
	"errors"
	"github.com/AssetMantle/modules/modules/metas/internal/common"
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
	return errors.New(queryResponse.Error)
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
	if error != nil {
		return &QueryResponse{
			Success: false,
			Error:   error.Error(),
			List:    collection.GetList(),
		}
	}

	return &QueryResponse{
		Success: true,
		Error:   "",
		List:    collection.GetList(),
	}
}
