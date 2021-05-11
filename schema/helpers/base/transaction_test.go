/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	"encoding/json"
	clientContext "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransaction(t *testing.T) {
	codec := base.MakeCodec()
	context, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Transaction := NewTransaction("test", "", "", base.TestTransactionRequestPrototype, base.TestMessagePrototype,
		base.TestTransactionKeeperPrototype).InitializeKeeper(Mapper, parametersPrototype()).(transaction)
	require.Equal(t, "TestMessage", base.TestMessagePrototype().Route())
	require.NotNil(t, base.TestMessagePrototype().GetSignBytes())
	require.Equal(t, nil, base.TestTransactionKeeperPrototype().Transact(context, nil).GetError())

	// GetName
	require.Equal(t, "test", Transaction.GetName())

	// DecodeTransactionRequest
	message, Error := Transaction.DecodeTransactionRequest(json.RawMessage(`{"BaseReq":{"from":"addr"},"ID":"id"}`))
	require.Equal(t, nil, Error)
	require.Equal(t, sdkTypes.AccAddress("addr"), message.GetSigners()[0])

	// RegisterCodec : No Panics
	require.NotPanics(t, func() { Transaction.RegisterCodec(codec) })

	// Command : No Panics
	command := Transaction.Command(codec)
	require.Equal(t, `ABCIQuery: Post failed: Post "http://localhost:26657": dial tcp 127.0.0.1:26657: connect: connection refused`,
		command.ExecuteContext(context.Context()).Error())
	// HandleMessage
	_, Error = Transaction.HandleMessage(context, message)
	require.Nil(t, Error)

	// RESTRequestHandler : No Panics
	cliContext := clientContext.NewCLIContext().WithCodec(codec).WithChainID("test")

	//RPC ERROR
	request1 := codec.MustMarshalJSON(base.TransactionRequest{
		BaseReq: rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test"},
		ID:      "ID",
	})
	testRequest1, Error := http.NewRequest("GET", "/test", bytes.NewBuffer(request1))
	require.Nil(t, Error)
	responseRecorder := httptest.NewRecorder()
	Transaction.RESTRequestHandler(cliContext).ServeHTTP(responseRecorder, testRequest1)
	require.Equal(t, `{"error":"The specified item could not be found in the keyring"}`, responseRecorder.Body.String())
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	//invalid request
	request2 := codec.MustMarshalJSON(struct{}{})
	testRequest2, Error := http.NewRequest("GET", "/test", bytes.NewBuffer(request2))
	require.Nil(t, Error)
	responseRecorder = httptest.NewRecorder()
	Transaction.RESTRequestHandler(cliContext).ServeHTTP(responseRecorder, testRequest2)
	require.Equal(t, `{"error":"failed to decode JSON payload: JSON encoding of interfaces require non-empty type field."}`, responseRecorder.Body.String())
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	//validate fail
	request3 := codec.MustMarshalJSON(base.TransactionRequest{
		BaseReq: rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"},
		ID:      "ID",
	})
	testRequest3, Error := http.NewRequest("GET", "/test", bytes.NewBuffer(request3))
	require.Nil(t, Error)
	responseRecorder = httptest.NewRecorder()
	Transaction.RESTRequestHandler(cliContext).ServeHTTP(responseRecorder, testRequest3)
	require.Equal(t, `{"error":"chain-id required but not specified"}{"error":""}`, responseRecorder.Body.String())
	require.Equal(t, http.StatusUnauthorized, responseRecorder.Code)

	//Simulate RPC error
	request4 := codec.MustMarshalJSON(base.TransactionRequest{
		BaseReq: rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test", Simulate: true},
		ID:      "ID",
	})
	testRequest4, Error := http.NewRequest("GET", "/test", bytes.NewBuffer(request4))
	require.Nil(t, Error)
	responseRecorder = httptest.NewRecorder()
	Transaction.RESTRequestHandler(cliContext).ServeHTTP(responseRecorder, testRequest4)
	require.Equal(t, `{"error":"ABCIQuery: Post failed: Post \"http://localhost:26657\": dial tcp 127.0.0.1:26657: connect: connection refused"}`, responseRecorder.Body.String())
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	viper.Set(flags.FlagGenerateOnly, true)
	//Generate Only
	request5 := codec.MustMarshalJSON(base.TransactionRequest{
		BaseReq: rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test"},
		ID:      "ID",
	})
	testRequest5, Error := http.NewRequest("GET", "/test", bytes.NewBuffer(request5))
	require.Nil(t, Error)
	responseRecorder = httptest.NewRecorder()
	Transaction.RESTRequestHandler(cliContext).ServeHTTP(responseRecorder, testRequest5)
	require.Equal(t, http.StatusOK, responseRecorder.Code)

}
