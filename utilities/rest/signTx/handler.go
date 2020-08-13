/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package signTx

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strings"
)

func handler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var request request
		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.Codec, &request) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		Keyring, Error := keyring.New(sdkTypes.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), os.ExpandEnv("$HOME/.assetClient"), strings.NewReader(keys.DefaultKeyPass))
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "Key not found in keyring")
		}

		fromAddress, fromName, Error := context.GetFromFields(Keyring, request.BaseRequest.From, false)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}
		txBuilder := types.NewTxBuilder(
			authClient.GetTxEncoder(cliContext.Codec), request.BaseRequest.AccountNumber, request.BaseRequest.Sequence, 0, 0,
			request.BaseRequest.Simulate, request.BaseRequest.ChainID, request.BaseRequest.Memo, request.BaseRequest.Fees, request.BaseRequest.GasPrices,
		)
		//adding account sequence
		num, seq, Error := types.NewAccountRetriever(authClient.Codec, cliContext).GetAccountNumberSequence(fromAddress)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}
		txBuilder = txBuilder.WithAccountNumber(num)
		txBuilder = txBuilder.WithSequence(seq)

		signedStdTx, Error := txBuilder.SignStdTx(fromName, keys.DefaultKeyPass, request.StdTx, false)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}
		rest.PostProcessResponse(responseWriter, cliContext, newResponse(signedStdTx))
	}
}

func RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc("/signTx", handler(cliContext)).Methods("POST")
}
