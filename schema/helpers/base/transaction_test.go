// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/utilities/test"
	"github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

func TestTransaction(t *testing.T) {
	context, storeKey, _ := test.SetupTest(t)
	var legacyAmino = sdkCodec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Transaction := NewTransaction("test", "", "", base.TestTransactionRequestPrototype, base.TestMessagePrototype,
		base.TestTransactionKeeperPrototype, nil, nil).InitializeKeeper(Mapper, parameterManagerPrototype()).(transaction)
	require.Equal(t, "TestMessage", base.TestMessagePrototype().(*base.TestMessage).Route())
	require.NotNil(t, base.TestMessagePrototype().(*base.TestMessage).GetSignBytes())
	_, err := base.TestTransactionKeeperPrototype().Transact(context, nil)
	require.Equal(t, nil, err)

	// GetName
	require.Equal(t, "test", Transaction.GetName())

	// DecodeTransactionRequest
	message, err := Transaction.DecodeTransactionRequest(json.RawMessage(`{"BaseReq":{"from":"addr"},"ID":"id"}`))
	require.Equal(t, nil, err)
	require.Equal(t, sdkTypes.AccAddress("addr"), message.GetSigners()[0])

	// RegisterLegacyAminoCodec : No Panics
	require.NotPanics(t, func() { Transaction.RegisterLegacyAminoCodec(legacyAmino) })

	// Command : No Panics
	command := Transaction.Command()
	err = command.ParseFlags([]string{"--node", "tcp://localhost:26657"})
	require.Nil(t, err)
	require.Equal(t, `ABCIQuery: Post failed: Post "http://localhost:26657": dial tcp 127.0.0.1:26657: connect: connection refused`,
		command.ExecuteContext(context).Error())
	// HandleQuery
	_, err = Transaction.HandleMessage(context, message.(helpers.Message))
	require.Nil(t, err)

	// RPC ERROR
	request1 := legacyAmino.MustMarshalJSON(base.TransactionRequest{
		BaseReq: rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test"},
		ID:      "ID",
	})
	testRequest1, err := http.NewRequest("GET", "/test", bytes.NewBuffer(request1))
	require.Nil(t, err)
	responseRecorder := httptest.NewRecorder()
	Transaction.RESTRequestHandler(TestClientContext).ServeHTTP(responseRecorder, testRequest1)
	require.Equal(t, `{"error":"ABCIQuery: Post failed: Post \"http://localhost:26657\": dial tcp 127.0.0.1:26657: connect: connection refused"}`, responseRecorder.Body.String())
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// invalid request
	request2 := legacyAmino.MustMarshalJSON(struct{}{})
	testRequest2, err := http.NewRequest("GET", "/test", bytes.NewBuffer(request2))
	require.Nil(t, err)
	responseRecorder = httptest.NewRecorder()
	Transaction.RESTRequestHandler(TestClientContext).ServeHTTP(responseRecorder, testRequest2)
	require.Equal(t, `{"error":"failed to decode JSON payload: JSON encoding of interfaces require non-empty type field."}`, responseRecorder.Body.String())
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// validate fail
	request3 := legacyAmino.MustMarshalJSON(base.TransactionRequest{
		BaseReq: rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"},
		ID:      "ID",
	})
	testRequest3, err := http.NewRequest("GET", "/test", bytes.NewBuffer(request3))
	require.Nil(t, err)
	responseRecorder = httptest.NewRecorder()
	Transaction.RESTRequestHandler(TestClientContext).ServeHTTP(responseRecorder, testRequest3)
	require.Equal(t, `{"error":"chain-id required but not specified"}{"error":""}`, responseRecorder.Body.String())
	require.Equal(t, http.StatusUnauthorized, responseRecorder.Code)

	// Simulate RPC error
	request4 := legacyAmino.MustMarshalJSON(base.TransactionRequest{
		BaseReq: rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test", Simulate: true},
		ID:      "ID",
	})
	testRequest4, err := http.NewRequest("GET", "/test", bytes.NewBuffer(request4))
	require.Nil(t, err)
	responseRecorder = httptest.NewRecorder()
	Transaction.RESTRequestHandler(TestClientContext).ServeHTTP(responseRecorder, testRequest4)
	require.Equal(t, `{"error":"ABCIQuery: Post failed: Post \"http://localhost:26657\": dial tcp 127.0.0.1:26657: connect: connection refused"}`, responseRecorder.Body.String())
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	viper.Set(flags.FlagGenerateOnly, true)
	// Generate Only
	request5 := legacyAmino.MustMarshalJSON(base.TransactionRequest{
		BaseReq: rest.BaseReq{From: "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", ChainID: "test"},
		ID:      "ID",
	})
	testRequest5, err := http.NewRequest("GET", "/test", bytes.NewBuffer(request5))
	require.Nil(t, err)
	responseRecorder = httptest.NewRecorder()
	Transaction.RESTRequestHandler(TestClientContext).ServeHTTP(responseRecorder, testRequest5)
	require.Equal(t, http.StatusOK, responseRecorder.Code)

}
