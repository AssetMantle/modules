package base

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type module struct {
	moduleName        string
	storeKey          string
	defaultParamspace string
	queryRoute        string
	transactionRoute  string
	genesisState      utilities.GenesisState
	mapper            utilities.Mapper
	auxiliaryList     []utilities.Auxiliary
	queryList         []utilities.Query
	transactionList   []utilities.Transaction
}

var _ utilities.Module = (*module)(nil)

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
func (module module) DefaultGenesis(jsonMarshaler codec.JSONMarshaler) json.RawMessage {
	return jsonMarshaler.MustMarshalJSON(module.genesisState.Default())
}
func (module module) ValidateGenesis(_ codec.JSONMarshaler, rawMessage json.RawMessage) error {
	genesisState := module.genesisState.Unmarshall(rawMessage)
	return genesisState.Validate()
}
func (module module) RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	for _, query := range module.queryList {
		//TODO pick from query flags
		router.HandleFunc(fmt.Sprintf("/%v/%v/{%v}", query.GetModuleName(), query.GetName(), "id"), query.RESTQueryHandler(cliContext)).Methods("GET")
	}

	for _, transaction := range module.transactionList {
		router.HandleFunc(transaction.GetRoute(), transaction.RESTRequestHandler(cliContext)).Methods("POST")
		router.HandleFunc("/keys/add", RESTKeysHandler(cliContext)).Methods("POST")
		router.HandleFunc("/recover/keys", RESTKeysRecoverHandler(cliContext)).Methods("POST")
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
func (module module) InitGenesis(context sdkTypes.Context, _ codec.JSONMarshaler, rawMessage json.RawMessage) []abciTypes.ValidatorUpdate {
	genesisState := module.genesisState.Unmarshall(rawMessage)
	genesisState.Initialize(context)
	return []abciTypes.ValidatorUpdate{}
}
func (module module) ExportGenesis(context sdkTypes.Context, _ codec.JSONMarshaler) json.RawMessage {
	return module.genesisState.Export(context).Marshall()
}
func (module module) BeginBlock(_ sdkTypes.Context, _ abciTypes.RequestBeginBlock) {}

func (module module) EndBlock(_ sdkTypes.Context, _ abciTypes.RequestEndBlock) []abciTypes.ValidatorUpdate {
	return []abciTypes.ValidatorUpdate{}
}
func (module module) GetStoreKey() string {
	return module.storeKey
}
func (module module) GetDefaultParamspace() string {
	return module.defaultParamspace
}
func (module module) GetAuxiliaryKeepers(auxiliaryNames ...string) []utilities.AuxiliaryKeeper {
	var auxiliaryKeeperList []utilities.AuxiliaryKeeper
	for _, auxiliaryName := range auxiliaryNames {
		for _, auxiliary := range module.auxiliaryList {
			if auxiliary.GetName() == auxiliaryName {
				auxiliaryKeeperList = append(auxiliaryKeeperList, auxiliary.GetKeeper())
			}
		}
	}
	return auxiliaryKeeperList
}
func (module module) InitializeKeepers(codec *codec.Codec, storeKey sdkTypes.StoreKey, _ params.Subspace, auxiliaryKeepers ...interface{}) {
	mapper := module.mapper.InitializeMapper(codec, storeKey)

	for _, auxiliary := range module.auxiliaryList {
		auxiliary.InitializeKeeper(mapper)
	}

	for _, transaction := range module.transactionList {
		transaction.InitializeKeeper(mapper, auxiliaryKeepers...)
	}

	for _, query := range module.queryList {
		query.InitializeKeeper(mapper)
	}

	return
}
func NewModule(moduleName string, storeKey string, defaultParamspace string, queryRoute string, transactionRoute string, genesisState utilities.GenesisState, mapper utilities.Mapper, auxiliaryList []utilities.Auxiliary, queryList []utilities.Query, transactionList []utilities.Transaction) utilities.Module {
	return module{
		moduleName:        moduleName,
		storeKey:          storeKey,
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
