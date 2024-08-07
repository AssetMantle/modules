// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"
	"fmt"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	"math/rand"

	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/AssetMantle/modules/helpers"
)

type module struct {
	name             string
	consensusVersion uint64

	auxiliariesPrototype      func() helpers.Auxiliaries
	blockPrototype            func() helpers.Block
	genesisPrototype          func() helpers.Genesis
	invariantsPrototype       func() helpers.Invariants
	mapperPrototype           func() helpers.Mapper
	parameterManagerPrototype func() helpers.ParameterManager
	queriesPrototype          func() helpers.Queries
	simulatorPrototype        func() helpers.Simulator
	transactionsPrototype     func() helpers.Transactions

	auxiliaries      helpers.Auxiliaries
	genesis          helpers.Genesis
	mapper           helpers.Mapper
	parameterManager helpers.ParameterManager
	queries          helpers.Queries
	transactions     helpers.Transactions
	block            helpers.Block
}

var _ helpers.Module = (*module)(nil)

func (module module) IsOnePerModuleType() {
	//TODO implement me
	panic("implement me")
}

func (module module) IsAppModule() {
	//TODO implement me
	panic("implement me")
}
func (module module) Name() string {
	return module.name
}
func (module module) GetTransactions() helpers.Transactions {
	return module.transactions
}
func (module module) RegisterLegacyAminoCodec(legacyAmino *sdkCodec.LegacyAmino) {
	for _, transaction := range module.transactionsPrototype().Get() {
		transaction.RegisterLegacyAminoCodec(legacyAmino)
	}
}
func (module module) RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	for _, transaction := range module.transactionsPrototype().Get() {
		transaction.RegisterInterfaces(interfaceRegistry)
	}
}
func (module module) DefaultGenesis(jsonCodec sdkCodec.JSONCodec) json.RawMessage {
	return module.genesisPrototype().Default().Encode(jsonCodec)
}
func (module module) ValidateGenesis(jsonCodec sdkCodec.JSONCodec, _ client.TxEncodingConfig, rawMessage json.RawMessage) error {
	genesisState := module.genesisPrototype().Decode(jsonCodec, rawMessage)
	return genesisState.ValidateBasic(module.parameterManagerPrototype())
}
func (module module) RegisterRESTRoutes(context client.Context, router *mux.Router) {
	router.HandleFunc("/"+module.Name()+"/parameters", module.parameterManagerPrototype().RESTQueryHandler(context)).Methods("GET")

	for _, query := range module.queriesPrototype().Get() {
		router.HandleFunc("/"+module.Name()+"/"+query.GetName(), query.RESTQueryHandler(context)).Methods("GET")
	}

	for _, transaction := range module.transactionsPrototype().Get() {
		router.HandleFunc("/"+module.Name()+"/"+transaction.GetName(), transaction.RESTRequestHandler(context)).Methods("POST")
	}
}
func (module module) RegisterGRPCGatewayRoutes(context client.Context, serveMux *runtime.ServeMux) {
	for _, query := range module.queriesPrototype().Get() {
		query.RegisterGRPCGatewayRoute(context, serveMux)
	}

	for _, transaction := range module.transactionsPrototype().Get() {
		transaction.RegisterGRPCGatewayRoute(context, serveMux)
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
	commandList := make([]*cobra.Command, len(module.transactionsPrototype().Get()))

	for i, transaction := range module.transactionsPrototype().Get() {
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
	commandList := make([]*cobra.Command, len(module.queriesPrototype().Get()))

	for i, query := range module.queriesPrototype().Get() {
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
func (module module) ProposalMsgs(simulationState sdkModuleTypes.SimulationState) []simulationTypes.WeightedProposalMsg {
	//TODO implement me
	panic("implement me")
}
func (module module) ProposalContents(simulationState sdkModuleTypes.SimulationState) []simulationTypes.WeightedProposalContent {
	return module.simulatorPrototype().WeightedProposalContentList(simulationState)
}
func (module module) RandomizedParams(r *rand.Rand) []simulationTypes.LegacyParamChange {
	return module.simulatorPrototype().ParamChangeList(r)
}
func (module module) RegisterStoreDecoder(storeDecoderRegistry sdkTypes.StoreDecoderRegistry) {
	storeDecoderRegistry[module.name] = module.mapperPrototype().StoreDecoder
}
func (module module) WeightedOperations(simulationState sdkModuleTypes.SimulationState) []simulationTypes.WeightedOperation {
	return module.simulatorPrototype().WeightedOperations(simulationState, module)
}
func (module module) RegisterInvariants(invariantRegistry sdkTypes.InvariantRegistry) {
	module.invariantsPrototype().Register(invariantRegistry)
}

// TODO remove if unnecessary
//
//	func (module module) Route() sdkTypes.Route {
//		return sdkTypes.NewRoute(module.Name(), func(context sdkTypes.Context, msg sdkTypes.Msg) (*sdkTypes.Result, error) {
//			if module.transactions == nil {
//				panic(fmt.Errorf("transactions for module %s not initialized", module.Name()))
//			}
//
//			if message, ok := msg.(helpers.Message); ok {
//				if transaction := module.transactions.GetTransaction(message.Type()); transaction != nil {
//					return transaction.HandleMessage(sdkTypes.WrapSDKContext(context.WithEventManager(sdkTypes.NewEventManager())), message)
//				}
//			}
//			return nil, fmt.Errorf("message type %T is not supported by module %s", msg, module.Name())
//		})
//	}
func (module module) QuerierRoute() string {
	return module.name
}

//	func (module module) LegacyQuerierHandler(_ *sdkCodec.LegacyAmino) sdkTypes.Querier {
//		return func(context sdkTypes.Context, path []string, requestQuery abciTypes.RequestQuery) ([]byte, error) {
//			if module.queries == nil {
//				panic(fmt.Errorf("queries for module %s not initialized", module.Name()))
//			}
//
//			if query := module.queries.GetQuery(path[0]); query != nil {
//				return query.HandleQuery(sdkTypes.WrapSDKContext(context), requestQuery)
//			}
//
//			if path[0] == "parameters" {
//				return CodecPrototype().MarshalJSON(module.parameterManager.Fetch(sdkTypes.WrapSDKContext(context)).Get())
//			}
//
//			return nil, fmt.Errorf("unknown query path, %v for module %v", path[0], module.Name())
//		}
//	}
func (module module) RegisterServices(configurator sdkModuleTypes.Configurator) {
	for _, query := range module.queries.Get() {
		query.RegisterService(configurator)
	}

	for _, transaction := range module.transactions.Get() {
		transaction.RegisterService(configurator)
	}
}
func (module module) ConsensusVersion() uint64 {
	return module.consensusVersion
}
func (module module) InitGenesis(context sdkTypes.Context, jsonCodec sdkCodec.JSONCodec, rawMessage json.RawMessage) []abciTypes.ValidatorUpdate {
	genesisState := module.genesisPrototype().Decode(jsonCodec, rawMessage)

	if module.mapper == nil || module.parameterManager == nil {
		panic(fmt.Errorf("mapper or parameter manager for module %s not initialized", module.Name()))
	}

	genesisState.Import(sdkTypes.WrapSDKContext(context), module.mapper, module.parameterManager)

	return []abciTypes.ValidatorUpdate{}
}
func (module module) ExportGenesis(context sdkTypes.Context, jsonCodec sdkCodec.JSONCodec) json.RawMessage {
	if module.mapper == nil || module.parameterManager == nil {
		panic(fmt.Errorf("mapper or parameter manager for module %s not initialized", module.Name()))
	}

	return module.genesisPrototype().Export(sdkTypes.WrapSDKContext(context), module.mapper, module.parameterManager).Encode(jsonCodec)
}
func (module module) BeginBlock(context sdkTypes.Context, beginBlockRequest abciTypes.RequestBeginBlock) {
	module.block.Begin(sdkTypes.WrapSDKContext(context), beginBlockRequest)
}
func (module module) EndBlock(context sdkTypes.Context, endBlockRequest abciTypes.RequestEndBlock) []abciTypes.ValidatorUpdate {
	module.block.End(sdkTypes.WrapSDKContext(context), endBlockRequest)
	return []abciTypes.ValidatorUpdate{}
}
func (module module) GetAuxiliary(auxiliaryName string) helpers.Auxiliary {
	if module.auxiliaries != nil {
		if auxiliary := module.auxiliaries.GetAuxiliary(auxiliaryName); auxiliary != nil {
			return auxiliary
		}
	}

	panic(fmt.Errorf("auxiliary %v not found/initialized", auxiliaryName))
}
func (module module) DecodeModuleTransactionRequest(transactionName string, rawMessage json.RawMessage) (sdkTypes.Msg, error) {
	if transaction := module.transactionsPrototype().GetTransaction(transactionName); transaction != nil {
		return transaction.DecodeTransactionRequest(rawMessage)
	}

	return nil, fmt.Errorf("transaction %s is not supported by module %s", transactionName, module.Name())
}
func (module module) Initialize(kvStoreKey *storeTypes.KVStoreKey, paramsSubspace paramsTypes.Subspace, auxiliaryKeepers ...interface{}) helpers.Module {
	module.mapper = module.mapperPrototype().Initialize(kvStoreKey)

	module.genesis = module.genesisPrototype()

	module.parameterManager = module.parameterManagerPrototype().Initialize(paramsSubspace.WithKeyTable(module.parameterManagerPrototype().GetKeyTable()))

	auxiliaryList := make([]helpers.Auxiliary, len(module.auxiliariesPrototype().Get()))

	for i, auxiliary := range module.auxiliariesPrototype().Get() {
		auxiliaryList[i] = auxiliary.Initialize(module.mapper, module.parameterManager, auxiliaryKeepers...)
	}

	module.auxiliaries = NewAuxiliaries(auxiliaryList...)

	for _, auxiliary := range auxiliaryList {
		auxiliaryKeepers = append(auxiliaryKeepers, auxiliary)
	}

	transactionList := make([]helpers.Transaction, len(module.transactionsPrototype().Get()))

	for i, transaction := range module.transactionsPrototype().Get() {
		transactionList[i] = transaction.InitializeKeeper(module.mapper, module.parameterManager, auxiliaryKeepers...)
	}

	module.transactions = NewTransactions(transactionList...)

	queryList := make([]helpers.Query, len(module.queriesPrototype().Get()))

	for i, query := range module.queriesPrototype().Get() {
		queryList[i] = query.Initialize(module.mapper, module.parameterManager, auxiliaryKeepers...)
	}

	module.queries = NewQueries(queryList...)

	module.block = module.blockPrototype().Initialize(module.mapper, module.parameterManager, auxiliaryKeepers...)

	return module
}

func NewModule(name string, consensusVersion uint64, auxiliariesPrototype func() helpers.Auxiliaries, blockPrototype func() helpers.Block, genesisPrototype func() helpers.Genesis, invariantsPrototype func() helpers.Invariants, mapperPrototype func() helpers.Mapper, parameterManagerPrototype func() helpers.ParameterManager, queriesPrototype func() helpers.Queries, simulatorPrototype func() helpers.Simulator, transactionsPrototype func() helpers.Transactions) helpers.Module {
	return module{
		name:                      name,
		consensusVersion:          consensusVersion,
		auxiliariesPrototype:      auxiliariesPrototype,
		blockPrototype:            blockPrototype,
		genesisPrototype:          genesisPrototype,
		invariantsPrototype:       invariantsPrototype,
		mapperPrototype:           mapperPrototype,
		parameterManagerPrototype: parameterManagerPrototype,
		queriesPrototype:          queriesPrototype,
		simulatorPrototype:        simulatorPrototype,
		transactionsPrototype:     transactionsPrototype,
	}
}
