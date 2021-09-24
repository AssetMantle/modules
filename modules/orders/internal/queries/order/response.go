/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package order

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/common"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)
func (queryResponse QueryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse QueryResponse) GetError() error {
	return errors.New(queryResponse.Error)
}
func (queryResponse QueryResponse) LegacyAminoEncode() ([]byte, error) {
	return common.LegacyAminoCodec.MarshalJSON(queryResponse)
}
func (queryResponse QueryResponse) LegacyAminoDecode(bytes []byte) (helpers.QueryResponse, error) {
	if Error := common.LegacyAminoCodec.UnmarshalJSON(bytes, &queryResponse); Error != nil {
		return nil, Error
	}

	return &queryResponse, nil
}
func (queryResponse QueryResponse) Encode(cdc codec.JSONMarshaler) ([]byte, error) {
	return cdc.MarshalJSON(&queryResponse)
}

func (queryResponse QueryResponse) Decode(cdc codec.JSONMarshaler, bytes []byte) (helpers.QueryResponse, error) {
	if Error := cdc.UnmarshalJSON(bytes, &queryResponse); Error != nil {
		return nil, Error
	}
	return &queryResponse, nil
}

func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}

func newQueryResponse(collection helpers.Collection, error error) QueryResponse {
	success := true
	if error != nil {
		success = false
		return QueryResponse{
			Success: success,
			Error:   "yes",
			List:    collection.GetList(),
		}
	}
	return QueryResponse{
		Success: success,
		Error:   "no",
		List:    collection.GetList(),
	}

}