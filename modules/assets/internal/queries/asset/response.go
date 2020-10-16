/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package asset

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryResponse struct {
	Success bool               `json:"success"`
	Error   error              `json:"error"`
	Assets  helpers.Collection `json:"assets" valid:"required~required field assets missing"`
}

var _ helpers.QueryResponse = (*queryResponse)(nil)

func (queryResponse queryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse queryResponse) GetError() error {
	return queryResponse.Error
}
func (queryResponse queryResponse) Encode() ([]byte, error) {
	//todo
	panic("implement me")
}

func (queryResponse queryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	//todo
	panic("implement me")
}
func responsePrototype() helpers.QueryResponse {
	return queryResponse{}
}
func newQueryResponse(assets helpers.Collection, error error) helpers.QueryResponse {
	success := true
	if error != nil {
		success = false
	}
	return queryResponse{
		Success: success,
		Error:   error,
		Assets:  assets,
	}
}
