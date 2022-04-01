// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/utilities/random"
	"github.com/persistenceOne/persistenceSDK/utilities/rest/queuing"
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
func (transaction transaction) Command(codec *codec.Codec) *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		bufioReader := bufio.NewReader(command.InOrStdin())
		transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(authClient.GetTxEncoder(codec))
		cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)

		transactionRequest, err := transaction.requestPrototype().FromCLI(transaction.cliCommand, cliContext)
		if err != nil {
			return err
		}

		var msg sdkTypes.Msg

		msg, err = transactionRequest.MakeMsg()
		if err != nil {
			return err
		}

		if err = msg.ValidateBasic(); err != nil {
			return err
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
			sdkTypes.NewAttribute(sdkTypes.AttributeKeyModule, message.Route()),
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

		err := transactionRequest.Validate()
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
			return
		}

		baseReq := transactionRequest.GetBaseReq()

		var msg sdkTypes.Msg
		msg, err = transactionRequest.MakeMsg()
		// TODO write one method
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
			return
		}

		baseReq = baseReq.Sanitize()
		if !baseReq.ValidateBasic(responseWriter) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		if err = msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
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

		var simAndExec bool
		var gas uint64

		simAndExec, gas, err = flags.ParseGas(baseReq.Gas)
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
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

			txBuilder, err = authClient.EnrichWithGas(txBuilder, cliContext, msgList)
			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			if baseReq.Simulate {
				rest.WriteSimulationResponse(responseWriter, cliContext.Codec, txBuilder.Gas())
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

		cliContext = cliContext.WithFromAddress(fromAddress)
		cliContext = cliContext.WithFromName(fromName)
		cliContext = cliContext.WithBroadcastMode(viper.GetString(flags.FlagBroadcastMode))

		if queuing.KafkaState.IsEnabled {
			responseWriter.WriteHeader(http.StatusAccepted)

			output := queuing.SendToKafka(queuing.NewKafkaMsgFromRest(
				msg,
				queuing.TicketID(random.GenerateUniqueIdentifier(transaction.name)),
				baseReq,
				cliContext),
				cliContext.Codec,
			)

			if _, err = responseWriter.Write(output); err != nil {
				log.Printf("could not write response: %v", err)
			}
		} else {
			var accountNumber uint64
			var sequence uint64

			accountNumber, sequence, err = types.NewAccountRetriever(cliContext).GetAccountNumberSequence(fromAddress)
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

			response, err = cliContext.BroadcastTx(stdMsg)
			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			var output []byte

			output, err = cliContext.Codec.MarshalJSON(response)
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

func (transaction transaction) RegisterCodec(codec *codec.Codec) {
	transaction.messagePrototype().RegisterCodec(codec)
	transaction.requestPrototype().RegisterCodec(codec)
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
