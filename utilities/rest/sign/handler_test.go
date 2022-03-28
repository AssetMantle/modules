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

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
)

func TestHandler(t *testing.T) {
	Codec := codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	Codec.RegisterConcrete(request{}, "request", nil)
	Codec.RegisterConcrete(response{}, "response", nil)
	base.TestMessagePrototype().RegisterCodec(Codec)

	clientContext := context.NewCLIContext().WithCodec(Codec)

	handler := handler(clientContext)
	viper.Set(flags.FlagKeyringBackend, cryptoKeys.BackendTest)

	keyring, err := cryptoKeys.NewKeyring(sdk.KeyringServiceName(), cryptoKeys.BackendTest, viper.GetString(flags.FlagHome), strings.NewReader(""))
	require.NoError(t, err)

	router := mux.NewRouter()
	RegisterRESTRoutes(clientContext, router)

	t.Cleanup(func() {
		_ = keyring.Delete("keyName1", "", true)
		_ = keyring.Delete("keyName2", "", true)
		_ = keyring.Delete("keyName3", "", true)
	})
	_, err = keyring.CreateAccount("keyName1", "wage thunder live sense resemble foil apple course spin horse glass mansion midnight laundry acoustic rhythm loan scale talent push green direct brick please",
		cryptoKeys.DefaultBIP39Passphrase, keys.DefaultKeyPass, sdkTypes.FullFundraiserPath, cryptoKeys.Secp256k1)
	require.Nil(t, err)

	address := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"

	var sdkAddress sdkTypes.AccAddress
	sdkAddress, err = sdkTypes.AccAddressFromBech32(address)

	// signWithout chainID
	requestBody1, err := Codec.MarshalJSON(request{
		BaseRequest: rest.BaseReq{From: address},
		Type:        "cosmos-sdk/StdTx",
		StdTx:       auth.NewStdTx([]sdkTypes.Msg{base.NewTestMessage(sdkAddress, "id")}, auth.NewStdFee(10, sdkTypes.NewCoins()), nil, ""),
	})
	require.Nil(t, err)
	testRequest1, err := http.NewRequest("POST", "/sign", bytes.NewBuffer(requestBody1))
	require.Nil(t, err)
	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest1)
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	require.Equal(t, `{"error":"Chain-ID required but not specified"}`, responseRecorder.Body.String())

	// with wrong key
	requestBody2, err := Codec.MarshalJSON(request{
		BaseRequest: rest.BaseReq{From: "address", ChainID: "test"},
		Type:        "cosmos-sdk/StdTx",
		StdTx:       auth.NewStdTx([]sdkTypes.Msg{base.NewTestMessage(sdkAddress, "id")}, auth.NewStdFee(20, sdkTypes.NewCoins()), nil, ""),
	})
	require.Nil(t, err)
	testRequest2, err := http.NewRequest("POST", "/sign", bytes.NewBuffer(requestBody2))
	require.Nil(t, err)
	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest2)
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	require.Equal(t, `{"error":"The specified item could not be found in the keyring"}`, responseRecorder.Body.String())

	// RPC client offline
	requestBody3, err := Codec.MarshalJSON(request{
		BaseRequest: rest.BaseReq{From: address, ChainID: "test"},
		Type:        "cosmos-sdk/StdTx",
		StdTx:       auth.NewStdTx([]sdkTypes.Msg{base.NewTestMessage(sdkAddress, "id")}, auth.NewStdFee(30, sdkTypes.NewCoins()), nil, ""),
	})
	require.Nil(t, err)
	testRequest3, err := http.NewRequest("POST", "/sign", bytes.NewBuffer(requestBody3))
	require.Nil(t, err)
	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, testRequest3)
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	require.Equal(t, `{"error":"no RPC client is defined in offline mode"}`, responseRecorder.Body.String())

}
