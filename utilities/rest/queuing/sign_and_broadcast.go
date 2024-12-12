// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func signAndBroadcastMultiple(kafkaMsgList []kafkaMsg, context client.Context) ([]byte, error) {
	var txBytes []byte

	msgList := make([]sdkTypes.Msg, len(kafkaMsgList))
	for _, kafkaMsg := range kafkaMsgList {
		msgList = append(msgList, kafkaMsg.Msg)
	}

	for i, kafkaMsg := range kafkaMsgList {
		context := cliCtxFromKafkaMsg(kafkaMsg, context)

		gasAdj, err := parseGasAdjustment(kafkaMsg.CommonTransactionRequest.GetGasAdjustment())
		if err != nil {
			return nil, err
		}

		gasSetting, err := flags.ParseGasSetting(kafkaMsg.CommonTransactionRequest.GetGas())
		if err != nil {
			return nil, err
		}

		keyBase, err := keyring.New(sdkTypes.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), strings.NewReader(keys.DefaultKeyPass), context.Codec)
		if err != nil {
			return nil, err
		}

		accountNumber, sequence, err := types.AccountRetriever{}.GetAccountNumberSequence(context, context.FromAddress)
		if err != nil {
			return nil, err
		}

		kafkaMsg.CommonTransactionRequest = kafkaMsg.CommonTransactionRequest.SetAccountNumber(accountNumber)

		var count = uint64(0)

		for j := 0; j < i; j++ {
			if accountNumber == kafkaMsgList[j].CommonTransactionRequest.GetAccountNumber() {
				count++
			}
		}

		sequence += count

		txFactory := tx.Factory{}.
			WithAccountNumber(kafkaMsg.CommonTransactionRequest.GetAccountNumber()).
			WithSequence(kafkaMsg.CommonTransactionRequest.GetSequence()).
			WithGas(gasSetting.Gas).
			WithGasAdjustment(gasAdj).
			WithMemo(kafkaMsg.CommonTransactionRequest.GetMemo()).
			WithChainID(kafkaMsg.CommonTransactionRequest.GetChainID()).
			WithSimulateAndExecute(kafkaMsg.CommonTransactionRequest.IsSimulated()).
			WithTxConfig(context.TxConfig).
			WithTimeoutHeight(kafkaMsg.CommonTransactionRequest.GetTimeoutHeight()).
			WithFees(kafkaMsg.CommonTransactionRequest.GetFees().String()).
			WithGasPrices(kafkaMsg.CommonTransactionRequest.GetGasPrices().String()).
			WithKeybase(keyBase)

		if kafkaMsg.CommonTransactionRequest.IsSimulated() || gasSetting.Simulate {
			if gasAdj < 0 {
				return nil, errors.New("Error invalid gas adjustment")
			}

			_, adjusted, err := tx.CalculateGas(context, txFactory, kafkaMsg.Msg)

			if err != nil {
				return nil, err
			}

			txFactory = txFactory.WithGas(adjusted)

			if kafkaMsg.CommonTransactionRequest.IsSimulated() {
				val, _ := simulationResponse(context.LegacyAmino, txFactory.Gas())
				return val, nil
			}
		}

		txBuilder, err := txFactory.BuildUnsignedTx(msgList...)
		if err != nil {
			return nil, err
		}

		err = tx.Sign(txFactory, context.FromName, txBuilder, true)
		if err != nil {
			return nil, err
		}

		txBytes, err = context.TxConfig.TxEncoder()(txBuilder.GetTx())

		if err != nil {
			return nil, err
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
