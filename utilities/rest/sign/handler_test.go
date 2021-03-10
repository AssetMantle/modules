/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package sign

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/cosmos/cosmos-sdk/crypto/hd"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	Codec := codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	Codec.RegisterConcrete(request{}, "request", nil)
	Codec.RegisterConcrete(response{}, "response", nil)
	base.TestMessagePrototype().RegisterCodec(Codec)

	clientContext := client.Context{}.WithLegacyAmino(Codec)
	handler := handler(clientContext)
	viper.Set(flags.FlagKeyringBackend, cryptoKeys.BackendTest)
	viper.Set(flags.FlagHome, t.TempDir())

	keyring, Error := cryptoKeys.New(sdk.KeyringServiceName(), cryptoKeys.BackendTest, viper.GetString(flags.FlagHome), strings.NewReader(""))
	require.NoError(t, Error)

	router := mux.NewRouter()
	RegisterRESTRoutes(clientContext, router)

	t.Cleanup(func() {
		_ = keyring.Delete("keyName1")
		_ = keyring.Delete("keyName2")
		_ = keyring.Delete("keyName3")
	})
	_, Error = keyring.NewAccount("keyName1", "wage thunder live sense resemble foil apple course spin horse glass mansion midnight laundry acoustic rhythm loan scale talent push green direct brick please",
		cryptoKeys.DefaultBIP39Passphrase, sdkTypes.FullFundraiserPath, hd.Secp256k1)
	require.Nil(t, Error)

	address := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	sdkAddress, Error := sdkTypes.AccAddressFromBech32(address)

	// signWithout chainID
	requestBody1, Error := Codec.MarshalJSON(request{
		BaseRequest: rest.BaseReq{From: address},
		Type:        "cosmos-sdk/StdTx",
		StdTx:       legacytx.NewStdTx([]sdkTypes.Msg{base.NewTestMessage(sdkAddress, "id")}, legacytx.NewStdFee(10, sdkTypes.NewCoins()), nil, ""),
	})
	require.Nil(t, Error)
	testRequest1, Error := http.NewRequest("POST", "/sign", bytes.NewBuffer(requestBody1))
	require.Nil(t, Error)
	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest1)
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	require.Equal(t, `{"error":"Chain-ID required but not specified"}`, responseRecorder.Body.String())

	// with wrong key
	requestBody2, Error := Codec.MarshalJSON(request{
		BaseRequest: rest.BaseReq{From: "address", ChainID: "test"},
		Type:        "cosmos-sdk/StdTx",
		StdTx:       legacytx.NewStdTx([]sdkTypes.Msg{base.NewTestMessage(sdkAddress, "id")}, legacytx.NewStdFee(20, sdkTypes.NewCoins()), nil, ""),
	})
	require.Nil(t, Error)
	testRequest2, Error := http.NewRequest("POST", "/sign", bytes.NewBuffer(requestBody2))
	require.Nil(t, Error)
	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest2)
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	require.Equal(t, `{"error":"The specified item could not be found in the keyring"}`, responseRecorder.Body.String())

	// RPC client offline
	requestBody3, Error := Codec.MarshalJSON(request{
		BaseRequest: rest.BaseReq{From: address, ChainID: "test"},
		Type:        "cosmos-sdk/StdTx",
		StdTx:       legacytx.NewStdTx([]sdkTypes.Msg{base.NewTestMessage(sdkAddress, "id")}, legacytx.NewStdFee(30, sdkTypes.NewCoins()), nil, ""),
	})
	require.Nil(t, Error)
	testRequest3, Error := http.NewRequest("POST", "/sign", bytes.NewBuffer(requestBody3))
	require.Nil(t, Error)
	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest3)
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	require.Equal(t, `{"error":"no RPC client is defined in offline mode"}`, responseRecorder.Body.String())

}
