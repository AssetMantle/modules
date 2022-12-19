// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/utilities/rest/queuing"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"reflect"
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

//TODO: Replace EmitEvent with EmitTypedEvent
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
		if !rest.ReadRESTReq(responseWriter, httpRequest, context.LegacyAmino, &transactionRequest) {
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
			tx.WriteGeneratedTxResponse(context, responseWriter, baseReq, []sdkTypes.Msg{msg}...)
			return
		}

		gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(responseWriter, baseReq.GasAdjustment, flags.DefaultGasAdjustment)
		if !ok {
			return
		}

		gasSetting, err := flags.ParseGasSetting(baseReq.Gas)
		if rest.CheckBadRequestError(responseWriter, err) {
			return
		}

		transactionFactory := tx.Factory{}.
			WithFees(baseReq.Fees.String()).
			WithGasPrices(baseReq.GasPrices.String()).
			WithAccountNumber(baseReq.AccountNumber).
			WithSequence(baseReq.Sequence).
			WithGas(gasSetting.Gas).
			WithGasAdjustment(gasAdj).
			WithMemo(baseReq.Memo).
			WithChainID(baseReq.ChainID).
			WithSimulateAndExecute(baseReq.Simulate).
			WithTxConfig(context.TxConfig).
			WithTimeoutHeight(baseReq.TimeoutHeight)

		msgList := []sdkTypes.Msg{msg}

		if baseReq.Simulate || gasSetting.Simulate {
			if gasAdj < 0 {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, errors.ErrOutOfGas.Error())
				return
			}

			_, adjusted, err := tx.CalculateGas(context, transactionFactory, msgList...)
			if rest.CheckInternalServerError(responseWriter, err) {
				return
			}

			transactionFactory = transactionFactory.WithGas(adjusted)

			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			if baseReq.Simulate {
				rest.WriteSimulationResponse(responseWriter, context.LegacyAmino, transactionFactory.Gas())
				return
			}
		}

		var fromAddress sdkTypes.AccAddress
		var fromName string

		fromAddress, fromName, _, err = client.GetFromFields(context.Keyring, baseReq.From, viper.GetBool(flags.FlagGenerateOnly))
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
				context.LegacyAmino,
			)

			if _, err = responseWriter.Write(output); err != nil {
				log.Printf("could not write response: %v", err)
			}
		} else {
			var accountNumber uint64
			var sequence uint64

			accountNumber, sequence, err = types.AccountRetriever{}.GetAccountNumberSequence(context, fromAddress)
			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			transactionFactory = transactionFactory.WithAccountNumber(accountNumber).WithSequence(sequence)

			var transactionBuilder client.TxBuilder

			transactionBuilder, err = transactionFactory.BuildUnsignedTx(msgList...)

			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			err = tx.Sign(transactionFactory, fromName, transactionBuilder, true)

			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			var response *sdkTypes.TxResponse

			var transactionBytes []byte

			transactionBytes, err = context.TxConfig.TxEncoder()(transactionBuilder.GetTx())

			if err != nil {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
				return
			}

			// broadcast to a node

			response, err = context.BroadcastTx(transactionBytes)

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
