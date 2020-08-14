/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package signTx

import (
	"bytes"
	"fmt"
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

		if request.BaseRequest.ChainID == "" {
			request.BaseRequest.ChainID = viper.GetString(flags.FlagChainID)
			if request.BaseRequest.ChainID == "" {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "Chain-ID required but not specified")
				return
			}
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

		num, seq, Error := types.NewAccountRetriever(authClient.Codec, cliContext).GetAccountNumberSequence(fromAddress)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}
		txBuilder = txBuilder.WithAccountNumber(num)
		txBuilder = txBuilder.WithSequence(seq)

		stdSignature, Error := types.MakeSignature(txBuilder.Keybase(), fromName, keys.DefaultKeyPass, types.StdSignMsg{
			ChainID:       txBuilder.ChainID(),
			AccountNumber: txBuilder.AccountNumber(),
			Sequence:      txBuilder.Sequence(),
			Fee:           request.StdTx.Fee,
			Msgs:          request.StdTx.GetMsgs(),
			Memo:          request.StdTx.GetMemo(),
		})
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		signers := request.StdTx.GetSigners()
		request.StdTx.Signatures = append(request.StdTx.Signatures, stdSignature)
		pubicKeys := request.StdTx.GetPubKeys()
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

		rest.PostProcessResponse(responseWriter, cliContext, newResponse(request.StdTx))
	}
}

func RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc("/signTx", handler(cliContext)).Methods("POST")
}
