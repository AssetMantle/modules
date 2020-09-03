/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package identity

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
)

type queryResponse struct {
	Success    bool                    `json:"success"`
	Error      error                   `json:"error"`
	Identities mappers.InterIdentities `json:"identities" valid:"required~required field identities missing"`
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

func NewQueryResponse(identities mappers.InterIdentities, error error) helpers.QueryResponse {
	success := true
	if error != nil {
		success = false
	}
	return queryResponse{
		Success:    success,
		Error:      error,
		Identities: identities,
	}
}
