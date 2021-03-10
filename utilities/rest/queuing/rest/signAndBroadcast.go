/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package rest

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/viper"
)

func SignAndBroadcastMultiple(brs []rest.BaseReq, cliContextList []client.Context, msgList []sdkTypes.Msg) ([]byte, error) {
	var stdTxs client.TxBuilder

	var txBytes []byte

	for i := range brs {
		gasAdj, _, Error := ParseFloat64OrReturnBadRequest(brs[i].GasAdjustment, flags.DefaultGasAdjustment)
		if Error != nil {
			return nil, Error
		}

		gasSetting, Error := flags.ParseGasSetting(brs[0].Gas)
		if Error != nil {
			return nil, Error
		}

		keyBase, Error := cryptoKeys.New(sdkTypes.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), strings.NewReader(keys.DefaultKeyPass))
		if Error != nil {
			return nil, Error
		}

		accountNumber, sequence, Error := cliContextList[i].AccountRetriever.GetAccountNumberSequence(cliContextList[i], cliContextList[i].FromAddress)
		if Error != nil {
			return nil, Error
		}
		brs[i].AccountNumber = accountNumber
		var count = uint64(0)
		for j := 0; j < i; j++ {
			if accountNumber == brs[j].AccountNumber {
				count++
				break
			}
		}
		if count > 0 {
			continue
		}

		txf := tx.Factory{}.
			WithAccountNumber(accountNumber).
			WithSequence(sequence).
			WithGas(gasSetting.Gas).
			WithGasAdjustment(gasAdj).
			WithMemo(brs[i].Memo).
			WithChainID(brs[i].ChainID).
			WithSimulateAndExecute(brs[i].Simulate).
			WithTxConfig(cliContextList[i].TxConfig).
			WithTimeoutHeight(brs[i].TimeoutHeight).
			WithFees(brs[i].Fees.String()).
			WithGasPrices(brs[i].GasPrices.String()).
			WithKeybase(keyBase)

		if brs[i].Simulate || gasSetting.Simulate {
			if gasAdj < 0 {
				return nil, errors.ErrorInvalidGasAdjustment
			}

			_, adjusted, err := tx.CalculateGas(cliContextList[i].QueryWithData, txf, msgList...)
			return nil, err

			txf = txf.WithGas(adjusted)

			if brs[i].Simulate {
				val, Error := SimulationResponse(cliContextList[i].LegacyAmino, txf.Gas())
				return val, Error
			}
		}

		txBuilder, err := tx.BuildUnsignedTx(txf, msgList...)
		if err != nil {
			return nil, err
		}
		err = tx.Sign(txf, cliContextList[i].FromName, txBuilder, false)
		if err != nil {
			return nil, err
		}

		if i == 0 {
			stdTxs.SetMsgs(txBuilder.GetTx().GetMsgs()...)
			stdTxs.SetGasLimit(txBuilder.GetTx().GetGas())
			stdTxs.SetFeeAmount(txBuilder.GetTx().GetFee())
			stdTxs.SetMemo(txBuilder.GetTx().GetMemo())
			stdTxs.SetTimeoutHeight(txBuilder.GetTx().GetTimeoutHeight())
		}
		signaturesV2, Error := txBuilder.GetTx().GetSignaturesV2()
		if err != nil {
			return nil, err
		}
		stdTxs.SetSignatures(signaturesV2...)

		if i == len(brs)-1 {
			txBytes, err = cliContextList[i].TxConfig.TxEncoder()(stdTxs.GetTx())
			if err != nil {
				return txBytes, err
			}
		}
	}

	response, Error := cliContextList[0].BroadcastTx(txBytes)
	if Error != nil {
		return nil, Error
	}

	responseBytes, Error := cliContextList[0].LegacyAmino.MarshalJSON(response)
	if Error != nil {
		return responseBytes, Error
	}

	wrappedResponse := rest.NewResponseWithHeight(cliContextList[0].Height, responseBytes)

	output, Error := cliContextList[0].LegacyAmino.MarshalJSON(wrappedResponse)
	if Error != nil {
		return output, Error
	}

	return output, nil
}
