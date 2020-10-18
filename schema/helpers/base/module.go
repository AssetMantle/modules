/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/gorilla/mux"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"math/rand"
)

type module struct {
	moduleName string

	simulatorPrototype    func() helpers.Simulator
	parametersPrototype   func() helpers.Parameters
	genesisPrototype      func() helpers.Genesis
	auxiliariesPrototype  func() helpers.Auxiliaries
	queriesPrototype      func() helpers.Queries
	transactionsPrototype func() helpers.Transactions

	mapper       helpers.Mapper
	genesis      helpers.Genesis
	parameters   helpers.Parameters
	auxiliaries  helpers.Auxiliaries
	queries      helpers.Queries
	transactions helpers.Transactions
}

var _ helpers.Module = (*module)(nil)

func (module module) GenerateGenesisState(_ *sdkTypesModule.SimulationState) {
	return
}

func (module module) ProposalContents(_ sdkTypesModule.SimulationState) []simulation.WeightedProposalContent {
	return nil
}

func (module module) RandomizedParams(_ *rand.Rand) []simulation.ParamChange {
	return nil
}

func (module module) RegisterStoreDecoder(storeDecoderRegistry sdkTypes.StoreDecoderRegistry) {
	storeDecoderRegistry[module.moduleName] = module.mapper.StoreDecoder
}

func (module module) WeightedOperations(_ sdkTypesModule.SimulationState) []simulation.WeightedOperation {
	return nil
}

func (module module) Name() string {
	return module.moduleName
}
func (module module) RegisterCodec(codec *codec.Codec) {
	module.mapper.RegisterCodec(codec)
	for _, transaction := range module.transactionsPrototype().GetList() {
		transaction.RegisterCodec(codec)
	}
}
func (module module) DefaultGenesis() json.RawMessage {
	return module.genesisPrototype().Default().Encode()
}
func (module module) ValidateGenesis(rawMessage json.RawMessage) error {
	genesisState := module.genesisPrototype().Decode(rawMessage)
	return genesisState.Validate()
}
func (module module) RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	if module.queries == nil || module.transactions == nil {
		panic(xprtErrors.UninitializedUsage)
	}
	for _, query := range module.queries.GetList() {
		router.HandleFunc(module.Name()+"/"+query.GetName()+fmt.Sprintf("/{%s}", query.GetName()), query.RESTQueryHandler(cliContext)).Methods("GET")
	}

	for _, transaction := range module.transactions.GetList() {
		router.HandleFunc(module.Name()+"/"+transaction.GetName(), transaction.RESTRequestHandler(cliContext)).Methods("POST")
	}
}
func (module module) GetTxCmd(codec *codec.Codec) *cobra.Command {
	rootTransactionCommand := &cobra.Command{
		Use:                        module.moduleName,
		Short:                      "Get root transaction command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	var commandList []*cobra.Command
	for _, transaction := range module.transactionsPrototype().GetList() {
		commandList = append(commandList, transaction.Command(codec))
	}
	rootTransactionCommand.AddCommand(
		commandList...,
	)
	return rootTransactionCommand
}
func (module module) GetQueryCmd(codec *codec.Codec) *cobra.Command {
	rootQueryCommand := &cobra.Command{
		Use:                        module.moduleName,
		Short:                      "Get root query command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	var commandList []*cobra.Command
	for _, query := range module.queriesPrototype().GetList() {
		commandList = append(commandList, query.Command(codec))
	}
	rootQueryCommand.AddCommand(
		commandList...,
	)
	return rootQueryCommand
}
func (module module) RegisterInvariants(_ sdkTypes.InvariantRegistry) {}
func (module module) Route() string {
	return module.moduleName
}
func (module module) NewHandler() sdkTypes.Handler {
	return func(context sdkTypes.Context, msg sdkTypes.Msg) (*sdkTypes.Result, error) {
		context = context.WithEventManager(sdkTypes.NewEventManager())
		if module.transactions == nil {
			panic(xprtErrors.UninitializedUsage)
		}
		if transaction := module.transactions.Get(msg.Type()); transaction != nil {
			return transaction.HandleMessage(context, msg)
		}
		return nil, errors.New(fmt.Sprintf("Unknown message type, %v for module %v", msg.Type(), module.Name()))
	}
}
func (module module) QuerierRoute() string {
	return module.moduleName
}
func (module module) NewQuerierHandler() sdkTypes.Querier {
	return func(context sdkTypes.Context, path []string, requestQuery abciTypes.RequestQuery) ([]byte, error) {
		if module.queries == nil {
			panic(xprtErrors.UninitializedUsage)
		}
		if query := module.queries.Get(path[0]); query != nil {
			return query.HandleMessage(context, requestQuery)
		}
		return nil, errors.New(fmt.Sprintf("Unknown query path, %v for module %v", path[0], module.Name()))
	}
}
func (module module) InitGenesis(context sdkTypes.Context, rawMessage json.RawMessage) []abciTypes.ValidatorUpdate {
	genesisState := module.genesisPrototype().Decode(rawMessage)
	if module.mapper == nil || module.parameters == nil {
		panic(xprtErrors.UninitializedUsage)
	}
	genesisState.Import(context, module.mapper, module.parameters)
	return []abciTypes.ValidatorUpdate{}
}
func (module module) ExportGenesis(context sdkTypes.Context) json.RawMessage {
	if module.mapper == nil || module.parameters == nil {
		panic(xprtErrors.UninitializedUsage)
	}
	return module.genesisPrototype().Export(context, module.mapper, module.parameters).Encode()
}
func (module module) BeginBlock(_ sdkTypes.Context, _ abciTypes.RequestBeginBlock) {}

func (module module) EndBlock(_ sdkTypes.Context, _ abciTypes.RequestEndBlock) []abciTypes.ValidatorUpdate {
	return []abciTypes.ValidatorUpdate{}
}
func (module module) GetAuxiliary(auxiliaryName string) helpers.Auxiliary {
	if module.auxiliaries != nil {
		if auxiliary := module.auxiliaries.Get(auxiliaryName); auxiliary != nil {
			return auxiliary
		}
	}
	panic(fmt.Sprintf("auxiliary %v not found/initialized", auxiliaryName))
}

func (module module) DecodeModuleTransactionRequest(transactionName string, rawMessage json.RawMessage) (sdkTypes.Msg, error) {
	if transaction := module.transactionsPrototype().Get(transactionName); transaction != nil {
		return transaction.DecodeTransactionRequest(rawMessage)
	}
	return nil, xprtErrors.IncorrectMessage
}

func (module module) Initialize(kvStoreKey *sdkTypes.KVStoreKey, paramsSubspace params.Subspace, auxiliaryKeepers ...interface{}) helpers.Module {
	module.mapper = module.mapper.Initialize(kvStoreKey)
	//TODO initialize genesis
	module.genesis = module.genesisPrototype().Initialize(nil, nil)
	module.parameters = module.parametersPrototype().Initialize(paramsSubspace.WithKeyTable(module.parametersPrototype().GetKeyTable()))

	var auxiliaryList []helpers.Auxiliary
	for _, auxiliary := range module.auxiliariesPrototype().GetList() {
		auxiliaryList = append(auxiliaryList, auxiliary.Initialize(module.mapper, module.parameters, auxiliaryKeepers...))
	}
	module.auxiliaries = NewAuxiliaries(auxiliaryList...)

	var transactionList []helpers.Transaction
	for _, transaction := range module.transactionsPrototype().GetList() {
		transactionList = append(transactionList, transaction.InitializeKeeper(module.mapper, module.parameters, auxiliaryKeepers...))
	}
	module.transactions = NewTransactions(transactionList...)

	var queryList []helpers.Query
	for _, query := range module.queriesPrototype().GetList() {
		queryList = append(queryList, query.Initialize(module.mapper, module.parameters, auxiliaryKeepers...))
	}
	module.queries = NewQueries(queryList...)
	return module
}

func NewModule(moduleName string, simulatorPrototype func() helpers.Simulator, parametersPrototype func() helpers.Parameters, genesisPrototype func() helpers.Genesis, auxiliariesPrototype func() helpers.Auxiliaries, queriesPrototype func() helpers.Queries, transactionsPrototype func() helpers.Transactions) helpers.Module {
	return module{
		moduleName:            moduleName,
		simulatorPrototype:    simulatorPrototype,
		parametersPrototype:   parametersPrototype,
		genesisPrototype:      genesisPrototype,
		auxiliariesPrototype:  auxiliariesPrototype,
		queriesPrototype:      queriesPrototype,
		transactionsPrototype: transactionsPrototype,
	}
}
