/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package sign

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func handler(cliContext client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var request request
		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.LegacyAmino, &request) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		if request.BaseRequest.ChainID == "" {
			request.BaseRequest.ChainID = viper.GetString(flags.FlagChainID)
			if request.BaseRequest.ChainID == "" {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "Chain-ID required but not specified")
				return
			}
		}

		kr, Error := keyring.New(sdkTypes.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), strings.NewReader(keys.DefaultKeyPass))
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}
		fromAddress, fromName, _, Error := client.GetFromFields(kr, request.BaseRequest.From, viper.GetBool(flags.FlagGenerateOnly))
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}
		cliContext.FromAddress = fromAddress
		cliContext.FromName = fromName
		cliContext.From = request.BaseRequest.From

		txf := tx.Factory{}.
			WithAccountNumber(request.BaseRequest.AccountNumber).
			WithSequence(request.BaseRequest.Sequence).
			WithMemo(request.BaseRequest.Memo).
			WithChainID(request.BaseRequest.ChainID).
			WithSimulateAndExecute(request.BaseRequest.Simulate).
			WithTxConfig(cliContext.TxConfig).
			WithTimeoutHeight(request.BaseRequest.TimeoutHeight).
			WithFees(request.BaseRequest.Fees.String()).
			WithGasPrices(request.BaseRequest.GasPrices.String())

		txf, Error = tx.PrepareFactory(cliContext, txf)
		if rest.CheckBadRequestError(responseWriter, Error) {
			return
		}

		txBuilder, Error := tx.BuildUnsignedTx(txf, request.StdTx.Msgs...)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}
		Error = authClient.SignTx(txf, cliContext, fromName, txBuilder, false, false)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		signers := txBuilder.GetTx().GetSigners()
		pubicKeys := txBuilder.GetTx().GetPubKeys()
		if rest.CheckBadRequestError(responseWriter, Error) {
			return
		}

		if len(pubicKeys) > len(signers) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "cannot add more signatures than signers")
			return
		}

		for i, publicKey := range pubicKeys {
			if !bytes.Equal(publicKey.Address(), signers[i]) {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest,
					fmt.Sprintf("pubKey does not match signer address %s with signer index: %d", signers[i], i))
				return
			}
		}

		rest.PostProcessResponse(responseWriter, cliContext, newResponse(request.StdTx, nil))
	}
}

func RegisterRESTRoutes(cliContext client.Context, router *mux.Router) {
	router.HandleFunc("/sign", handler(cliContext)).Methods("POST")
}
