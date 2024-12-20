// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"context"
	"fmt"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	"reflect"
	"testing"

	tendermintDB "github.com/cometbft/cometbft-db"
	abciTypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
)

func CreateAssetsTestInput(t *testing.T) context.Context {
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

		{"+ve", fields{mapper.Prototype(), parameters.Prototype()}, args{CreateAssetsTestInput(t), abciTypes.RequestBeginBlock{}}},
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

		{"+ve", fields{mapper.Prototype(), parameters.Prototype()}, args{CreateAssetsTestInput(t), abciTypes.RequestEndBlock{}}},
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
		{"+ve", fields{testMapper, testParameter}, args{testMapper, testParameter, []interface{}{}}, testBlock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := block{
				mapper:           tt.fields.mapper,
				parameterManager: tt.fields.parameterManager,
			}
			if got := block.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.in2...); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Block
	}{

		{"+ve", block{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
