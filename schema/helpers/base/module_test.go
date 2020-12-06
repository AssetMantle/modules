package base

import (
	"encoding/json"
	clientContext "github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"math/rand"
	"testing"
)

var auxiliariesPrototype = func() helpers.Auxiliaries { return auxiliaries{} }
var genesisPrototype = func() helpers.Genesis { return genesis{} }
var mapperPrototype = func() helpers.Mapper { return NewMapper(base.KeyPrototype, base.MappablePrototype) }
var parametersPrototype = func() helpers.Parameters { return parameters{} }
var queriesPrototype = func() helpers.Queries { return queries{} }
var simulatorPrototype = func() helpers.Simulator { return nil }
var transactionsPrototype = func() helpers.Transactions { return transactions{} }

func TestModule(t *testing.T) {
	context, storeKey := base.SetupTest(t)
	codec := base.MakeCodec()
	Module := NewModule("test", auxiliariesPrototype, genesisPrototype,
		mapperPrototype, parametersPrototype, queriesPrototype, simulatorPrototype, transactionsPrototype).Initialize(storeKey, params.NewSubspace(codec, storeKey, nil, "test")).(module)

	//AppModuleBasic
	Module.Name()

	// RegisterCodec
	Module.RegisterCodec(codec)

	// DefaultGenesis
	// Module.DefaultGenesis()

	// ValidateGenesis
	// Module.ValidateGenesis(json.RawMessage{})

	// RegisterRESTRoutes
	cliContext := clientContext.NewCLIContext().WithCodec(codec).WithChainID("test")
	router := mux.NewRouter()
	Module.RegisterRESTRoutes(cliContext, router)

	// GetTxCmd
	Module.GetTxCmd(codec)
	Module.GetQueryCmd(codec)

	//AppModule
	Module.RegisterInvariants(nil)
	Module.Route()
	Module.NewHandler()
	Module.QuerierRoute()
	Module.NewQuerierHandler()
	Module.BeginBlock(context, abciTypes.RequestBeginBlock{})
	Module.EndBlock(context, abciTypes.RequestEndBlock{})
	//	Module.InitGenesis(context, json.RawMessage{})
	//	Module.ExportGenesis(context)

	// AppModuleSimulation
	Module.GenerateGenesisState(&sdkModule.SimulationState{})
	Module.ProposalContents(sdkModule.SimulationState{})
	Module.RandomizedParams(&rand.Rand{})
	Module.RegisterStoreDecoder(sdkTypes.StoreDecoderRegistry{})
	Module.WeightedOperations(sdkModule.SimulationState{})

	//types.Module
	//	Module.GetAuxiliary("")
	_, Error := Module.DecodeModuleTransactionRequest("", json.RawMessage{})
	require.Nil(t, Error)

}
