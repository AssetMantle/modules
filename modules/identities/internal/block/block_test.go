// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/identities/internal/mapper"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
)

func CreateTestInput(t *testing.T) sdkTypes.Context {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
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

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return context
}

func Test_block_Begin(t *testing.T) {
	context := CreateTestInput(t)
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

		{"+ve", fields{mapper: mapper.Prototype(), parameters: parameters.Prototype()}, args{in0: context, in1: abciTypes.RequestBeginBlock{}}},
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

		{"+ve", fields{mapper.Prototype(), parameters.Prototype()}, args{CreateTestInput(t), abciTypes.RequestEndBlock{}}},
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
	block := Prototype()
	block.Initialize(mapper.Prototype(), parameters.Prototype(), []helpers.Auxiliary{})
}
