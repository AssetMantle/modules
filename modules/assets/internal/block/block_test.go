// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"fmt"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/assets/internal/mapper"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
)

func CreateAssetsTestInput(t *testing.T) sdkTypes.Context {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return context
}

func Test_block_Begin(t *testing.T) {
	context := CreateAssetsTestInput(t)
	type fields struct {
		mapper     helpers.Mapper
		parameters helpers.Parameters
	}
	type args struct {
		in0 sdkTypes.Context
		in1 abciTypes.RequestBeginBlock
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{

		{"+ve", fields{mapper.Prototype(), parameters.Prototype()}, args{context, abciTypes.RequestBeginBlock{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := block{
				mapper:     tt.fields.mapper,
				parameters: tt.fields.parameters,
			}
			block.Begin(tt.args.in0, tt.args.in1)
		})
	}
}

func Test_block_End(t *testing.T) {
	context := CreateAssetsTestInput(t)

	type fields struct {
		mapper     helpers.Mapper
		parameters helpers.Parameters
	}
	type args struct {
		in0 sdkTypes.Context
		in1 abciTypes.RequestEndBlock
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{

		{"+ve", fields{mapper.Prototype(), parameters.Prototype()}, args{context, abciTypes.RequestEndBlock{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := block{
				mapper:     tt.fields.mapper,
				parameters: tt.fields.parameters,
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
		mapper     helpers.Mapper
		parameters helpers.Parameters
	}
	type args struct {
		mapper     helpers.Mapper
		parameters helpers.Parameters
		in2        []interface{}
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
				mapper:     tt.fields.mapper,
				parameters: tt.fields.parameters,
			}
			if got := block.Initialize(tt.args.mapper, tt.args.parameters, tt.args.in2...); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
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
