// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/AssetMantle/modules/utilities/test"
	"github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

func TestQuery(t *testing.T) {
	context, storeKey, _ := test.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Query := NewQuery("test", "t", "testQuery", "test", base.TestQueryRequestPrototype,
		base.TestQueryResponsePrototype, base.TestQueryKeeperPrototype,
		nil,
		nil,
	).Initialize(Mapper, parameterManagerPrototype()).(query)

	require.Equal(t, nil, base.TestQueryKeeperPrototype().(base.TestQueryKeeper).Help(context, nil))
	require.Equal(t, nil, base.TestQueryRequestPrototype().Validate())
	require.Equal(t, false, base.TestQueryResponsePrototype().IsSuccessful())
	require.Equal(t, nil, base.TestQueryResponsePrototype().GetError())
	encodedResponse, err := base.TestQueryResponsePrototype().Encode()
	require.Nil(t, err)
	decodedResponse, err := base.TestQueryResponsePrototype().Decode(encodedResponse)
	require.Nil(t, err)
	require.Equal(t, Query.responsePrototype(), decodedResponse)

	// GetName
	require.Equal(t, "test", Query.GetName())

	// HandleQuery
	encodedRequest, err := Query.requestPrototype().Encode()
	require.Nil(t, err)

	_, err = Query.HandleQuery(context, abciTypes.RequestQuery{Data: encodedRequest})
	require.Nil(t, err)

	command := Query.Command()
	command.SetArgs([]string{
		"test"})
	err = command.ParseFlags([]string{"--node", "tcp://localhost:26657"})
	require.Nil(t, err)
	require.Equal(t, `ABCIQuery: Post failed: Post "http://localhost:26657": dial tcp 127.0.0.1:26657: connect: connection refused`,
		command.ExecuteContext(context).Error())

	// require.Equal(t, nil, command.ExecuteContext(context.Context()))

	// RESTQueryHandler
	Query.RESTQueryHandler(TestClientContext)

	// RPC ERROR
	testRequest1, err := http.NewRequest("GET", "/test", nil)
	require.Nil(t, err)
	responseRecorder := httptest.NewRecorder()
	Query.RESTQueryHandler(TestClientContext).ServeHTTP(responseRecorder, testRequest1)
	require.Equal(t, responseRecorder.Code, http.StatusInternalServerError)

}
