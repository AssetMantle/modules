/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypesModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type module struct {
	name string

	auxiliariesPrototype  func() helpers.Auxiliaries
	genesisPrototype      func() helpers.Genesis
	mapperPrototype       func() helpers.Mapper
	parametersPrototype   func() helpers.Parameters
	queriesPrototype      func() helpers.Queries
	simulatorPrototype    func() helpers.Simulator
	transactionsPrototype func() helpers.Transactions
	blockPrototype        func() helpers.Block

	auxiliaries  helpers.Auxiliaries
	genesis      helpers.Genesis
	mapper       helpers.Mapper
	parameters   helpers.Parameters
	queries      helpers.Queries
	transactions helpers.Transactions
	block        helpers.Block
}

var _ helpers.Module = (*module)(nil)

func (module module) GenerateGenesisState(simulationState *sdkTypesModule.SimulationState) {
	module.simulatorPrototype().RandomizedGenesisState(simulationState)
}

func (module module) ProposalContents(_ sdkTypesModule.SimulationState) []simulation.WeightedProposalContent {
	return module.simulatorPrototype().WeightedProposalContentList()
}

func (module module) RandomizedParams(r *rand.Rand) []simulation.ParamChange {
	return module.simulatorPrototype().ParamChangeList(r)
}

func (module module) RegisterStoreDecoder(storeDecoderRegistry sdkTypes.StoreDecoderRegistry) {
	storeDecoderRegistry[module.name] = module.mapperPrototype().StoreDecoder
}

func (module module) WeightedOperations(_ sdkTypesModule.SimulationState) []simulation.WeightedOperation {
	return nil
}

func (module module) Name() string {
	return module.name
}
func (module module) RegisterCodec(codec *codec.Codec) {
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
	for _, query := range module.queriesPrototype().GetList() {
		router.HandleFunc("/"+module.Name()+"/"+query.GetName()+fmt.Sprintf("/{%s}", query.GetName()), query.RESTQueryHandler(cliContext)).Methods("GET")
	}

	for _, transaction := range module.transactionsPrototype().GetList() {
		router.HandleFunc("/"+module.Name()+"/"+transaction.GetName(), transaction.RESTRequestHandler(cliContext)).Methods("POST")
	}
}
func (module module) GetTxCmd(codec *codec.Codec) *cobra.Command {
	rootTransactionCommand := &cobra.Command{
		Use:                        module.name,
		Short:                      "Get root transaction command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	commandList := make([]*cobra.Command, len(module.transactionsPrototype().GetList()))

	for i, transaction := range module.transactionsPrototype().GetList() {
		commandList[i] = transaction.Command(codec)
	}

	rootTransactionCommand.AddCommand(
		commandList...,
	)

	return rootTransactionCommand
}
func (module module) GetQueryCmd(codec *codec.Codec) *cobra.Command {
	rootQueryCommand := &cobra.Command{
		Use:                        module.name,
		Short:                      "Get root query command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	commandList := make([]*cobra.Command, len(module.queriesPrototype().GetList()))

	for i, query := range module.queriesPrototype().GetList() {
		commandList[i] = query.Command(codec)
	}

	rootQueryCommand.AddCommand(
		commandList...,
	)

	return rootQueryCommand
}
func (module module) RegisterInvariants(_ sdkTypes.InvariantRegistry) {}
func (module module) Route() string {
	return module.name
}
func (module module) NewHandler() sdkTypes.Handler {
	return func(context sdkTypes.Context, msg sdkTypes.Msg) (*sdkTypes.Result, error) {
		context = context.WithEventManager(sdkTypes.NewEventManager())

		if module.transactions == nil {
			panic(errors.UninitializedUsage)
		}

		if transaction := module.transactions.Get(msg.Type()); transaction != nil {
			return transaction.HandleMessage(context, msg)
		}

		return nil, fmt.Errorf("unknown message type, %v for module %v", msg.Type(), module.Name())
	}
}
func (module module) QuerierRoute() string {
	return module.name
}
func (module module) NewQuerierHandler() sdkTypes.Querier {
	return func(context sdkTypes.Context, path []string, requestQuery abciTypes.RequestQuery) ([]byte, error) {
		if module.queries == nil {
			panic(errors.UninitializedUsage)
		}

		if query := module.queries.Get(path[0]); query != nil {
			return query.HandleMessage(context, requestQuery)
		}

		return nil, fmt.Errorf("unknown query path, %v for module %v", path[0], module.Name())
	}
}
func (module module) InitGenesis(context sdkTypes.Context, rawMessage json.RawMessage) []abciTypes.ValidatorUpdate {
	genesisState := module.genesisPrototype().Decode(rawMessage)

	if module.mapper == nil || module.parameters == nil {
		panic(errors.UninitializedUsage)
	}

	genesisState.Import(context, module.mapper, module.parameters)

	return []abciTypes.ValidatorUpdate{}
}
func (module module) ExportGenesis(context sdkTypes.Context) json.RawMessage {
	if module.mapper == nil || module.parameters == nil {
		panic(errors.UninitializedUsage)
	}

	return module.genesisPrototype().Export(context, module.mapper, module.parameters).Encode()
}
func (module module) BeginBlock(context sdkTypes.Context, beginBlockRequest abciTypes.RequestBeginBlock) {
	module.block.Begin(context, beginBlockRequest)
}

func (module module) EndBlock(context sdkTypes.Context, endBlockRequest abciTypes.RequestEndBlock) []abciTypes.ValidatorUpdate {
	module.block.End(context, endBlockRequest)
	return []abciTypes.ValidatorUpdate{}
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

	return nil, errors.IncorrectMessage
}

func (module module) Initialize(kvStoreKey *sdkTypes.KVStoreKey, paramsSubspace params.Subspace, auxiliaryKeepers ...interface{}) helpers.Module {
	module.mapper = module.mapperPrototype().Initialize(kvStoreKey)

	module.genesis = module.genesisPrototype().Initialize(module.genesisPrototype().GetMappableList(), module.genesisPrototype().GetParameterList())

	module.parameters = module.parametersPrototype().Initialize(paramsSubspace.WithKeyTable(module.parametersPrototype().GetKeyTable()))

	auxiliaryList := make([]helpers.Auxiliary, len(module.auxiliariesPrototype().GetList()))

	for i, auxiliary := range module.auxiliariesPrototype().GetList() {
		auxiliaryList[i] = auxiliary.Initialize(module.mapper, module.parameters, auxiliaryKeepers...)
	}

	module.auxiliaries = NewAuxiliaries(auxiliaryList...)

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

func NewModule(name string, auxiliariesPrototype func() helpers.Auxiliaries, genesisPrototype func() helpers.Genesis, mapperPrototype func() helpers.Mapper, parametersPrototype func() helpers.Parameters, queriesPrototype func() helpers.Queries, simulatorPrototype func() helpers.Simulator, transactionsPrototype func() helpers.Transactions, blockPrototype func() helpers.Block) helpers.Module {
	return module{
		name:                  name,
		auxiliariesPrototype:  auxiliariesPrototype,
		genesisPrototype:      genesisPrototype,
		mapperPrototype:       mapperPrototype,
		parametersPrototype:   parametersPrototype,
		queriesPrototype:      queriesPrototype,
		simulatorPrototype:    simulatorPrototype,
		transactionsPrototype: transactionsPrototype,
		blockPrototype:        blockPrototype,
	}
}
