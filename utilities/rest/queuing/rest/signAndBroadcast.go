/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package rest

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func SignAndBroadcastMultiple(brs []rest.BaseReq, cliContextList []context.CLIContext, msgList []sdkTypes.Msg) ([]byte, error) {
	var stdTxs types.StdTx

	var txBytes []byte

	for i := range brs {
		gasAdj, _, Error := ParseFloat64OrReturnBadRequest(brs[i].GasAdjustment, flags.DefaultGasAdjustment)
		if Error != nil {
			return nil, Error
		}

		simAndExec, gas, Error := flags.ParseGas(brs[i].Gas)

		if Error != nil {
			return nil, Error
		}

		keyBase, Error := cryptoKeys.NewKeyring(sdkTypes.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), strings.NewReader(keys.DefaultKeyPass))
		if Error != nil {
			return nil, Error
		}

		accountNumber, sequence, Error := types.NewAccountRetriever(cliContextList[i]).GetAccountNumberSequence(cliContextList[i].FromAddress)
		if Error != nil {
			return nil, Error
		}

		brs[i].AccountNumber = accountNumber

		var count = uint64(0)

		for j := 0; j < i; j++ {
			if accountNumber == brs[j].AccountNumber {
				count++
			}
		}

		sequence += count
		txBuilder := types.NewTxBuilder(
			authClient.GetTxEncoder(cliContextList[i].Codec), accountNumber, sequence, gas, gasAdj,
			brs[i].Simulate, brs[i].ChainID, brs[i].Memo, brs[i].Fees, brs[i].GasPrices,
		)

		txBuilder = txBuilder.WithKeybase(keyBase)

		if brs[i].Simulate || simAndExec {
			if gasAdj < 0 {
				return nil, errors.New("Error invalid gas adjustment")
			}

			txBuilder, Error = authClient.EnrichWithGas(txBuilder, cliContextList[i], []sdkTypes.Msg{msgList[i]})
			if Error != nil {
				return nil, Error
			}

			if brs[i].Simulate {
				val, _ := SimulationResponse(cliContextList[i].Codec, txBuilder.Gas())
				return val, nil
			}
		}

		stdMsg, Error := txBuilder.BuildSignMsg(msgList)
		if Error != nil {
			return nil, Error
		}

		stdTx := auth.NewStdTx(stdMsg.Msgs, stdMsg.Fee, nil, stdMsg.Memo)

		stdTx, Error = txBuilder.SignStdTx(cliContextList[i].FromName, keys.DefaultKeyPass, stdTx, true)
		if Error != nil {
			return nil, Error
		}

		if i == 0 {
			stdTxs.Msgs = stdTx.Msgs
			stdTxs.Fee = stdTx.Fee
			stdTxs.Memo = stdTx.Memo
		}

		if count == 0 {
			stdTxs.Signatures = append(stdTxs.Signatures, stdTx.Signatures...)
		}

		if i == len(brs)-1 {
			txBytes, Error = txBuilder.TxEncoder()(stdTxs)
			if Error != nil {
				return nil, Error
			}
		}
	}

	response, Error := cliContextList[0].BroadcastTx(txBytes)
	if Error != nil {
		return nil, Error
	}

	output, Error := cliContextList[0].Codec.MarshalJSON(response)
	if Error != nil {
		return nil, Error
	}

	return output, nil
}
