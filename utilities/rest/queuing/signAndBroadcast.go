// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
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

func signAndBroadcastMultiple(kafkaMsgList []kafkaMsg, context client.Context) ([]byte, error) {
	var stdTxs types.StdTx

	var txBytes []byte

	msgList := make([]sdkTypes.Msg, len(kafkaMsgList))
	for _, kafkaMsg := range kafkaMsgList {
		msgList = append(msgList, kafkaMsg.Msg)
	}

	for i, kafkaMsg := range kafkaMsgList {
		context := cliCtxFromKafkaMsg(kafkaMsg, context)

		gasAdj, err := parseGasAdjustment(kafkaMsg.BaseRequest.GasAdjustment)
		if err != nil {
			return nil, err
		}

		simAndExec, gas, err := flags.ParseGas(kafkaMsg.BaseRequest.Gas)
		if err != nil {
			return nil, err
		}

		keyBase, err := cryptoKeys.NewKeyring(sdkTypes.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), strings.NewReader(keys.DefaultKeyPass))
		if err != nil {
			return nil, err
		}

		accountNumber, sequence, err := types.NewAccountRetriever(context).GetAccountNumberSequence(context.FromAddress)
		if err != nil {
			return nil, err
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
			authClient.GetTxEncoder(context.Codec), accountNumber, sequence, gas, gasAdj,
			kafkaMsg.BaseRequest.Simulate, kafkaMsg.BaseRequest.ChainID, kafkaMsg.BaseRequest.Memo, kafkaMsg.BaseRequest.Fees, kafkaMsg.BaseRequest.GasPrices,
		)

		txBuilder = txBuilder.WithKeybase(keyBase)

		if kafkaMsg.BaseRequest.Simulate || simAndExec {
			if gasAdj < 0 {
				return nil, errors.New("Error invalid gas adjustment")
			}

			txBuilder, err = authClient.EnrichWithGas(txBuilder, context, []sdkTypes.Msg{kafkaMsg.Msg})
			if err != nil {
				return nil, err
			}

			if kafkaMsg.BaseRequest.Simulate {
				val, _ := simulationResponse(context.Codec, txBuilder.Gas())
				return val, nil
			}
		}

		stdMsg, err := txBuilder.BuildSignMsg(msgList)
		if err != nil {
			return nil, err
		}

		stdTx := auth.NewStdTx(stdMsg.Msgs, stdMsg.Fee, nil, stdMsg.Memo)

		stdTx, err = txBuilder.SignStdTx(context.FromName, keys.DefaultKeyPass, stdTx, true)
		if err != nil {
			return nil, err
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
			txBytes, err = txBuilder.TxEncoder()(stdTxs)
			if err != nil {
				return nil, err
			}
		}
	}

	response, err := cliCtxFromKafkaMsg(kafkaMsgList[0], context).BroadcastTx(txBytes)
	if err != nil {
		return nil, err
	}

	output, err := cliCtxFromKafkaMsg(kafkaMsgList[0], context).Codec.MarshalJSON(response)
	if err != nil {
		return nil, err
	}

	return output, nil
}
