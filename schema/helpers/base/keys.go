package base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type NewKeyBody struct {
	Name string `json:"name"`
}

type NewKeyRecoverBody struct {
	Name     string `json:"name"`
	Mnemonic string `json:"mnemonic"`
}

func RESTKeysHandler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var newKey NewKeyBody

		body, err := ioutil.ReadAll(httpRequest.Body)

		if err = json.Unmarshal(body, &newKey); rest.CheckBadRequestError(responseWriter, err) {
			return
		}

		info, mnemonic, Error := keyring.NewInMemory().NewMnemonic(newKey.Name, keyring.English, sdk.FullFundraiserPath, hd.Secp256k1)
		if Error != nil {
			fmt.Printf("Error in generating mnemonic: %s\n", err)
			return
		}

		keyOutput, Error := cryptoKeys.Bech32KeyOutput(info)
		if Error != nil {
			fmt.Printf("Error in generating mnemonic: %s\n", err)
			return
		}

		keyOutput.Mnemonic = mnemonic
		bz, Error := json.Marshal(keyOutput)
		if Error != nil {
			fmt.Printf("Error in generating mnemonic: %s\n", err)
			return
		}

		responseWriter.Header().Set("Content-Type", "application/json")
		if _, err := responseWriter.Write(bz); err != nil {
			log.Printf("could not write response: %v", err)
		}
	}
}

func RESTKeysRecoverHandler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var newKeyRecover NewKeyRecoverBody

		body, err := ioutil.ReadAll(httpRequest.Body)

		if err = json.Unmarshal(body, &newKeyRecover); rest.CheckBadRequestError(responseWriter, err) {
			return
		}

		seed := newKeyRecover.Mnemonic
		from := newKeyRecover.Name

		info, err := keyring.NewInMemory().NewAccount(from, seed, keyring.DefaultBIP39Passphrase, sdk.FullFundraiserPath, hd.Secp256k1)
		if err != nil {
			return
		}

		keyOutput, err := cryptoKeys.Bech32KeyOutput(info)
		if err != nil {
			return
		}

		keyOutput.Mnemonic = seed
		bz, err := json.Marshal(keyOutput)

		responseWriter.Header().Set("Content-Type", "application/json")
		if _, err := responseWriter.Write(bz); err != nil {
			log.Printf("could not write response: %v", err)
		}
	}
}
