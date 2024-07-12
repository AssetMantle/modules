// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	"testing"

	tendermintDB "github.com/cometbft/cometbft-db"
	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/mapper"
	"github.com/AssetMantle/modules/x/orders/parameters"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

func CreateTestInput(t *testing.T) (sdkTypes.Context, helpers.Mapper, helpers.Auxiliary, helpers.Auxiliary, helpers.Auxiliary) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	Mapper := mapper.Prototype().Initialize(storeKey)
	codec := baseHelpers.TestCodec()
	ParamsKeeper := paramsKeeper.NewKeeper(
		codec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	parameterManager := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))
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
	block.Begin(sdkTypes.WrapSDKContext(context), abciTypes.RequestBeginBlock{})

	block.End(sdkTypes.WrapSDKContext(context), abciTypes.RequestEndBlock{})
}

func Test_block_End(t *testing.T) {
	context, Mapper, transferAuxiliary, supplementAuxiliary, scrubAuxiliary := CreateTestInput(t)
	testContext := context.WithBlockHeight(1)
	testContext1 := context.WithBlockHeight(-1)
	type fields struct {
		mapper              helpers.Mapper
		parameterManager    helpers.ParameterManager
		supplementAuxiliary helpers.Auxiliary
		transferAuxiliary   helpers.Auxiliary
		scrubAuxiliary      helpers.Auxiliary
	}
	type args struct {
		context sdkTypes.Context
		in1     abciTypes.RequestEndBlock
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve with block height", fields{Mapper, parameters.Prototype(), supplementAuxiliary, transferAuxiliary, scrubAuxiliary}, args{testContext, abciTypes.RequestEndBlock{Height: int64(1)}}},
		{"-ve without block height", fields{Mapper, parameters.Prototype(), supplementAuxiliary, transferAuxiliary, scrubAuxiliary}, args{context, abciTypes.RequestEndBlock{}}},
		{"-ve with -ve block height", fields{Mapper, parameters.Prototype(), supplementAuxiliary, transferAuxiliary, scrubAuxiliary}, args{testContext1, abciTypes.RequestEndBlock{Height: int64(-1)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := block{
				mapper: tt.fields.mapper,
			}
			block.End(sdkTypes.WrapSDKContext(tt.args.context), tt.args.in1)
		})
	}
}
