/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type TestQueryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*TestQueryKeeper)(nil)

func (t TestQueryKeeper) Help(_ sdkTypes.Context, _ helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	return nil
}

func (t TestQueryKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return TestQueryKeeper{mapper: mapper}
}

func (t TestQueryKeeper) Enquire(_ sdkTypes.Context, _ helpers.QueryRequest) helpers.QueryResponse {
	return testQueryResponse{}
}

func TestQueryKeeperPrototype() helpers.QueryKeeper {
	return TestQueryKeeper{}
}

type testQueryRequest struct {
}

var _ helpers.QueryRequest = (*testQueryRequest)(nil)

func (t testQueryRequest) Validate() error {
	return nil
}

func (t testQueryRequest) FromCLI(_ helpers.CLICommand, _ context.CLIContext) helpers.QueryRequest {
	return t
}

func (t testQueryRequest) FromMap(_ map[string]string) helpers.QueryRequest {
	return t
}

func (t testQueryRequest) Encode() ([]byte, error) {
	return json.Marshal(t)
}

func (t testQueryRequest) Decode(bytes []byte) (helpers.QueryRequest, error) {
	var queryRequest testQueryRequest
	Error := json.Unmarshal(bytes, &queryRequest)

	return queryRequest, Error
}

func TestQueryRequestPrototype() helpers.QueryRequest {
	return testQueryRequest{}
}

type testQueryResponse struct {
	Success bool
	Error   error
}

var _ helpers.QueryResponse = (*testQueryResponse)(nil)

func (t testQueryResponse) IsSuccessful() bool {
	return t.Success
}

func (t testQueryResponse) GetError() error {
	return t.Error
}

func (t testQueryResponse) Encode() ([]byte, error) {
	return json.Marshal(t)
}

func (t testQueryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	var queryResponse testQueryResponse
	Error := json.Unmarshal(bytes, &queryResponse)

	return queryResponse, Error
}
func TestQueryResponsePrototype() helpers.QueryResponse {
	return testQueryResponse{}
}
