/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package add

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	Codec := codec.New()
	schema.RegisterCodec(Codec)
	Codec.RegisterConcrete(request{}, "request", nil)
	Codec.RegisterConcrete(response{}, "response", nil)
	clientContext := context.NewCLIContext().WithCodec(Codec)
	handler := handler(clientContext)
	viper.Set(flags.FlagKeyringBackend, keys.BackendTest)
	viper.Set(flags.FlagHome, t.TempDir())

	keyring, Error := cryptoKeys.NewKeyring(sdk.KeyringServiceName(), keys.BackendTest, t.TempDir(), strings.NewReader(""))
	require.NoError(t, Error)

	router := mux.NewRouter()
	RegisterRESTRoutes(clientContext, router)

	t.Cleanup(func() {
		_ = keyring.Delete("keyName1", "", true)
		_ = keyring.Delete("keyName2", "", true)
		_ = keyring.Delete("keyName3", "", true)
	})

	getResponse := func(responseBytes []byte) response {
		var Response rest.ResponseWithHeight
		Error := Codec.UnmarshalJSON(responseBytes, &Response)
		require.Nil(t, Error)

		var ResponseValue response
		Error = Codec.UnmarshalJSON(Response.Result, &ResponseValue)
		return ResponseValue
	}

	// create account without mnemonic
	requestBody1, Error := Codec.MarshalJSON(request{
		Name: "testKey1",
	})
	require.Nil(t, Error)
	testRequest1, Error := http.NewRequest("POST", "/keys/add", bytes.NewBuffer(requestBody1))
	require.Nil(t, Error)
	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest1)
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	response1 := getResponse(responseRecorder.Body.Bytes())
	require.Nil(t, response1.Error)
	require.Equal(t, true, response1.Success)

	// create account with mnemonic
	requestBody2, Error := Codec.MarshalJSON(request{
		Name:     "testKey2",
		Mnemonic: "wage thunder live sense resemble foil apple course spin horse glass mansion midnight laundry acoustic rhythm loan scale talent push green direct brick please",
	})
	require.Nil(t, Error)
	testRequest2, Error := http.NewRequest("POST", "/keys/add", bytes.NewBuffer(requestBody2))
	require.Nil(t, Error)

	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest2)
	require.Equal(t, responseRecorder.Code, http.StatusOK)

	response2 := getResponse(responseRecorder.Body.Bytes())
	require.Nil(t, response2.Error)
	require.Equal(t, true, response2.Success)

	// invalid mnemonic
	requestBody3, Error := Codec.MarshalJSON(request{
		Name:     "testKey3",
		Mnemonic: "wage brick please",
	})
	require.Nil(t, Error)
	testRequest3, Error := http.NewRequest("POST", "/keys/add", bytes.NewBuffer(requestBody3))
	require.Nil(t, Error)

	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest3)
	require.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
	require.Equal(t, `{"error":"invalid mnemonic"}`, responseRecorder.Body.String())

	// invalid request
	testRequest4, Error := http.NewRequest("POST", "/keys/add", bytes.NewBuffer([]byte{}))
	require.Nil(t, Error)

	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest4)
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	//Retry adding same account
	requestBody5, Error := Codec.MarshalJSON(request{
		Name: "testKey1",
	})
	require.Nil(t, Error)
	testRequest5, Error := http.NewRequest("POST", "/keys/add", bytes.NewBuffer(requestBody5))
	require.Nil(t, Error)

	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest5)
	require.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
	require.Equal(t, `{"error":"Account for keyname testKey1 already exists"}`, responseRecorder.Body.String())

}
