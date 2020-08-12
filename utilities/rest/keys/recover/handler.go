/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package recover

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
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
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}

		info, Error := Keyring.NewAccount(request.Name, request.Mnemonic, keyring.DefaultBIP39Passphrase, sdkTypes.FullFundraiserPath, hd.Secp256k1)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}

		keyOutput, Error := keyring.Bech32KeyOutput(info)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}

		rest.PostProcessResponse(responseWriter, cliContext, newResponse(keyOutput))
	}
}

func RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc("/keys/recover", handler(cliContext)).Methods("POST")
}
