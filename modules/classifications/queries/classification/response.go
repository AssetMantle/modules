/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package classification

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
)

type queryResponse struct {
	Success         bool                    `json:"success"`
	Error           error                   `json:"error"`
	Classifications mappers.Classifications `json:"classifications" valid:"required~required field classifications missing"`
}

var _ helpers.QueryResponse = (*queryResponse)(nil)

func (queryResponse queryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse queryResponse) GetError() error {
	return queryResponse.Error
}
func queryResponsePrototype() helpers.QueryResponse {
	return queryResponse{}
}
func newQueryResponse(classifications mappers.Classifications, error error) helpers.QueryResponse {
	success := true
	if error != nil {
		success = false
	}
	return queryResponse{
		Success:         success,
		Error:           error,
		Classifications: classifications,
	}
}
