// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"context"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/metas/mapper"
	"github.com/AssetMantle/modules/x/metas/parameters"
)

type TestKeepers struct {
	MetasKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := mapper.Prototype().Initialize(storeKey)
	encodingConfig := simapp.MakeTestEncodingConfig()
	appCodec := encodingConfig.Marshaler
	ParamsKeeper := paramsKeeper.NewKeeper(
		appCodec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	parameterManager := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		MetasKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return Context, keepers

}

func Test_auxiliaryKeeper_Help(t *testing.T) {
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		context context.Context
		request helpers.AuxiliaryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.AuxiliaryResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryKeeper := auxiliaryKeeper{
				mapper: tt.fields.mapper,
			}
			got, err := auxiliaryKeeper.Help(tt.args.context, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Help() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Help() got = %v, want %v", got, tt.want)
			}
		})
	}
}
