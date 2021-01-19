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

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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

		fromAddress, fromName, Error := context.GetFromFields(strings.NewReader(keys.DefaultKeyPass), request.BaseRequest.From, false)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		txBuilder := types.NewTxBuilder(
			authClient.GetTxEncoder(cliContext.Codec), request.BaseRequest.AccountNumber, request.BaseRequest.Sequence, 0, 0,
			request.BaseRequest.Simulate, request.BaseRequest.ChainID, request.BaseRequest.Memo, request.BaseRequest.Fees, request.BaseRequest.GasPrices,
		)

		accountNumber, sequence, Error := types.NewAccountRetriever(cliContext).GetAccountNumberSequence(fromAddress)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		txBuilder = txBuilder.WithAccountNumber(accountNumber)
		txBuilder = txBuilder.WithSequence(sequence)

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

		rest.PostProcessResponse(responseWriter, cliContext, newResponse(request.StdTx, nil))
	}
}

func RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc("/sign", handler(cliContext)).Methods("POST")
}
