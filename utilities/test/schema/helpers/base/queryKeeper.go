// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/AssetMantle/schema/x/helpers"
)

type TestQueryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*TestQueryKeeper)(nil)

func (t TestQueryKeeper) Help(_ context.Context, _ helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	return nil, nil
}

func (t TestQueryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return TestQueryKeeper{mapper: mapper}
}

func (t TestQueryKeeper) Enquire(_ context.Context, _ helpers.QueryRequest) helpers.QueryResponse {
	return &TestQueryResponse{}
}

func TestQueryKeeperPrototype() helpers.QueryKeeper {
	return TestQueryKeeper{}
}

type testQueryRequest struct {
}

var _ helpers.QueryRequest = (*testQueryRequest)(nil)

func (testQueryRequest testQueryRequest) Validate() error {
	return nil
}

func (testQueryRequest testQueryRequest) FromCLI(_ helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	return testQueryRequest, nil
}

func (testQueryRequest testQueryRequest) FromMap(_ map[string]string) (helpers.QueryRequest, error) {
	return testQueryRequest, nil
}

func (testQueryRequest testQueryRequest) Encode() ([]byte, error) {
	return json.Marshal(testQueryRequest)
}

func (testQueryRequest) Decode(bytes []byte) (helpers.QueryRequest, error) {
	var queryRequest testQueryRequest
	err := json.Unmarshal(bytes, &queryRequest)

	return queryRequest, err
}

func TestQueryRequestPrototype() helpers.QueryRequest {
	return testQueryRequest{}
}

var _ helpers.QueryResponse = (*TestQueryResponse)(nil)

func (t *TestQueryResponse) IsSuccessful() bool {
	return t.Success
}

func (t *TestQueryResponse) Encode() ([]byte, error) {
	return json.Marshal(t)
}

func (t *TestQueryResponse) GetError() error {
	return errors.New(t.Error)
}

func (t *TestQueryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	var queryResponse TestQueryResponse
	err := json.Unmarshal(bytes, &queryResponse)

	return &queryResponse, err
}
func TestQueryResponsePrototype() helpers.QueryResponse {
	return &TestQueryResponse{}
}
