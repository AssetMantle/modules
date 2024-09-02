// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/utilities/rest"
	"github.com/AssetMantle/modules/utilities/rest/queuing"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type transaction struct {
	serviceName      string
	cliCommand       helpers.CLICommand
	keeper           helpers.TransactionKeeper
	requestPrototype func() helpers.TransactionRequest
	messagePrototype func() helpers.Message
	keeperPrototype  func() helpers.TransactionKeeper
	serviceRegistrar func(grpc.ServiceRegistrar, helpers.TransactionKeeper)
}

var _ helpers.Transaction = (*transaction)(nil)

func (transaction transaction) GetServicePath() string {
	splits := strings.Split(transaction.serviceName, ".")
	return "/" + splits[3] + "/" + splits[5]
}
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
		transactionRequest, err := transaction.requestPrototype().FromHTTPRequest(httpRequest)
		if err != nil {
			rest.CheckBadRequestError(responseWriter, err)
			return
		}

		if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
			return
		}

		commonTransactionRequest := transactionRequest.GetCommonTransactionRequest().Sanitize()
		if !commonTransactionRequest.ValidateBasic(responseWriter) {
			rest.CheckBadRequestError(responseWriter, fmt.Errorf("invalid base request "))
		}

		msg, err := transactionRequest.MakeMsg()
		if err != nil {
			rest.CheckBadRequestError(responseWriter, err)
			return
		}

		if rest.CheckInternalServerError(responseWriter, queuing.QueueOrBroadcastTransaction(context.WithOutput(responseWriter), commonTransactionRequest, msg)) {
			return
		}
	}
}
func (transaction transaction) RegisterLegacyAminoCodec(legacyAmino *sdkCodec.LegacyAmino) {
	transaction.requestPrototype().RegisterLegacyAminoCodec(legacyAmino)
}
func (transaction transaction) RegisterInterfaces(interfaceRegistry codecTypes.InterfaceRegistry) {
	transaction.messagePrototype().RegisterInterface(interfaceRegistry)
}
func (transaction transaction) RegisterService(configurator sdkModuleTypes.Configurator) {
	if transaction.keeper == nil {
		panic(fmt.Errorf("keeper for transaction %s is not initialized", transaction.serviceName))
	}
	transaction.serviceRegistrar(configurator.MsgServer(), transaction.keeper)
}
func (transaction transaction) InitializeKeeper(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaryKeepers ...interface{}) helpers.Transaction {
	transaction.keeper = transaction.keeperPrototype().Initialize(mapper, parameterManager, auxiliaryKeepers).(helpers.TransactionKeeper)
	return transaction
}

func NewTransaction(serviceName string, short string, long string, requestPrototype func() helpers.TransactionRequest, messagePrototype func() helpers.Message, keeperPrototype func() helpers.TransactionKeeper, serviceRegistrar func(grpc.ServiceRegistrar, helpers.TransactionKeeper), flagList ...helpers.CLIFlag) helpers.Transaction {
	splits := strings.Split(serviceName, ".")
	return transaction{
		serviceName:      serviceName,
		cliCommand:       NewCLICommand(splits[5], short, long, flagList),
		requestPrototype: requestPrototype,
		messagePrototype: messagePrototype,
		keeperPrototype:  keeperPrototype,
		serviceRegistrar: serviceRegistrar,
	}
}
