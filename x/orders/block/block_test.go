// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	storeTypes "cosmossdk.io/store/types"
	"testing"

	cosmosDB "github.com/cosmos/cosmos-db"
	"cosmossdk.io/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"cosmossdk.io/store"
	storeMetrics "cosmossdk.io/store/metrics"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/mapper"
	"github.com/AssetMantle/modules/x/orders/parameters"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

func CreateTestInput(t *testing.T) (sdkTypes.Context, helpers.Mapper, helpers.Auxiliary, helpers.Auxiliary, helpers.Auxiliary) {

	storeKey := storeTypes.NewKVStoreKey("test")
	paramsStoreKey := storeTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := storeTypes.NewTransientStoreKey("testParamsTransient")

	memDB := cosmosDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB, log.NewNopLogger(), storeMetrics.NewNoOpMetrics())
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	Mapper := mapper.Prototype().Initialize(storeKey)

	parameterManager := parameters.Prototype().Initialize(storeKey)
	transferAuxiliary := transfer.Auxiliary.Initialize(Mapper, parameterManager)
	supplementAuxiliary := supplement.Auxiliary.Initialize(Mapper, parameterManager)
	scrubAuxiliary := scrub.Auxiliary.Initialize(Mapper, parameterManager)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
		Height:  1000,
	}, false, log.NewNopLogger())

	return context, Mapper, transferAuxiliary, supplementAuxiliary, scrubAuxiliary
}

func Test_Block_Methods(t *testing.T) {
	block := Prototype()
	context, Mapper, transferAuxiliary, supplementAuxiliary, _ := CreateTestInput(t)
	block = block.Initialize(Mapper, parameters.Prototype(), transferAuxiliary, supplementAuxiliary)
	block.Begin(sdkTypes.WrapSDKContext(context))

	block.End(sdkTypes.WrapSDKContext(context))
}

func Test_block_End(t *testing.T) {
	context, Mapper, transferAuxiliary, supplementAuxiliary, scrubAuxiliary := CreateTestInput(t)
	_ = context.WithBlockHeight(1)
	_ = context.WithBlockHeight(-1)
	type fields struct {
		mapper              helpers.Mapper
		parameterManager    helpers.ParameterManager
		supplementAuxiliary helpers.Auxiliary
		transferAuxiliary   helpers.Auxiliary
		scrubAuxiliary      helpers.Auxiliary
	}
	type args struct {
		context sdkTypes.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"-ve without block height", fields{Mapper, parameters.Prototype(), supplementAuxiliary, transferAuxiliary, scrubAuxiliary}, args{context}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := block{
				mapper: tt.fields.mapper,
			}
			block.End(sdkTypes.WrapSDKContext(tt.args.context))
		})
	}
}
