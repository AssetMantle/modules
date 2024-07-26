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

func QueueOrBroadcastTransaction(context client.Context, commonTransactionRequest rest.CommonTransactionRequest, msg sdkTypes.Msg) (err error) {
	fromAddress, fromName, _, err := client.GetFromFields(context, context.Keyring, commonTransactionRequest.GetFrom())
	if err != nil {
		return err
	}

	context = context.WithFromAddress(fromAddress).WithFromName(fromName).WithSkipConfirmation(true).WithGenerateOnly(commonTransactionRequest.IsSimulated())

	if KafkaState.IsEnabled {
		if err = context.PrintBytes(SendToKafka(newKafkaMsgFromRest(msg, TicketID(random.GenerateUniqueIdentifier(reflect.TypeOf(msg).String())), commonTransactionRequest, context), context.LegacyAmino)); err != nil {
			return err
		} else {
			return nil
		}
	}

	gasAdjustment := flags.DefaultGasAdjustment
	if len(commonTransactionRequest.GetGasAdjustment()) != 0 {
		if gasAdjustment, err = strconv.ParseFloat(commonTransactionRequest.GetGasAdjustment(), 64); err != nil {
			return err
		}
	}

	gasSetting, err := flags.ParseGasSetting(commonTransactionRequest.GetGas())
	if err != nil {
		return err
	}

	transactionFactory := tx.Factory{}.
		WithFees(commonTransactionRequest.GetFees().String()).
		WithGasPrices(commonTransactionRequest.GetGasPrices().String()).
		WithAccountNumber(commonTransactionRequest.GetAccountNumber()).
		WithAccountRetriever(context.AccountRetriever).
		WithSequence(commonTransactionRequest.GetSequence()).
		WithGas(gasSetting.Gas).
		WithGasAdjustment(gasAdjustment).
		WithMemo(commonTransactionRequest.GetMemo()).
		WithChainID(commonTransactionRequest.GetChainID()).
		WithSimulateAndExecute(gasSetting.Simulate || commonTransactionRequest.IsSimulated()).
		WithTxConfig(context.TxConfig).
		WithTimeoutHeight(commonTransactionRequest.GetTimeoutHeight()).
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
