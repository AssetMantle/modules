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
	moduleName        string
	defaultParamspace string
	queryRoute        string
	transactionRoute  string
	genesisState      helpers.GenesisState
	mapper            helpers.Mapper
	auxiliaryList     []helpers.Auxiliary
	queryList         []helpers.Query
	transactionList   []helpers.Transaction
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
	module.genesisState.RegisterCodec(codec)
	for _, transaction := range module.transactionList {
		transaction.RegisterCodec(codec)
	}
	for _, query := range module.queryList {
		query.RegisterCodec(codec)
	}
}
func (module module) DefaultGenesis() json.RawMessage {
	return module.genesisState.Default().Marshall()
}
func (module module) ValidateGenesis(rawMessage json.RawMessage) error {
	genesisState := module.genesisState.Unmarshall(rawMessage)
	return genesisState.Validate()
}
func (module module) RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	for _, query := range module.queryList {
		router.HandleFunc(query.GetRoute()+"/{query}", query.RESTQueryHandler(cliContext)).Methods("GET")
	}

	for _, transaction := range module.transactionList {
		router.HandleFunc(transaction.GetRoute(), transaction.RESTRequestHandler(cliContext)).Methods("POST")
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
	for _, transaction := range module.transactionList {
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
	for _, query := range module.queryList {
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

		for _, transaction := range module.transactionList {
			if msg.Type() == transaction.GetName() {
				return transaction.HandleMessage(context, msg)
			}
		}
		return nil, errors.New(fmt.Sprintf("Unknown message type, %v for module %v", msg.Type(), module.Name()))
	}
}
func (module module) QuerierRoute() string {
	return module.moduleName
}
func (module module) NewQuerierHandler() sdkTypes.Querier {
	return func(context sdkTypes.Context, path []string, requestQuery abciTypes.RequestQuery) ([]byte, error) {
		for _, query := range module.queryList {
			if query.GetName() == path[0] {
				return query.HandleMessage(context, requestQuery)
			}
		}
		return nil, errors.New(fmt.Sprintf("Unknown query path, %v for module %v", path[0], module.Name()))
	}
}
func (module module) InitGenesis(context sdkTypes.Context, rawMessage json.RawMessage) []abciTypes.ValidatorUpdate {
	genesisState := module.genesisState.Unmarshall(rawMessage)
	genesisState.Initialize(context, module.mapper)
	return []abciTypes.ValidatorUpdate{}
}
func (module module) ExportGenesis(context sdkTypes.Context) json.RawMessage {
	return module.genesisState.Export(context, module.mapper).Marshall()
}
func (module module) BeginBlock(_ sdkTypes.Context, _ abciTypes.RequestBeginBlock) {}

func (module module) EndBlock(_ sdkTypes.Context, _ abciTypes.RequestEndBlock) []abciTypes.ValidatorUpdate {
	return []abciTypes.ValidatorUpdate{}
}
func (module module) GetKVStoreKey() *sdkTypes.KVStoreKey {
	return module.mapper.GetKVStoreKey()
}
func (module module) GetDefaultParamspace() string {
	return module.defaultParamspace
}
func (module module) GetAuxiliary(auxiliaryName string) helpers.Auxiliary {
	for _, auxiliary := range module.auxiliaryList {
		if auxiliary.GetName() == auxiliaryName {
			return auxiliary
		}
	}
	panic(fmt.Sprintf("auxiliary %v not found/initialized", auxiliaryName))
}

func (module module) DecodeModuleTransactionRequest(transactionName string, rawMessage json.RawMessage) (sdkTypes.Msg, error) {
	for _, transaction := range module.transactionList {
		if transaction.GetName() == transactionName {
			return transaction.DecodeTransactionRequest(rawMessage)
		}
	}
	return nil, xprtErrors.IncorrectMessage
}

func (module module) Initialize(auxiliaryKeepers ...interface{}) helpers.Module {

	for _, auxiliary := range module.auxiliaryList {
		auxiliary.InitializeKeeper(module.mapper, auxiliaryKeepers...)
	}

	for _, transaction := range module.transactionList {
		transaction.InitializeKeeper(module.mapper, auxiliaryKeepers...)
	}

	for _, query := range module.queryList {
		query.InitializeKeeper(module.mapper, auxiliaryKeepers...)
	}

	return module
}
func NewModule(moduleName string, defaultParamspace string, queryRoute string, transactionRoute string, genesisState helpers.GenesisState, mapper helpers.Mapper, auxiliaryList []helpers.Auxiliary, queryList []helpers.Query, transactionList []helpers.Transaction) helpers.Module {
	return module{
		moduleName:        moduleName,
		defaultParamspace: defaultParamspace,
		queryRoute:        queryRoute,
		transactionRoute:  transactionRoute,
		genesisState:      genesisState,
		mapper:            mapper,
		auxiliaryList:     auxiliaryList,
		queryList:         queryList,
		transactionList:   transactionList,
	}
}
