package queuing

import (
	"github.com/AssetMantle/modules/utilities/rest"
	"reflect"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/utilities/random"
)

func QueueOrBroadcastTransaction(context client.Context, baseReq rest.BaseReq, msg sdkTypes.Msg) (err error) {
	fromAddress, fromName, _, err := client.GetFromFields(context, context.Keyring, baseReq.From)
	if err != nil {
		return err
	}

	context = context.WithFromAddress(fromAddress).WithFromName(fromName).WithSkipConfirmation(true).WithGenerateOnly(baseReq.Simulate)

	if KafkaState.IsEnabled {
		if err = context.PrintBytes(SendToKafka(NewKafkaMsgFromRest(msg, TicketID(random.GenerateUniqueIdentifier(reflect.TypeOf(msg).String())), baseReq, context), context.LegacyAmino)); err != nil {
			return err
		} else {
			return nil
		}
	}

	gasAdjustment := flags.DefaultGasAdjustment
	if len(baseReq.GasAdjustment) != 0 {
		if gasAdjustment, err = strconv.ParseFloat(baseReq.GasAdjustment, 64); err != nil {
			return err
		}
	}

	gasSetting, err := flags.ParseGasSetting(baseReq.Gas)
	if err != nil {
		return err
	}

	transactionFactory := tx.Factory{}.
		WithFees(baseReq.Fees.String()).
		WithGasPrices(baseReq.GasPrices.String()).
		WithAccountNumber(baseReq.AccountNumber).
		WithAccountRetriever(context.AccountRetriever).
		WithSequence(baseReq.Sequence).
		WithGas(gasSetting.Gas).
		WithGasAdjustment(gasAdjustment).
		WithMemo(baseReq.Memo).
		WithChainID(baseReq.ChainID).
		WithSimulateAndExecute(gasSetting.Simulate || baseReq.Simulate).
		WithTxConfig(context.TxConfig).
		WithTimeoutHeight(baseReq.TimeoutHeight).
		WithKeybase(context.Keyring)

	if context.GenerateOnly {
		transactionFactory, err = transactionFactory.Prepare(context)
		if err != nil {
			return err
		}
	}

	if err := tx.GenerateOrBroadcastTxWithFactory(context, transactionFactory, msg); err != nil {
		return err
	} else {
		return nil
	}
}
