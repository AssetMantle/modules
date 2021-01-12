/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	clientContext "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQuery(t *testing.T) {
	context, storeKey, _ := base.SetupTest(t)
	codec := base.MakeCodec()
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Query := NewQuery("test", "t", "testQuery", "test", base.TestQueryRequestPrototype,
		base.TestQueryResponsePrototype, base.TestQueryKeeperPrototype).Initialize(Mapper, parametersPrototype()).(query)

	require.Equal(t, nil, base.TestQueryKeeperPrototype().(base.TestQueryKeeper).Help(context, nil))
	require.Equal(t, nil, base.TestQueryRequestPrototype().Validate())
	require.Equal(t, false, base.TestQueryResponsePrototype().IsSuccessful())
	require.Equal(t, nil, base.TestQueryResponsePrototype().GetError())
	encodedResponse, Error := base.TestQueryResponsePrototype().Encode()
	require.Nil(t, Error)
	decodedResponse, Error := base.TestQueryResponsePrototype().Decode(encodedResponse)
	require.Nil(t, Error)
	require.Equal(t, Query.responsePrototype(), decodedResponse)

	// GetName
	require.Equal(t, "test", Query.GetName())

	// HandleMessage
	encodedRequest, Error := Query.requestPrototype().Encode()
	require.Nil(t, Error)

	_, Error = Query.HandleMessage(context, abciTypes.RequestQuery{Data: encodedRequest})
	require.Nil(t, Error)

	command := Query.Command(codec)
	command.SetArgs([]string{
		"test"})
	require.Equal(t, `ABCIQuery: Post failed: Post "http://localhost:26657": dial tcp 127.0.0.1:26657: connect: connection refused`,
		command.ExecuteContext(context.Context()).Error())

	// RESTQueryHandler
	cliContext := clientContext.NewCLIContext().WithCodec(codec).WithChainID("test")
	Query.RESTQueryHandler(cliContext)

	//RPC ERROR
	testRequest1, Error := http.NewRequest("GET", "/test", nil)
	require.Nil(t, Error)
	responseRecorder := httptest.NewRecorder()
	Query.RESTQueryHandler(cliContext).ServeHTTP(responseRecorder, testRequest1)
	require.Equal(t, responseRecorder.Code, http.StatusInternalServerError)

}
