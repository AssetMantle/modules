/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package ownable

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/common"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryResponse struct {
	Success bool         `json:"success"`
	Error   error        `json:"error" swaggertype:"string"`
	Value   sdkTypes.Dec `json:"value" swaggertype:"string"`
}

var _ helpers.QueryResponse = (*queryResponse)(nil)

func (queryResponse queryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse queryResponse) GetError() error {
	return queryResponse.Error
}
func (queryResponse queryResponse) Encode() ([]byte, error) {
	return common.Codec.MarshalJSON(queryResponse)
}
func (queryResponse queryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	if Error := common.Codec.UnmarshalJSON(bytes, &queryResponse); Error != nil {
		return nil, Error
	}

	return queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return queryResponse{}
}
func newQueryResponse(value sdkTypes.Dec, error error) helpers.QueryResponse {
	success := true
	if error != nil {
		success = false
	}

	return queryResponse{
		Success: success,
		Error:   error,
		Value:   value,
	}
}
