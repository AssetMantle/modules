/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func SignAndBroadcastMultiples(brs []rest.BaseReq, cliContextList []context.CLIContext, msgList []sdkTypes.Msg) ([]byte, error) {
	var stdTxs types.StdTx
	for i := range brs {
		cdc := cliContextList[i].Codec
		gasAdj, _, err := ParseFloat64OrReturnBadRequest(brs[i].GasAdjustment, flags.DefaultGasAdjustment)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		simAndExec, gas, err := flags.ParseGas(brs[i].Gas)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		keyBase, err := cryptoKeys.NewKeyring(sdkTypes.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), os.ExpandEnv("$HOME/.assetClient"), strings.NewReader(keys.DefaultKeyPass))
		if err != nil {
			panic(fmt.Errorf("couldn't acquire keyring: %v", err))
		}

		address, err := sdkTypes.AccAddressFromBech32(brs[i].From)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		//adding account sequence
		num, seq, err := types.NewAccountRetriever(cliContextList[i]).GetAccountNumberSequence(address)
		if err != nil {
			fmt.Printf("Error in NewAccountRetriever: %s\n", err)
			return nil, nil
		}

		txBuilder := types.NewTxBuilder(
			authClient.GetTxEncoder(cliContextList[i].Codec), brs[i].AccountNumber, brs[i].Sequence, gas, gasAdj,
			brs[i].Simulate, brs[i].ChainID, brs[i].Memo, brs[i].Fees, brs[i].GasPrices,
		)

		txBuilder = txBuilder.WithKeybase(keyBase)

		if brs[i].Simulate || simAndExec {
			if gasAdj < 0 {
				return nil, errors.New("Error invalid gas adjustment")
			}

			txBuilder, err = authClient.EnrichWithGas(txBuilder, cliContextList[i], []sdkTypes.Msg{msgList[i]})
			if err != nil {
				return nil, errors.New(err.Error())
			}

			if brs[i].Simulate {
				val, _ := SimulationResponse(cdc, txBuilder.Gas())
				return val, nil
			}
		}

		txBuilder = txBuilder.WithAccountNumber(num)
		txBuilder = txBuilder.WithSequence(seq)
		fromName := cliContextList[i].GetFromName()

		//build and sign
		stdMsg, err := txBuilder.BuildAndSign(fromName, keys.DefaultKeyPass, msgList)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		stdSignMsg, err := txBuilder.BuildSignMsg(msgList)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		var count = uint64(0)
		for j := 0; j < i; j++ {
			if txBuilder.AccountNumber() == brs[j].AccountNumber {
				count++
			}
		}

		if i == 0 {
			stdTxs.Msgs = stdSignMsg.Msgs
			stdTxs.Fee = stdSignMsg.Fee
			stdTxs.Memo = stdSignMsg.Memo
		}

		// broadcast to a node
		res, err := cliContextList[i].BroadcastTx(stdMsg)
		if err != nil {
			fmt.Printf("Error in broadcast: %s\n", err)
			return nil, nil
		}

		output, err := cliContextList[i].Codec.MarshalJSON(res)

		fmt.Printf("output: %s\n", output)
		return output, nil

	}
	return nil, nil
}
