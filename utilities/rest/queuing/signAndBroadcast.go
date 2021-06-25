/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	clientContext "github.com/cosmos/cosmos-sdk/client/context"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func signAndBroadcastMultiple(kafkaMsgList []kafkaMsg, cliContext clientContext.CLIContext) ([]byte, error) {
	var stdTxs types.StdTx

	var txBytes []byte

	var msgList []sdkTypes.Msg
	for _, kafkaMsg := range kafkaMsgList {
		msgList = append(msgList, kafkaMsg.Msg)
	}

	for i, kafkaMsg := range kafkaMsgList {
		msgCLIContext := cliCtxFromKafkaMsg(kafkaMsg, cliContext)

		gasAdj, Error := parseGasAdjustment(kafkaMsg.BaseRequest.GasAdjustment)
		if Error != nil {
			return nil, Error
		}

		simAndExec, gas, Error := flags.ParseGas(kafkaMsg.BaseRequest.Gas)
		if Error != nil {
			return nil, Error
		}

		keyBase, Error := cryptoKeys.NewKeyring(sdkTypes.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), strings.NewReader(keys.DefaultKeyPass))
		if Error != nil {
			return nil, Error
		}

		accountNumber, sequence, Error := types.NewAccountRetriever(cliContext).GetAccountNumberSequence(msgCLIContext.FromAddress)
		if Error != nil {
			return nil, Error
		}

		kafkaMsg.BaseRequest.AccountNumber = accountNumber

		var count = uint64(0)

		for j := 0; j < i; j++ {
			if accountNumber == kafkaMsgList[j].BaseRequest.AccountNumber {
				count++
			}
		}

		sequence += count
		txBuilder := types.NewTxBuilder(
			authClient.GetTxEncoder(cliContext.Codec), accountNumber, sequence, gas, gasAdj,
			kafkaMsg.BaseRequest.Simulate, kafkaMsg.BaseRequest.ChainID, kafkaMsg.BaseRequest.Memo, kafkaMsg.BaseRequest.Fees, kafkaMsg.BaseRequest.GasPrices,
		)

		txBuilder = txBuilder.WithKeybase(keyBase)

		if kafkaMsg.BaseRequest.Simulate || simAndExec {
			if gasAdj < 0 {
				return nil, errors.New("Error invalid gas adjustment")
			}

			txBuilder, Error = authClient.EnrichWithGas(txBuilder, cliContext, []sdkTypes.Msg{kafkaMsg.Msg})
			if Error != nil {
				return nil, Error
			}

			if kafkaMsg.BaseRequest.Simulate {
				val, _ := simulationResponse(cliContext.Codec, txBuilder.Gas())
				return val, nil
			}
		}

		stdMsg, Error := txBuilder.BuildSignMsg(msgList)
		if Error != nil {
			return nil, Error
		}

		stdTx := auth.NewStdTx(stdMsg.Msgs, stdMsg.Fee, nil, stdMsg.Memo)

		stdTx, Error = txBuilder.SignStdTx(msgCLIContext.FromName, keys.DefaultKeyPass, stdTx, true)
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

		if i == len(kafkaMsgList)-1 {
			txBytes, Error = txBuilder.TxEncoder()(stdTxs)
			if Error != nil {
				return nil, Error
			}
		}
	}

	response, Error := cliCtxFromKafkaMsg(kafkaMsgList[0], cliContext).BroadcastTx(txBytes)
	if Error != nil {
		return nil, Error
	}

	output, Error := cliCtxFromKafkaMsg(kafkaMsgList[0], cliContext).Codec.MarshalJSON(response)
	if Error != nil {
		return nil, Error
	}

	return output, nil
}
