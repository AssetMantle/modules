// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

type transaction struct {
	name             string
	cliCommand       helpers.CLICommand
	keeper           helpers.TransactionKeeper
	requestPrototype func() helpers.TransactionRequest
	messagePrototype func() helpers.Message
	keeperPrototype  func() helpers.TransactionKeeper
}

func (transaction transaction) Service() (*grpc.ServiceDesc, interface{}) {
	// TODO implement me
	panic("implement me")
}

var _ helpers.Transaction = (*transaction)(nil)

func (transaction transaction) GetName() string { return transaction.name }
func (transaction transaction) Command() *cobra.Command {
	runE := func(command *cobra.Command, args []string) error {
		cliContext, err := client.GetClientTxContext(command)
		if err != nil {
			return err
		}

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

		return tx.GenerateOrBroadcastTxCLI(cliContext, command.Flags(), msg)
	}

	return transaction.cliCommand.CreateCommand(runE)
}

func (transaction transaction) HandleMessage(context sdkTypes.Context, message helpers.Message) (*sdkTypes.Result, error) {
	if transactionResponse := transaction.keeper.Transact(context, message); !transactionResponse.IsSuccessful() {
		return nil, nil
	}

	return &sdkTypes.Result{Events: message.GenerateOnSuccessEvents().ToABCIEvents()}, nil
}

func (transaction transaction) RESTRequestHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := transaction.requestPrototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, context.LegacyAmino, &transactionRequest) {
			return
		} else if reflect.TypeOf(transaction.requestPrototype()) != reflect.TypeOf(transactionRequest) { // unmarshalling can result in a different implementation of the same interface
			rest.CheckBadRequestError(responseWriter, constants.InvalidRequest)
			return
		}

		if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
			return
		}

		baseReq := transactionRequest.GetRequest()
		if !baseReq.ValidateBasic(responseWriter) {
			rest.CheckBadRequestError(responseWriter, constants.InvalidRequest)
			return
		}

		_, err := transactionRequest.MakeMsg()
		if rest.CheckBadRequestError(responseWriter, err) {
			return
		}

		// TODO, allow for fully sign and broadcast txn also - commented for now
		// if viper.GetBool(flags.FlagGenerateOnly) {
		//	tx.WriteGeneratedTxResponse(context, responseWriter, baseReq, msg)
		//	return
		// }
	}
}

func (transaction transaction) RegisterCodec(codec *codec.LegacyAmino) {
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
