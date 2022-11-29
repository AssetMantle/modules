// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package meta

import (
	"github.com/AssetMantle/modules/modules/metas/internal/common"
	"github.com/AssetMantle/modules/schema/helpers"
)

type queryResponse struct {
	Success bool               `json:"success"`
	Error   error              `json:"error" swaggertype:"string"`
	List    []helpers.Mappable `json:"list"`
}

func (queryResponse queryResponse) Reset() {
	// TODO implement me
	panic("implement me")
}

func (queryResponse queryResponse) String() string {
	// TODO implement me
	panic("implement me")
}

func (queryResponse queryResponse) ProtoMessage() {
	// TODO implement me
	panic("implement me")
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
	if err := common.Codec.UnmarshalJSON(bytes, &queryResponse); err != nil {
		return nil, err
	}

	return queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return queryResponse{}
}
func newQueryResponse(collection helpers.Collection, error error) helpers.QueryResponse {
	success := true
	if error != nil {
		success = false
	}

	return queryResponse{
		Success: success,
		Error:   error,
		List:    collection.GetList(),
	}
}
