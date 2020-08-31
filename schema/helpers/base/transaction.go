/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	"log"
	"net/http"
	"strings"
)

type transaction struct {
	moduleName                  string
	name                        string
	route                       string
	transactionKeeper           helpers.TransactionKeeper
	cliCommand                  helpers.CLICommand
	registerCodec               func(*codec.Codec)
	initializeKeeper            func(helpers.Mapper, []interface{}) helpers.TransactionKeeper
	transactionRequestPrototype func() helpers.TransactionRequest
}

//declaring global variable
var KafkaBool = false
var KafkaState queuing.KafkaState

var _ helpers.Transaction = (*transaction)(nil)

func (transaction transaction) GetModuleName() string { return transaction.moduleName }
func (transaction transaction) GetName() string       { return transaction.name }
func (transaction transaction) GetRoute() string      { return transaction.route }
func (transaction transaction) Command(codec *codec.Codec) *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		bufioReader := bufio.NewReader(command.InOrStdin())
		transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(authClient.GetTxEncoder(codec))
		cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)

		transactionRequest, Error := transaction.transactionRequestPrototype().FromCLI(transaction.cliCommand, cliContext)
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

	if transactionResponse := transaction.transactionKeeper.Transact(context, message); !transactionResponse.IsSuccessful() {
		return nil, transactionResponse.GetError()
	}

	context.EventManager().EmitEvent(
		sdkTypes.NewEvent(
			sdkTypes.EventTypeMessage,
			sdkTypes.NewAttribute(sdkTypes.AttributeKeyModule, transaction.moduleName),
		),
	)

	return &sdkTypes.Result{Events: context.EventManager().Events()}, nil
}

func (transaction transaction) RESTRequestHandler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := transaction.transactionRequestPrototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.Codec, &transactionRequest) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
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
		//adding below commands to REST to have signed txs
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
			fmt.Printf("failed to get from fields: %v\n", Error)
			return
		}

		cliContext = cliContext.WithFromAddress(fromAddress)
		cliContext = cliContext.WithFromName(fromName)
		cliContext = cliContext.WithBroadcastMode(viper.GetString(flags.FlagBroadcastMode))

		if KafkaBool == true {
			ticketID := queuing.TicketIDGenerator(transaction.name)
			jsonResponse := queuing.SendToKafka(queuing.NewKafkaMsgFromRest(msg, ticketID, baseReq, cliContext), KafkaState, cliContext.Codec)
			responseWriter.WriteHeader(http.StatusAccepted)
			_, _ = responseWriter.Write(jsonResponse)
		} else {
			//adding account sequence
			accountNumber, sequence, Error := types.NewAccountRetriever(cliContext).GetAccountNumberSequence(fromAddress)
			if Error != nil {
				fmt.Printf("Error in NewAccountRetriever: %s\n", Error)
				return
			}

			txBuilder = txBuilder.WithAccountNumber(accountNumber)
			txBuilder = txBuilder.WithSequence(sequence)

			//build and sign
			stdMsg, Error := txBuilder.BuildAndSign(fromName, keys.DefaultKeyPass, msgList)
			if Error != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
				return
			}

			// broadcast to a node
			response, Error := cliContext.BroadcastTx(stdMsg)
			if Error != nil {
				fmt.Printf("Error in broadcast: %s\n", Error)
				return
			}

			output, Error := cliContext.Codec.MarshalJSON(response)

			responseWriter.Header().Set("Content-Type", "application/json")
			if _, Error := responseWriter.Write(output); Error != nil {
				log.Printf("could not write response: %v", Error)
			}
		}

	}
}

func (transaction transaction) RegisterCodec(codec *codec.Codec) {
	transaction.registerCodec(codec)
}
func (transaction transaction) DecodeTransactionRequest(rawMessage json.RawMessage) (sdkTypes.Msg, error) {
	transactionRequest, Error := transaction.transactionRequestPrototype().FromJSON(rawMessage)
	if Error != nil {
		return nil, Error
	}
	return transactionRequest.MakeMsg()
}

func (transaction *transaction) InitializeKeeper(mapper helpers.Mapper, auxiliaryKeepers ...interface{}) {
	transaction.transactionKeeper = transaction.initializeKeeper(mapper, auxiliaryKeepers)
}

func NewTransaction(module string, name string, route string, short string, long string, registerCodec func(*codec.Codec), initializeKeeper func(helpers.Mapper, []interface{}) helpers.TransactionKeeper, transactionRequestPrototype func() helpers.TransactionRequest, flagList []helpers.CLIFlag) helpers.Transaction {
	return &transaction{
		moduleName:                  module,
		name:                        name,
		route:                       route,
		cliCommand:                  NewCLICommand(name, short, long, flagList),
		registerCodec:               registerCodec,
		initializeKeeper:            initializeKeeper,
		transactionRequestPrototype: transactionRequestPrototype,
	}
}
