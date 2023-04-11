// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"errors"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/internal/mappable"
)

// type queryResponse struct {
//	Success bool               `json:"success"`
//	Error   error              `json:"error" swaggertype:"string"`
//	List    []helpers.Mappable `json:"list"`
// }

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func (queryResponse *QueryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse *QueryResponse) GetError() error {
	return errors.New(queryResponse.Error)
}
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
func newQueryResponse(collection helpers.Collection, error error) *QueryResponse {
	list := mappable.MappablesFromInterface(collection.GetList())
	if error != nil {
		return &QueryResponse{
			Success: false,
			Error:   error.Error(),
			List:    list,
		}
	}

	return &QueryResponse{
		Success: true,
		List:    list,
	}
}
