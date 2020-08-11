/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package rest

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client"
)

func SignStdTxFromRest(txBuilder auth.TxBuilder, cliCtx context.CLIContext, name string, stdTx auth.StdTx, appendSig bool, offline bool, password string) (auth.StdTx, error) {

	var signedStdTx auth.StdTx

	info, err := txBuilder.Keybase().Key(name)
	if err != nil {
		return signedStdTx, err
	}

	addr := info.GetPubKey().Address()

	if !isTxSigner(sdkTypes.AccAddress(addr), stdTx.GetSigners()) {
		return signedStdTx, fmt.Errorf("%s: %s", "Error invalid signer", name)
	}

	ad := sdkTypes.AccAddress(addr)
	if !offline {
		txBuilder, err = populateAccountFromState(txBuilder, cliCtx, ad)
		if err != nil {
			return signedStdTx, err
		}
	}

	return txBuilder.SignStdTx(name, password, stdTx, appendSig)
}

func isTxSigner(user sdkTypes.AccAddress, signers []sdkTypes.AccAddress) bool {
	for _, s := range signers {
		if bytes.Equal(user.Bytes(), s.Bytes()) {
			return true
		}
	}

	return false
}

func populateAccountFromState(
	txBuilder auth.TxBuilder, cliCtx context.CLIContext, addr sdkTypes.AccAddress,
) (auth.TxBuilder, error) {
	num, seq, err := auth.NewAccountRetriever(authClient.Codec, cliCtx).GetAccountNumberSequence(addr)
	if err != nil {
		return txBuilder, err
	}

	return txBuilder.WithAccountNumber(num).WithSequence(seq + txBuilder.Sequence()), nil
}
