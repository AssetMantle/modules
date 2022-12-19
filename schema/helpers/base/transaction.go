// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/utilities/rest/queuing"
)

type transaction struct {
	name             string
	cliCommand       helpers.CLICommand
	keeper           helpers.TransactionKeeper
	requestPrototype func() helpers.TransactionRequest
	messagePrototype func() helpers.Message
	keeperPrototype  func() helpers.TransactionKeeper
}

var _ helpers.Transaction = (*transaction)(nil)

func (transaction transaction) GetName() string { return transaction.name }
func (transaction transaction) Command() *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		context, err := client.GetClientTxContext(command)
		if err != nil {
			return err
		}

		transactionRequest, err := transaction.requestPrototype().FromCLI(transaction.cliCommand, context)
		if err != nil {
			return err
		}

		var msg sdkTypes.Msg
		if er := transactionRequest.Validate(); er != nil {
			return errorConstants.IncorrectFormat
		}
		msg, err = transactionRequest.MakeMsg()
		if err != nil {
			return err
		}

		if err = msg.ValidateBasic(); err != nil {
			return err
		}

		return tx.GenerateOrBroadcastTxCLI(context, command.Flags(), msg)
	}

	return transaction.cliCommand.CreateCommand(runE)
}

func (transaction transaction) HandleMessage(context sdkTypes.Context, message sdkTypes.Msg) (*sdkTypes.Result, error) {
	if transactionResponse := transaction.keeper.Transact(context, message); !transactionResponse.IsSuccessful() {
		return nil, transactionResponse.GetError()
	}

	context.EventManager().EmitEvent(
		sdkTypes.NewEvent(
			sdkTypes.EventTypeMessage,
			sdkTypes.NewAttribute(sdkTypes.AttributeKeyModule, message.Route()),
		),
	)

	return &sdkTypes.Result{Events: context.EventManager().Events()}, nil
}

func (transaction transaction) RESTRequestHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := transaction.requestPrototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, context.Codec, &transactionRequest) {
			return
		} else if reflect.TypeOf(transaction.requestPrototype()) != reflect.TypeOf(transactionRequest) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		err := transactionRequest.Validate()
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
			return
		}

		baseReq := transactionRequest.GetBaseReq()

		baseReq = baseReq.Sanitize()
		if !baseReq.ValidateBasic(responseWriter) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		var msg sdkTypes.Msg
		msg, err = transactionRequest.MakeMsg()
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
			return
		}

		if err = msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
			return
		}

		if viper.GetBool(flags.FlagGenerateOnly) {
			authClient.WriteGenerateStdTxResponse(responseWriter, context, baseReq, []sdkTypes.Msg{msg})
			return
		}

		gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(responseWriter, baseReq.GasAdjustment, flags.DefaultGasAdjustment)
		if !ok {
			return
		}

		var simAndExec bool
		var gas uint64

		simAndExec, gas, err = flags.ParseGas(baseReq.Gas)
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
			return
		}

		txBuilder := types.NewTxBuilder(
			authClient.GetTxEncoder(context.Codec), baseReq.AccountNumber, baseReq.Sequence, gas, gasAdj,
			baseReq.Simulate, baseReq.ChainID, baseReq.Memo, baseReq.Fees, baseReq.GasPrices,
		)
		msgList := []sdkTypes.Msg{msg}

		if baseReq.Simulate || simAndExec {
			if gasAdj < 0 {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, errors.ErrOutOfGas.Error())
				return
			}

			txBuilder, err = authClient.EnrichWithGas(txBuilder, context, msgList)
			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			if baseReq.Simulate {
				rest.WriteSimulationResponse(responseWriter, context.Codec, txBuilder.Gas())
				return
			}
		}

		var fromAddress sdkTypes.AccAddress
		var fromName string

		fromAddress, fromName, err = context.GetFromFields(strings.NewReader(keys.DefaultKeyPass), baseReq.From, viper.GetBool(flags.FlagGenerateOnly))
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
			return
		}

		context = context.WithFromAddress(fromAddress)
		context = context.WithFromName(fromName)
		context = context.WithBroadcastMode(viper.GetString(flags.FlagBroadcastMode))

		if queuing.KafkaState.IsEnabled {
			responseWriter.WriteHeader(http.StatusAccepted)

			output := queuing.SendToKafka(queuing.NewKafkaMsgFromRest(
				msg,
				queuing.TicketID(random.GenerateUniqueIdentifier(transaction.name)),
				baseReq,
				context),
				context.Codec,
			)

			if _, err = responseWriter.Write(output); err != nil {
				log.Printf("could not write response: %v", err)
			}
		} else {
			var accountNumber uint64
			var sequence uint64

			accountNumber, sequence, err = types.NewAccountRetriever(context).GetAccountNumberSequence(fromAddress)
			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			txBuilder = txBuilder.WithAccountNumber(accountNumber)
			txBuilder = txBuilder.WithSequence(sequence)

			var stdMsg []byte

			stdMsg, err = txBuilder.BuildAndSign(fromName, keys.DefaultKeyPass, msgList)
			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			// broadcast to a node
			var response sdkTypes.TxResponse

			response, err = context.BroadcastTx(stdMsg)
			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			var output []byte

			output, err = context.Codec.MarshalJSON(response)
			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			responseWriter.Header().Set("Content-Type", "application/json")
			if _, err = responseWriter.Write(output); err != nil {
				log.Printf("could not write response: %v", err)
			}
		}
	}
}

func (transaction transaction) RegisterLegacyAminoCodec(legacyAmino *sdkCodec.LegacyAmino) {
	transaction.messagePrototype().RegisterLegacyAminoCodec(legacyAmino)
	transaction.requestPrototype().RegisterLegacyAminoCodec(legacyAmino)
}
func (transaction transaction) DecodeTransactionRequest(rawMessage json.RawMessage) (sdkTypes.Msg, error) {
	transactionRequest, err := transaction.requestPrototype().FromJSON(rawMessage)
	if err != nil {
		return nil, err
	}

	return transactionRequest.MakeMsg()
}

func (transaction transaction) InitializeKeeper(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaryKeepers ...interface{}) helpers.Transaction {
	transaction.keeper = transaction.keeperPrototype().Initialize(mapper, parameters, auxiliaryKeepers).(helpers.TransactionKeeper)
	return transaction
}

func NewTransaction(name string, short string, long string, requestPrototype func() helpers.TransactionRequest, messagePrototype func() helpers.Message, keeperPrototype func() helpers.TransactionKeeper, flagList ...helpers.CLIFlag) helpers.Transaction {
	return transaction{
		name:             name,
		cliCommand:       NewCLICommand(name, short, long, flagList),
		requestPrototype: requestPrototype,
		messagePrototype: messagePrototype,
		keeperPrototype:  keeperPrototype,
	}
}
