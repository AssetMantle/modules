/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/utilities/rest/queuing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type transaction struct {
	name             string
	cliCommand       helpers.CLICommand
	keeper           helpers.TransactionKeeper
	requestPrototype func() helpers.TransactionRequest
	messagePrototype func() helpers.Message
	keeperPrototype  func() helpers.TransactionKeeper
}

// TODO remove

var KafkaBool = false
var KafkaState queuing.KafkaState

var _ helpers.Transaction = (*transaction)(nil)

func (transaction transaction) GetName() string { return transaction.name }
func (transaction transaction) Command(codec *codec.Codec) *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		bufioReader := bufio.NewReader(command.InOrStdin())
		transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(authClient.GetTxEncoder(codec))
		cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)

		transactionRequest, Error := transaction.requestPrototype().FromCLI(transaction.cliCommand, cliContext)
		if Error != nil {
			return Error
		}

		msg, Error := transactionRequest.MakeMsg()
		if Error != nil {
			return Error
		}

		if Error := msg.ValidateBasic(); Error != nil {
			return Error
		}

		return authClient.GenerateOrBroadcastMsgs(cliContext, transactionBuilder, []sdkTypes.Msg{msg})
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
			sdkTypes.NewAttribute(sdkTypes.AttributeKeyModule, transaction.name),
		),
	)

	return &sdkTypes.Result{Events: context.EventManager().Events()}, nil
}

func (transaction transaction) RESTRequestHandler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := transaction.requestPrototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.Codec, &transactionRequest) {
			return
		} else if reflect.TypeOf(transaction.requestPrototype()) != reflect.TypeOf(transactionRequest) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		Error := transactionRequest.Validate()
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		baseReq := transactionRequest.GetBaseReq()

		msg, Error := transactionRequest.MakeMsg()
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		baseReq = baseReq.Sanitize()
		if !baseReq.ValidateBasic(responseWriter) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		if Error := msg.ValidateBasic(); Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		if viper.GetBool(flags.FlagGenerateOnly) {
			authClient.WriteGenerateStdTxResponse(responseWriter, cliContext, baseReq, []sdkTypes.Msg{msg})
			return
		}

		gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(responseWriter, baseReq.GasAdjustment, flags.DefaultGasAdjustment)
		if !ok {
			return
		}

		simAndExec, gas, Error := flags.ParseGas(baseReq.Gas)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		txBuilder := types.NewTxBuilder(
			authClient.GetTxEncoder(cliContext.Codec), baseReq.AccountNumber, baseReq.Sequence, gas, gasAdj,
			baseReq.Simulate, baseReq.ChainID, baseReq.Memo, baseReq.Fees, baseReq.GasPrices,
		)
		msgList := []sdkTypes.Msg{msg}

		if baseReq.Simulate || simAndExec {
			if gasAdj < 0 {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, errors.ErrOutOfGas.Error())
				return
			}

			txBuilder, Error = authClient.EnrichWithGas(txBuilder, cliContext, msgList)
			if Error != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
				return
			}

			if baseReq.Simulate {
				rest.WriteSimulationResponse(responseWriter, cliContext.Codec, txBuilder.Gas())
				return
			}
		}

		fromAddress, fromName, Error := context.GetFromFields(strings.NewReader(keys.DefaultKeyPass), baseReq.From, viper.GetBool(flags.FlagGenerateOnly))
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		cliContext = cliContext.WithFromAddress(fromAddress)
		cliContext = cliContext.WithFromName(fromName)
		cliContext = cliContext.WithBroadcastMode(viper.GetString(flags.FlagBroadcastMode))

		if KafkaBool {
			ticketID := queuing.TicketIDGenerator(transaction.name)
			jsonResponse := queuing.SendToKafka(queuing.NewKafkaMsgFromRest(msg, ticketID, baseReq, cliContext), KafkaState, cliContext.Codec)

			responseWriter.WriteHeader(http.StatusAccepted)
			_, _ = responseWriter.Write(jsonResponse)
		} else {
			accountNumber, sequence, Error := types.NewAccountRetriever(cliContext).GetAccountNumberSequence(fromAddress)
			if Error != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
				return
			}

			txBuilder = txBuilder.WithAccountNumber(accountNumber)
			txBuilder = txBuilder.WithSequence(sequence)

			stdMsg, Error := txBuilder.BuildAndSign(fromName, keys.DefaultKeyPass, msgList)
			if Error != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
				return
			}

			// broadcast to a node
			response, Error := cliContext.BroadcastTx(stdMsg)
			if Error != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
				return
			}

			output, Error := cliContext.Codec.MarshalJSON(response)
			if Error != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
				return
			}

			responseWriter.Header().Set("Content-Type", "application/json")
			if _, Error := responseWriter.Write(output); Error != nil {
				log.Printf("could not write response: %v", Error)
			}
		}
	}
}

func (transaction transaction) RegisterCodec(codec *codec.Codec) {
	transaction.messagePrototype().RegisterCodec(codec)
	transaction.requestPrototype().RegisterCodec(codec)
}
func (transaction transaction) DecodeTransactionRequest(rawMessage json.RawMessage) (sdkTypes.Msg, error) {
	transactionRequest, Error := transaction.requestPrototype().FromJSON(rawMessage)
	if Error != nil {
		return nil, Error
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
