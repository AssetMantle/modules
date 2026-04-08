// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"context"
	storeTypes "cosmossdk.io/store/types"
	"testing"

	cosmosDB "github.com/cosmos/cosmos-db"
	"cosmossdk.io/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"cosmossdk.io/store"
	storeMetrics "cosmossdk.io/store/metrics"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
)

func CreateAssetsTestInput(t *testing.T) context.Context {
	storeKey := storeTypes.NewKVStoreKey("test")
	paramsStoreKey := storeTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := storeTypes.NewTransientStoreKey("testParamsTransient")

	memDB := cosmosDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB, log.NewNopLogger(), storeMetrics.NewNoOpMetrics())
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, nil)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, nil)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return sdkTypes.WrapSDKContext(context)
}

func Test_block_Begin(t *testing.T) {
	type fields struct {
		mapper           helpers.Mapper
		parameterManager helpers.ParameterManager
	}
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{

		{"valid", fields{mapper.Prototype(), parameters.Prototype()}, args{CreateAssetsTestInput(t)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := block{
				mapper:           tt.fields.mapper,
				parameterManager: tt.fields.parameterManager,
			}
			block.Begin(tt.args.in0)
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
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{

		{"valid", fields{mapper.Prototype(), parameters.Prototype()}, args{CreateAssetsTestInput(t)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := block{
				mapper:           tt.fields.mapper,
				parameterManager: tt.fields.parameterManager,
			}
			block.End(tt.args.in0)
		})
	}
}

func Test_block_Initialize(t *testing.T) {
	testMapper := mapper.Prototype()
	testParameter := parameters.Prototype()
	testBlock := block{testMapper, testParameter}
	type fields struct {
		mapper           helpers.Mapper
		parameterManager helpers.ParameterManager
	}
	type args struct {
		mapper           helpers.Mapper
		parameterManager helpers.ParameterManager
		in2              []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Block
	}{
		{"valid", fields{testMapper, testParameter}, args{testMapper, testParameter, []interface{}{}}, testBlock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := block{
				mapper:           tt.fields.mapper,
				parameterManager: tt.fields.parameterManager,
			}
			got := block.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.in2...)
			assert.NotNil(t, got)
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Block
	}{

		{"valid", block{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prototype()
			assert.Equal(t, tt.want, got, "Prototype()")
		})
	}
}
