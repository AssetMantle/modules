// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/utilities/rest/queuing"
)

type transaction struct {
	name                 string
	cliCommand           helpers.CLICommand
	keeper               helpers.TransactionKeeper
	requestPrototype     func() helpers.TransactionRequest
	messagePrototype     func() helpers.Message
	keeperPrototype      func() helpers.TransactionKeeper
	serviceRegistrar     func(grpc.Server, helpers.TransactionKeeper)
	grpcGatewayRegistrar func(client.Context, *runtime.ServeMux) error
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

		if err = transactionRequest.Validate(); err != nil {
			return err
		}

		msg, err := transactionRequest.MakeMsg()
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
func (transaction transaction) HandleMessage(context context.Context, message helpers.Message) (*sdkTypes.Result, error) {
	if transactionResponse, err := transaction.keeper.Transact(context, message); err != nil {
		return nil, err
	} else {
		return transactionResponse.GetResult(), nil
	}
}
func (transaction transaction) RESTRequestHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := transaction.requestPrototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, context.LegacyAmino, &transactionRequest) {
			return
		} else if reflect.TypeOf(transaction.requestPrototype()) != reflect.TypeOf(transactionRequest) {
			rest.CheckBadRequestError(responseWriter, errorConstants.InvalidRequest.Wrapf("expected %s, got %s", reflect.TypeOf(transaction.requestPrototype()), reflect.TypeOf(transactionRequest)))
			return
		}

		if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
			return
		}
		baseReq := transactionRequest.GetBaseReq()
		baseReq = baseReq.Sanitize()
		if !baseReq.ValidateBasic(responseWriter) {

		}

		msg, err := transactionRequest.MakeMsg()
		if err != nil {
			rest.CheckBadRequestError(responseWriter, err)
			return
		}

		if rest.CheckBadRequestError(responseWriter, msg.ValidateBasic()) {
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
			WithTimeoutHeight(baseReq.TimeoutHeight).
			WithKeybase(context.Keyring)

		msgList := []sdkTypes.Msg{msg}

		if baseReq.Simulate || gasSetting.Simulate {
			if gasAdj < 0 {
				rest.CheckBadRequestError(responseWriter, errors.ErrOutOfGas)
				return
			}

			_, adjusted, err := tx.CalculateGas(context, transactionFactory, msgList...)
			if rest.CheckInternalServerError(responseWriter, err) {
				return
			}

			transactionFactory = transactionFactory.WithGas(adjusted)

			if baseReq.Simulate {
				rest.WriteSimulationResponse(responseWriter, context.LegacyAmino, transactionFactory.Gas())
				return
			}
		}

		fromAddress, fromName, _, err := client.GetFromFields(context.Keyring, baseReq.From, viper.GetBool(flags.FlagGenerateOnly))
		if err != nil {
			rest.CheckBadRequestError(responseWriter, err)
			return
		}

		context = context.WithFromAddress(fromAddress)
		context = context.WithFromName(fromName)
		// TODO ***** from from client.toml, remove hardcode
		context = context.WithBroadcastMode(flags.BroadcastBlock)

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

			accountNumber, sequence, err := types.AccountRetriever{}.GetAccountNumberSequence(context, fromAddress)
			if err != nil {
				rest.CheckBadRequestError(responseWriter, err)
				return
			}

			transactionFactory = transactionFactory.WithAccountNumber(accountNumber).WithSequence(sequence)
			transactionBuilder, err := transactionFactory.BuildUnsignedTx(msgList...)
			if err != nil {
				rest.CheckBadRequestError(responseWriter, err)
				return
			}

			if rest.CheckBadRequestError(responseWriter, tx.Sign(transactionFactory, fromName, transactionBuilder, true)) {
				return
			}

			transactionBytes, err := context.TxConfig.TxEncoder()(transactionBuilder.GetTx())
			if err != nil {
				rest.CheckBadRequestError(responseWriter, err)
				return
			}

			response, err := context.BroadcastTx(transactionBytes)
			if err != nil {
				rest.CheckBadRequestError(responseWriter, err)
				return
			}

			output, err := context.Codec.MarshalJSON(response)
			if err != nil {
				rest.CheckBadRequestError(responseWriter, err)
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
func (transaction transaction) RegisterInterfaces(interfaceRegistry codecTypes.InterfaceRegistry) {
	transaction.messagePrototype().RegisterInterface(interfaceRegistry)
}
func (transaction transaction) RegisterService(configurator sdkModuleTypes.Configurator) {
	if transaction.keeper == nil {
		panic(errorConstants.UninitializedUsage.Wrapf("keeper for transaction %s is not initialized", transaction.name))
	}
	transaction.serviceRegistrar(configurator.MsgServer(), transaction.keeper)
}
func (transaction transaction) RegisterGRPCGatewayRoute(context client.Context, serveMux *runtime.ServeMux) {
	if err := transaction.grpcGatewayRegistrar(context, serveMux); err != nil {
		panic(err)
	}
}
func (transaction transaction) DecodeTransactionRequest(rawMessage json.RawMessage) (sdkTypes.Msg, error) {
	transactionRequest, err := transaction.requestPrototype().FromJSON(rawMessage)
	if err != nil {
		return nil, err
	}

	return transactionRequest.MakeMsg()
}
func (transaction transaction) InitializeKeeper(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaryKeepers ...interface{}) helpers.Transaction {
	transaction.keeper = transaction.keeperPrototype().Initialize(mapper, parameterManager, auxiliaryKeepers).(helpers.TransactionKeeper)
	return transaction
}

func NewTransaction(name string, short string, long string, requestPrototype func() helpers.TransactionRequest, messagePrototype func() helpers.Message, keeperPrototype func() helpers.TransactionKeeper, serviceRegistrar func(grpc.Server, helpers.TransactionKeeper), grpcGatewayRegistrar func(client.Context, *runtime.ServeMux) error, flagList ...helpers.CLIFlag) helpers.Transaction {
	return transaction{
		name:                 name,
		cliCommand:           NewCLICommand(name, short, long, flagList),
		requestPrototype:     requestPrototype,
		messagePrototype:     messagePrototype,
		keeperPrototype:      keeperPrototype,
		serviceRegistrar:     serviceRegistrar,
		grpcGatewayRegistrar: grpcGatewayRegistrar,
	}
}
