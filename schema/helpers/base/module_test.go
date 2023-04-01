// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"
	"math/rand"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModuleTypes "github.com/cosmos/cosmos-sdk/types/module"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/utilities/test"
	baseTestUtilities "github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

// TODO: Add grpc gateway handling for tests

var auxiliariesPrototype = func() helpers.Auxiliaries {
	return auxiliaries{[]helpers.Auxiliary{NewAuxiliary("testAuxiliary", baseTestUtilities.TestAuxiliaryKeeperPrototype)}}
}
var genesisPrototype = func() helpers.Genesis {
	return baseTestUtilities.Prototype()
}
var mapperPrototype = func() helpers.Mapper {
	return NewMapper(baseTestUtilities.KeyPrototype, baseTestUtilities.MappablePrototype)
}
var parameterManagerPrototype = func() helpers.ParameterManager {
	return NewParameterManager("", NewValidatableParameter(baseParameters.NewParameter(base.NewMetaProperty(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData"))), func(interface{}) error { return nil }))
}
var queriesPrototype = func() helpers.Queries {
	return queries{[]helpers.Query{NewQuery("testQuery", "q", "testQuery", "test", baseTestUtilities.TestQueryRequestPrototype,
		baseTestUtilities.TestQueryResponsePrototype, baseTestUtilities.TestQueryKeeperPrototype, nil, nil)}}
}
var simulatorPrototype = func() helpers.Simulator { return nil }
var transactionsPrototype = func() helpers.Transactions {
	return transactions{[]helpers.Transaction{NewTransaction("TestMessage", "", "", baseTestUtilities.TestTransactionRequestPrototype, baseTestUtilities.TestMessagePrototype,
		baseTestUtilities.TestTransactionKeeperPrototype, nil, nil)}}
}
var blockPrototype = func() helpers.Block { return baseTestUtilities.TestBlockPrototype() }

func TestModule(t *testing.T) {
	context, storeKey, transientStoreKey := test.SetupTest(t)

	codec := CodecPrototype()

	subspace := paramsTypes.NewSubspace(codec.GetProtoCodec(), codec.GetLegacyAmino(), storeKey, transientStoreKey, "test") // .WithKeyTable(parameterManagerPrototype().GetKeyTable())
	// subspace.SetParamSet(sdkTypes.UnwrapSDKContext(context), parameterManagerPrototype())
	Module := NewModule("test", 1, auxiliariesPrototype, blockPrototype, genesisPrototype, nil,
		mapperPrototype, parameterManagerPrototype, queriesPrototype, simulatorPrototype, transactionsPrototype).Initialize(storeKey, subspace).(module)

	// AppModuleBasic
	require.Equal(t, "test", Module.Name())

	// RegisterLegacyAminoCodec
	Module.RegisterLegacyAminoCodec(codec.GetLegacyAmino())

	require.NotPanics(t, func() {
		Module.DefaultGenesis(codec.GetProtoCodec())
	})

	require.NotPanics(t, func() {
	})
	require.Nil(t, Module.ValidateGenesis(codec.GetProtoCodec(), nil, Module.DefaultGenesis(codec.GetProtoCodec())))

	router := mux.NewRouter()
	require.NotPanics(t, func() {
		Module.RegisterRESTRoutes(TestClientContext, router)
	})

	// GetTxCmd
	require.Equal(t, "test", Module.GetTxCmd().Name())
	require.Equal(t, "test", Module.GetQueryCmd().Name())

	// AppModule
	require.NotPanics(t, func() {
		Module.RegisterInvariants(nil)
	})
	require.Equal(t, "test", Module.Route())

	response, err := Module.Route().Handler()(sdkTypes.UnwrapSDKContext(context), baseTestUtilities.NewTestMessage(sdkTypes.AccAddress("addr"), "id"))
	require.Nil(t, err)
	require.NotNil(t, response)

	require.Equal(t, "test", Module.QuerierRoute())

	encodedRequest, err := Module.queries.Get("testQuery").(query).requestPrototype().Encode()
	require.Nil(t, err)

	queryResponse, err := Module.LegacyQuerierHandler(codec.GetLegacyAmino())(sdkTypes.UnwrapSDKContext(context), []string{"testQuery"}, abciTypes.RequestQuery{Data: encodedRequest})
	require.Nil(t, err)
	require.NotNil(t, queryResponse)

	require.NotPanics(t, func() {
		Module.BeginBlock(sdkTypes.UnwrapSDKContext(context), abciTypes.RequestBeginBlock{})
	})
	endBlockResponse := Module.EndBlock(sdkTypes.UnwrapSDKContext(context), abciTypes.RequestEndBlock{})
	require.Equal(t, []abciTypes.ValidatorUpdate{}, endBlockResponse)

	require.NotPanics(t, func() {
		Module.InitGenesis(sdkTypes.UnwrapSDKContext(context), codec.GetProtoCodec(), Module.DefaultGenesis(codec.GetProtoCodec()))
	})

	require.Equal(t, Module.DefaultGenesis(codec.GetProtoCodec()), Module.ExportGenesis(sdkTypes.UnwrapSDKContext(context), codec.GetProtoCodec()))
	// AppModuleSimulation
	require.Panics(t, func() {
		Module.GenerateGenesisState(&sdkModuleTypes.SimulationState{})
		Module.ProposalContents(sdkModuleTypes.SimulationState{})
		Module.RandomizedParams(&rand.Rand{})
		Module.RegisterStoreDecoder(sdkTypes.StoreDecoderRegistry{})
		Module.WeightedOperations(sdkModuleTypes.SimulationState{})
	})

	// types.Module
	require.Equal(t, "testAuxiliary", Module.GetAuxiliary("testAuxiliary").GetName())
	_, err = Module.DecodeModuleTransactionRequest("TestMessage", json.RawMessage(`{"BaseReq":{"from":"addr"},"ID":"id"}`))
	require.Nil(t, err)
}
