// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

type module struct {
	name             string
	consensusVersion uint64

	auxiliariesPrototype  func() helpers.Auxiliaries
	blockPrototype        func() helpers.Block
	genesisPrototype      func() helpers.Genesis
	invariantsPrototype   func() helpers.Invariants
	mapperPrototype       func() helpers.Mapper
	migrationsPrototype   func() helpers.Migrations
	parametersPrototype   func() helpers.Parameters
	queriesPrototype      func() helpers.Queries
	simulatorPrototype    func() helpers.Simulator
	transactionsPrototype func() helpers.Transactions

	auxiliaries  helpers.Auxiliaries
	genesis      helpers.Genesis
	mapper       helpers.Mapper
	parameters   helpers.Parameters
	queries      helpers.Queries
	transactions helpers.Transactions
	block        helpers.Block
}

var _ helpers.Module = (*module)(nil)

func (module module) Name() string {
	return module.name
}
func (module module) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	for _, transaction := range module.transactionsPrototype().GetList() {
		transaction.RegisterCodec(codec)
	}
}
func (module module) RegisterInterfaces(_ types.InterfaceRegistry) {}
func (module module) DefaultGenesis(jsonCodec codec.JSONCodec) json.RawMessage {
	return module.genesisPrototype().Default().Encode(jsonCodec)
}
func (module module) ValidateGenesis(jsonCodec codec.JSONCodec, _ client.TxEncodingConfig, rawMessage json.RawMessage) error {
	genesisState := module.genesisPrototype().Decode(jsonCodec, rawMessage)
	return genesisState.Validate()
}
func (module module) RegisterRESTRoutes(context client.Context, router *mux.Router) {
	for _, query := range module.queriesPrototype().GetList() {
		router.HandleFunc("/"+module.Name()+"/"+query.GetName()+fmt.Sprintf("/{%s}", query.GetName()), query.RESTQueryHandler(context)).Methods("GET")
	}

	for _, transaction := range module.transactionsPrototype().GetList() {
		router.HandleFunc("/"+module.Name()+"/"+transaction.GetName(), transaction.RESTRequestHandler(context)).Methods("POST")
	}
}
func (module module) RegisterGRPCGatewayRoutes(context client.Context, serveMux *runtime.ServeMux) {
	for _, query := range module.queriesPrototype().GetList() {
		serveMux.Handle(query.GRPCGatewayHandler(context))
	}
}
func (module module) GetTxCmd() *cobra.Command {
	rootTransactionCommand := &cobra.Command{
		Use:                        module.name,
		Short:                      "GetProperty root transaction command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	commandList := make([]*cobra.Command, len(module.transactionsPrototype().GetList()))

	for i, transaction := range module.transactionsPrototype().GetList() {
		commandList[i] = transaction.Command()
	}

	rootTransactionCommand.AddCommand(
		commandList...,
	)

	return rootTransactionCommand
}
func (module module) GetQueryCmd() *cobra.Command {
	rootQueryCommand := &cobra.Command{
		Use:                        module.name,
		Short:                      "GetProperty root query command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	commandList := make([]*cobra.Command, len(module.queriesPrototype().GetList()))

	for i, query := range module.queriesPrototype().GetList() {
		commandList[i] = query.Command()
	}

	rootQueryCommand.AddCommand(
		commandList...,
	)

	return rootQueryCommand
}
func (module module) GenerateGenesisState(simulationState *sdkModuleTypes.SimulationState) {
	module.simulatorPrototype().RandomizedGenesisState(simulationState)
}
func (module module) ProposalContents(simulationState sdkModuleTypes.SimulationState) []simulationTypes.WeightedProposalContent {
	return module.simulatorPrototype().WeightedProposalContentList(simulationState)
}
func (module module) RandomizedParams(r *rand.Rand) []simulationTypes.ParamChange {
	return module.simulatorPrototype().ParamChangeList(r)
}
func (module module) RegisterStoreDecoder(storeDecoderRegistry sdkTypes.StoreDecoderRegistry) {
	storeDecoderRegistry[module.name] = module.mapperPrototype().StoreDecoder
}
func (module module) WeightedOperations(simulationState sdkModuleTypes.SimulationState) []simulationTypes.WeightedOperation {
	return module.simulatorPrototype().WeightedOperations(simulationState)
}
func (module module) RegisterInvariants(invariantRegistry sdkTypes.InvariantRegistry) {
	module.invariantsPrototype().RegisterInvariants(invariantRegistry)
}
func (module module) Route() sdkTypes.Route {
	return sdkTypes.NewRoute(module.Name(), func(context sdkTypes.Context, msg sdkTypes.Msg) (*sdkTypes.Result, error) {
		if module.transactions == nil {
			panic(constants.UninitializedUsage)
		}

		if message, ok := msg.(helpers.Message); ok {
			if transaction := module.transactions.Get(message.GetType()); transaction != nil {
				return transaction.HandleMessage(context.WithEventManager(sdkTypes.NewEventManager()), message)
			}
		}

		return nil, constants.IncorrectMessage
	})
}
func (module module) QuerierRoute() string {
	return module.name
}
func (module module) LegacyQuerierHandler(_ *codec.LegacyAmino) sdkTypes.Querier {
	return func(context sdkTypes.Context, path []string, requestQuery abciTypes.RequestQuery) ([]byte, error) {
		if module.queries == nil {
			panic(constants.UninitializedUsage)
		}

		if query := module.queries.Get(path[0]); query != nil {
			return query.HandleMessage(context, requestQuery)
		}

		return nil, fmt.Errorf("unknown query path, %v for module %v", path[0], module.Name())
	}
}
func (module module) RegisterServices(configurator sdkModuleTypes.Configurator) {
	for _, query := range module.queriesPrototype().GetList() {
		configurator.QueryServer().RegisterService(query.Service())
	}

	for _, transaction := range module.transactionsPrototype().GetList() {
		configurator.MsgServer().RegisterService(transaction.Service())
	}
}
func (module module) ConsensusVersion() uint64 {
	return module.consensusVersion
}
func (module module) InitGenesis(context sdkTypes.Context, jsonCodec codec.JSONCodec, rawMessage json.RawMessage) []abciTypes.ValidatorUpdate {
	genesisState := module.genesisPrototype().Decode(jsonCodec, rawMessage)

	if module.mapper == nil || module.parameters == nil {
		panic(constants.UninitializedUsage)
	}

	genesisState.Import(context, module.mapper, module.parameters)

	return []abciTypes.ValidatorUpdate{}
}
func (module module) ExportGenesis(context sdkTypes.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	if module.mapper == nil || module.parameters == nil {
		panic(constants.UninitializedUsage)
	}

	return module.genesisPrototype().Export(context, module.mapper, module.parameters).Encode(jsonCodec)
}
func (module module) BeginBlock(context sdkTypes.Context, beginBlockRequest abciTypes.RequestBeginBlock) {
	module.block.Begin(context, beginBlockRequest)
}
func (module module) EndBlock(context sdkTypes.Context, endBlockRequest abciTypes.RequestEndBlock) []abciTypes.ValidatorUpdate {
	return module.block.End(context, endBlockRequest)
}
func (module module) GetAuxiliary(auxiliaryName string) helpers.Auxiliary {
	if module.auxiliaries != nil {
		if auxiliary := module.auxiliaries.Get(auxiliaryName); auxiliary != nil {
			return auxiliary
		}
	}

	panic(fmt.Errorf("auxiliary %v not found/initialized", auxiliaryName))
}
func (module module) DecodeModuleTransactionRequest(transactionName string, rawMessage json.RawMessage) (sdkTypes.Msg, error) {
	if transaction := module.transactionsPrototype().Get(transactionName); transaction != nil {
		return transaction.DecodeTransactionRequest(rawMessage)
	}

	return nil, constants.IncorrectMessage
}
func (module module) Initialize(kvStoreKey *sdkTypes.KVStoreKey, paramsSubspace paramTypes.Subspace, auxiliaryKeepers ...interface{}) helpers.Module {
	module.mapper = module.mapperPrototype().Initialize(kvStoreKey)

	module.genesis = module.genesisPrototype().Initialize(module.genesisPrototype().GetMappableList(), module.genesisPrototype().GetParameterList())

	module.parameters = module.parametersPrototype().Initialize(paramsSubspace.WithKeyTable(module.parametersPrototype().GetKeyTable()))

	auxiliaryList := make([]helpers.Auxiliary, len(module.auxiliariesPrototype().GetList()))

	for i, auxiliary := range module.auxiliariesPrototype().GetList() {
		auxiliaryList[i] = auxiliary.Initialize(module.mapper, module.parameters, auxiliaryKeepers...)
	}

	module.auxiliaries = NewAuxiliaries(auxiliaryList...)

	for _, auxiliary := range auxiliaryList {
		auxiliaryKeepers = append(auxiliaryKeepers, auxiliary)
	}

	transactionList := make([]helpers.Transaction, len(module.transactionsPrototype().GetList()))

	for i, transaction := range module.transactionsPrototype().GetList() {
		transactionList[i] = transaction.InitializeKeeper(module.mapper, module.parameters, auxiliaryKeepers...)
	}

	module.transactions = NewTransactions(transactionList...)

	queryList := make([]helpers.Query, len(module.queriesPrototype().GetList()))

	for i, query := range module.queriesPrototype().GetList() {
		queryList[i] = query.Initialize(module.mapper, module.parameters, auxiliaryKeepers...)
	}

	module.queries = NewQueries(queryList...)

	module.block = module.blockPrototype().Initialize(module.mapper, module.parameters, auxiliaryKeepers...)

	return module
}

func NewModule(name string, consensusVersion uint64, auxiliariesPrototype func() helpers.Auxiliaries, blockPrototype func() helpers.Block, genesisPrototype func() helpers.Genesis, invariantsPrototype func() helpers.Invariants, mapperPrototype func() helpers.Mapper, migrationsPrototype func() helpers.Migrations, parametersPrototype func() helpers.Parameters, queriesPrototype func() helpers.Queries, simulatorPrototype func() helpers.Simulator, transactionsPrototype func() helpers.Transactions) helpers.Module {
	return module{
		name:                  name,
		consensusVersion:      consensusVersion,
		auxiliariesPrototype:  auxiliariesPrototype,
		blockPrototype:        blockPrototype,
		genesisPrototype:      genesisPrototype,
		invariantsPrototype:   invariantsPrototype,
		mapperPrototype:       mapperPrototype,
		migrationsPrototype:   migrationsPrototype,
		parametersPrototype:   parametersPrototype,
		queriesPrototype:      queriesPrototype,
		simulatorPrototype:    simulatorPrototype,
		transactionsPrototype: transactionsPrototype,
	}
}
