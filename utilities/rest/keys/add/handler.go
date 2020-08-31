/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package add

import (
	"errors"
	"fmt"
	"github.com/bartekn/go-bip39"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"

	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func handler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var request request
		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.Codec, &request) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		Keyring, Error := cryptoKeys.NewKeyring(sdkTypes.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), strings.NewReader(keys.DefaultKeyPass))
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}
		info, Error := Keyring.Get(request.Name)
		if Error == nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, errors.New(fmt.Sprintf("Account for keyname %v already exists", request.Name)).Error())
			return
		}

		if request.Mnemonic != "" && !bip39.IsMnemonicValid(request.Mnemonic) {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, errors.New("invalid mnemonic").Error())
			return
		}
		if request.Mnemonic == "" {
			var mnemonicEntropySize = 256
			entropySeed, Error := bip39.NewEntropy(mnemonicEntropySize)
			if Error != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
				return
			}

			request.Mnemonic, Error = bip39.NewMnemonic(entropySeed)
			if Error != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
				return
			}
		}

		info, Error = Keyring.CreateAccount(request.Name, request.Mnemonic, cryptoKeys.DefaultBIP39Passphrase, keys.DefaultKeyPass, sdkTypes.FullFundraiserPath, cryptoKeys.Secp256k1)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}

		keyOutput, Error := cryptoKeys.Bech32KeyOutput(info)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusInternalServerError, Error.Error())
			return
		}
		keyOutput.Mnemonic = request.Mnemonic
		rest.PostProcessResponse(responseWriter, cliContext, newResponse(keyOutput, nil))
	}
}

func RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc("/keys/add", handler(cliContext)).Methods("POST")
}
