package rest

import (
	"fmt"
	_ "github.com/Workiva/go-datastructures/threadsafe/err"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	//"log"
	//"net/http"

	//"log"
	//"net/http"
	"strings"

	//"net/http"

	//"github.com/cosmos/cosmos-sdk/client"
	context "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	_ "github.com/cosmos/cosmos-sdk/types"
	//"github.com/cosmos/cosmos-sdk/client/key"
	cTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	authClient "github.com/cosmos/cosmos-sdk/x/auth/client"
	tx "github.com/cosmos/cosmos-sdk/x/auth/client"
	//"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
)

func SignAndBroadcast(br rest.BaseReq, cliCtx context.CLIContext,
	mode, password string, msgs []cTypes.Msg) ([]byte, error) {

	cdc := cliCtx.Codec
	gasAdj, _, err := ParseFloat64OrReturnBadRequest(br.GasAdjustment, flags.DefaultGasAdjustment)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	simAndExec, gas, err := flags.ParseGas(br.Gas)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	keyBase, err := keyring.New(sdkTypes.KeyringServiceName(), "os", "home", strings.NewReader(keys.DefaultKeyPass))
	if err != nil {
		panic(fmt.Errorf("couldn't acquire keyring: %v", err))
	}

	txBldr := types.NewTxBuilder(
		authClient.GetTxEncoder(cliCtx.Codec), br.AccountNumber, br.Sequence, gas, gasAdj,
		br.Simulate, br.ChainID, br.Memo, br.Fees, br.GasPrices,
	)

	txBldr = txBldr.WithKeybase(keyBase)

	if br.Simulate || simAndExec {
		if gasAdj < 0 {
			return nil, errors.New("Error invalid gas adjustment")
		}

		txBldr, err = tx.EnrichWithGas(txBldr, cliCtx, msgs)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		if br.Simulate {
			return SimulationResponse(cdc, txBldr.Gas())
		}
	}

	stdMsg, err := txBldr.BuildSignMsg(msgs)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	stdTx := auth.NewStdTx(stdMsg.Msgs, stdMsg.Fee, nil, stdMsg.Memo)

	stdTx, err = SignStdTxFromRest(txBldr, cliCtx, cliCtx.GetFromName(), stdTx, true, false, password)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return BroadcastRest(cliCtx, cdc, stdTx, mode)

}

func SignAndBroadcastMultiples(brs []rest.BaseReq, cliCtxs []context.CLIContext, msgs []cTypes.Msg) ([]byte, error) {
	var stdTxs types.StdTx
	for i, _ := range brs {
		cdc := cliCtxs[i].Codec
		gasAdj, _, err := ParseFloat64OrReturnBadRequest(brs[i].GasAdjustment, flags.DefaultGasAdjustment)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		simAndExec, gas, err := flags.ParseGas(brs[i].Gas)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		keyBase, err := keyring.New(sdkTypes.KeyringServiceName(), "os", "home", strings.NewReader(keys.DefaultKeyPass))
		if err != nil {
			panic(fmt.Errorf("couldn't acquire keyring: %v", err))
		}

		address, err := cTypes.AccAddressFromBech32(brs[i].From)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		//adding account sequence
		num, seq, err := types.NewAccountRetriever(authClient.Codec, cliCtxs[i]).GetAccountNumberSequence(address)
		if err != nil {
			fmt.Printf("Error in NewAccountRetriever: %s\n", err)
			return nil, nil
		}

		txBldr := types.NewTxBuilder(
			authClient.GetTxEncoder(cliCtxs[i].Codec), brs[i].AccountNumber, brs[i].Sequence, gas, gasAdj,
			brs[i].Simulate, brs[i].ChainID, brs[i].Memo, brs[i].Fees, brs[i].GasPrices,
		)

		txBldr = txBldr.WithKeybase(keyBase)

		if brs[i].Simulate || simAndExec {
			if gasAdj < 0 {
				return nil, errors.New(err.Error())
			}

			txBldr, err = tx.EnrichWithGas(txBldr, cliCtxs[i], []cTypes.Msg{msgs[i]})
			if err != nil {
				return nil, errors.New(err.Error())
			}

			if brs[i].Simulate {
				val, _ := SimulationResponse(cdc, txBldr.Gas())
				return val, nil
			}
		}

		txBldr = txBldr.WithAccountNumber(num)
		txBldr = txBldr.WithSequence(seq)
		fromName := cliCtxs[i].GetFromName()

		//build and sign
		stdMsg, err := txBldr.BuildAndSign(fromName, keys.DefaultKeyPass, msgs)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		stdMsgs, err := txBldr.BuildSignMsg(msgs)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		var count = uint64(0)
		for j := 0; j < i; j++ {
			if txBldr.AccountNumber() == brs[j].AccountNumber {
				count++
			}
		}

		if i == 0 {
			stdTxs.Msgs = stdMsgs.Msgs
			stdTxs.Fee = stdMsgs.Fee
			stdTxs.Memo = stdMsgs.Memo
		}

		// broadcast to a node
		res, err := cliCtxs[i].BroadcastTx(stdMsg)
		if err != nil {
			fmt.Printf("Error in broadcast: %s\n", err)
			return nil, nil
		}

		output, err := cliCtxs[i].Codec.MarshalJSON(res)

		fmt.Printf("output: %s\n", output)
		return output, nil

	}
	return nil, nil
}

//older func used in comdex crust

func SignAndBroadcastMultiple(brs []rest.BaseReq, cliCtxs []context.CLIContext,
	mode []string, passwords []string, msgs []cTypes.Msg) ([]byte, error) {

	var stdTxs types.StdTx
	for i, _ := range brs {

		cdc := cliCtxs[i].Codec
		gasAdj, _, err := ParseFloat64OrReturnBadRequest(brs[i].GasAdjustment, flags.DefaultGasAdjustment)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		simAndExec, gas, err := flags.ParseGas(brs[i].Gas)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		keyBase, err := keyring.New(sdkTypes.KeyringServiceName(), "os", "home", strings.NewReader(keys.DefaultKeyPass))
		if err != nil {
			panic(fmt.Errorf("couldn't acquire keyring: %v", err))
		}

		address, err := cTypes.AccAddressFromBech32(brs[i].From)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		num, _, err := auth.NewAccountRetriever(authClient.Codec, cliCtxs[i]).GetAccountNumberSequence(address)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		brs[i].AccountNumber = num

		txBldr := types.NewTxBuilder(
			authClient.GetTxEncoder(cliCtxs[i].Codec), brs[i].AccountNumber, brs[i].Sequence, gas, gasAdj,
			brs[i].Simulate, brs[i].ChainID, brs[i].Memo, brs[i].Fees, brs[i].GasPrices,
		)

		txBldr = txBldr.WithKeybase(keyBase)

		if brs[i].Simulate || simAndExec {
			if gasAdj < 0 {
				return nil, errors.New(err.Error())
			}

			txBldr, err = tx.EnrichWithGas(txBldr, cliCtxs[i], []cTypes.Msg{msgs[i]})
			if err != nil {
				return nil, errors.New(err.Error())
			}

			if brs[i].Simulate {
				val, _ := SimulationResponse(cdc, txBldr.Gas())
				return val, nil
			}
		}

		stdMsg, err := txBldr.BuildSignMsg(msgs)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		stdTx := auth.NewStdTx(stdMsg.Msgs, stdMsg.Fee, nil, stdMsg.Memo)

		stdTx, err = SignStdTxFromRest(txBldr, cliCtxs[i], cliCtxs[i].GetFromName(), stdTx, true, false, passwords[i])
		if err != nil {
			return nil, errors.New(err.Error())
		}

		var count = uint64(0)
		for j := 0; j < i; j++ {
			if txBldr.AccountNumber() == brs[j].AccountNumber {
				count++
			}
		}
		txBldr = txBldr.WithSequence(count)

		if i == 0 {
			stdTxs.Msgs = stdTx.Msgs
			stdTxs.Fee = stdTx.Fee
			stdTxs.Memo = stdTx.Memo
		}

		if count == 0 {
			stdTxs.Signatures = append(stdTxs.Signatures, stdTx.Signatures...)
		}
	}
	val, _ := BroadcastRest(cliCtxs[0], cliCtxs[0].Codec, stdTxs, mode[0])

	return val, nil

}
