/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bufio"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/utilities/rest/queuing"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
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

		request := transaction.transactionRequestPrototype().FromCLI(transaction.cliCommand, cliContext)

		msg := request.MakeMsg()
		if Error := msg.ValidateBasic(); Error != nil {
			return Error
		}

		return authClient.GenerateOrBroadcastMsgs(cliContext, transactionBuilder, []sdkTypes.Msg{msg})
	}
	return transaction.cliCommand.CreateCommand(runE)
}

func (transaction transaction) HandleMessage(context sdkTypes.Context, message sdkTypes.Msg) (*sdkTypes.Result, error) {

	if Error := (transaction.transactionKeeper).Transact(context, message); Error != nil {
		return nil, Error
	}

	context.EventManager().EmitEvent(
		sdkTypes.NewEvent(
			sdkTypes.EventTypeMessage,
			sdkTypes.NewAttribute(sdkTypes.AttributeKeyModule, transaction.moduleName),
		),
	)

	return &sdkTypes.Result{Events: context.EventManager().ABCIEvents()}, nil
}

func (transaction transaction) RESTRequestHandler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		request := transaction.transactionRequestPrototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.Codec, &request) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		baseReq := request.GetBaseReq()
		msg := request.MakeMsg()

		baseReq = baseReq.Sanitize()
		if !baseReq.ValidateBasic(responseWriter) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		Error := msg.ValidateBasic()
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}
		//client.WriteGenerateStdTxResponse(responseWriter, cliContext, baseReq, []sdkTypes.Msg{msg})
		//adding below commands to REST to have signed txs
		gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(responseWriter, baseReq.GasAdjustment, flags.DefaultGasAdjustment)
		if !ok {
			return
		}

		simAndExec, gas, err := flags.ParseGas(baseReq.Gas)
		if rest.CheckBadRequestError(responseWriter, err) {
			return
		}

		txBuilder := types.NewTxBuilder(
			authClient.GetTxEncoder(cliContext.Codec), baseReq.AccountNumber, baseReq.Sequence, gas, gasAdj,
			baseReq.Simulate, baseReq.ChainID, baseReq.Memo, baseReq.Fees, baseReq.GasPrices,
		)

		msgList := []sdkTypes.Msg{msg}
		fromName := cliContext.GetFromName()

		if baseReq.Simulate || simAndExec {
			if gasAdj < 0 {
				rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, errors.ErrorInvalidGasAdjustment.Error())
				return
			}

			txBuilder, err = authClient.EnrichWithGas(txBuilder, cliContext, msgList)
			if rest.CheckInternalServerError(responseWriter, err) {
				return
			}

			if baseReq.Simulate {
				rest.WriteSimulationResponse(responseWriter, cliContext.Codec, txBuilder.Gas())
				return
			}
		}

		//using DefaultKeyPass as an input
		Keyring, err := keyring.New(sdkTypes.KeyringServiceName(), keyring.BackendPass, os.ExpandEnv("$HOME/.assetClient"), strings.NewReader(keys.DefaultKeyPass))
		if err != nil {
			panic(fmt.Errorf("couldn't acquire keyring: %v", err))
		}

		fromAddress, fromName, err := context.GetFromFields(Keyring, baseReq.From, false)
		if err != nil {
			fmt.Printf("failed to get from fields: %v\n", err)
			return
		}

		cliContext = cliContext.WithFromAddress(fromAddress)
		cliContext = cliContext.WithFromName(fromName)
		cliContext = cliContext.WithBroadcastMode("block")

		//adding account sequence
		num, seq, err := types.NewAccountRetriever(authClient.Codec, cliContext).GetAccountNumberSequence(fromAddress)
		if err != nil {
			fmt.Printf("Error in NewAccountRetriever: %s\n", err)
			return
		}

		txBuilder = txBuilder.WithAccountNumber(num)
		txBuilder = txBuilder.WithSequence(seq)

		//build and sign
		stdMsg, err := txBuilder.BuildAndSign(fromName, keys.DefaultKeyPass, msgList)
		if rest.CheckBadRequestError(responseWriter, err) {
			return
		}

		if KafkaBool == true {
			ticketID := queuing.TicketIDGenerator("assetM")
			jsonResponse := queuing.SendToKafka(queuing.NewKafkaMsgFromRest(msg, ticketID, baseReq, cliContext), KafkaState, cliContext.Codec)
			responseWriter.WriteHeader(http.StatusAccepted)
			_, _ = responseWriter.Write(jsonResponse)
		} else {
			// broadcast to a node
			res, err := cliContext.BroadcastTx(stdMsg)
			if err != nil {
				fmt.Printf("Error in broadcast: %s\n", err)
				return
			}

			output, err := cliContext.Codec.MarshalJSON(res)

			responseWriter.Header().Set("Content-Type", "application/json")
			if _, err := responseWriter.Write(output); err != nil {
				log.Printf("could not write response: %v", err)
			}
		}

	}
}

func (transaction transaction) RegisterCodec(codec *codec.Codec) {
	transaction.registerCodec(codec)
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
