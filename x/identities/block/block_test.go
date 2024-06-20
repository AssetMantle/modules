// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"context"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	"testing"

	tendermintDB "github.com/cometbft/cometbft-db"
	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/mapper"
	"github.com/AssetMantle/modules/x/identities/parameters"
)

func CreateTestInput(t *testing.T) context.Context {
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

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return sdkTypes.WrapSDKContext(context)
}

func Test_block_Begin(t *testing.T) {
	ctx := CreateTestInput(t)
	type fields struct {
		mapper           helpers.Mapper
		parameterManager helpers.ParameterManager
	}
	type args struct {
		in0 context.Context
		in1 abciTypes.RequestBeginBlock
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{

		{"+ve", fields{mapper: mapper.Prototype(), parameterManager: parameters.Prototype()}, args{in0: ctx, in1: abciTypes.RequestBeginBlock{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := block{
				mapper:           tt.fields.mapper,
				parameterManager: tt.fields.parameterManager,
			}
			block.Begin(tt.args.in0, tt.args.in1)
		})
	}
}

func Test_block_End(t *testing.T) {
	type fields struct {
		mapper           helpers.Mapper
		parameterManager helpers.ParameterManager
	}
	type args struct {
		in0 context.Context
		in1 abciTypes.RequestEndBlock
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{

		{"+ve", fields{mapper.Prototype(), parameters.Prototype()}, args{CreateTestInput(t), abciTypes.RequestEndBlock{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := block{
				mapper:           tt.fields.mapper,
				parameterManager: tt.fields.parameterManager,
			}
			block.End(tt.args.in0, tt.args.in1)
		})
	}
}

func Test_block_Initialize(t *testing.T) {
	block := Prototype()
	block.Initialize(mapper.Prototype(), parameters.Prototype(), []helpers.Auxiliary{})
}
