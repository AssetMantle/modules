// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/utilities/rest/queuing"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gogo/protobuf/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"net/http"
	"reflect"
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
		Context, err := client.GetClientTxContext(command)
		if err != nil {
			return err
		}

		transactionRequest, err := transaction.requestPrototype().FromCLI(transaction.cliCommand, Context)
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

		return tx.GenerateOrBroadcastTxCLI(Context, command.Flags(), msg)
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
			rest.CheckBadRequestError(responseWriter, fmt.Errorf("expected %s, got %s", reflect.TypeOf(transaction.requestPrototype()), reflect.TypeOf(transactionRequest)))
			return
		} else if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
			return
		}

		baseReq := transactionRequest.GetBaseReq().Sanitize()
		if !baseReq.ValidateBasic(responseWriter) {
			rest.CheckBadRequestError(responseWriter, fmt.Errorf("invalid base request"))
		}

		msg, err := transactionRequest.MakeMsg()
		if err != nil {
			rest.CheckBadRequestError(responseWriter, err)
			return
		}

		if rest.CheckInternalServerError(responseWriter, queuing.QueueOrBroadcastTransaction(context.WithOutput(responseWriter), baseReq, msg)) {
			return
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
		panic(fmt.Errorf("keeper for transaction %s is not initialized", transaction.name))
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
