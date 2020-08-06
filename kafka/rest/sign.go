/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package rest

import (
	"bytes"
	"fmt"
	_ "github.com/cosmos/cosmos-sdk/client/keys"
	_ "github.com/cosmos/cosmos-sdk/crypto/keyring"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client"
	_ "strings"

	context "github.com/cosmos/cosmos-sdk/client/context"
	//"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cTypes "github.com/cosmos/cosmos-sdk/types"
	//sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
)

func SignStdTxFromRest(txBldr auth.TxBuilder, cliCtx context.CLIContext, name string, stdTx auth.StdTx, appendSig bool, offline bool, password string) (auth.StdTx, error) {

	var signedStdTx auth.StdTx

	info, err := txBldr.Keybase().Key(name)
	if err != nil {
		return signedStdTx, err
	}

	addr := info.GetPubKey().Address()

	if !isTxSigner(cTypes.AccAddress(addr), stdTx.GetSigners()) {
		return signedStdTx, fmt.Errorf("%s: %s", "Error invalid signer", name)
	}

	ad := cTypes.AccAddress(addr)
	if !offline {
		txBldr, err = populateAccountFromState(txBldr, cliCtx, ad)
		if err != nil {
			return signedStdTx, err
		}
	}

	return txBldr.SignStdTx(name, password, stdTx, appendSig)
}

func isTxSigner(user cTypes.AccAddress, signers []cTypes.AccAddress) bool {
	for _, s := range signers {
		if bytes.Equal(user.Bytes(), s.Bytes()) {
			return true
		}
	}

	return false
}

func populateAccountFromState(
	txBldr auth.TxBuilder, cliCtx context.CLIContext, addr cTypes.AccAddress,
) (auth.TxBuilder, error) {
	num, seq, err := auth.NewAccountRetriever(authClient.Codec, cliCtx).GetAccountNumberSequence(addr)
	if err != nil {
		return txBldr, err
	}

	return txBldr.WithAccountNumber(num).WithSequence(seq + txBldr.Sequence()), nil
}
