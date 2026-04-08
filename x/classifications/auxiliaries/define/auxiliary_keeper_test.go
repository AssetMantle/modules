// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	storeTypes "cosmossdk.io/store/types"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/classifications/mapper"
	"github.com/AssetMantle/modules/x/classifications/parameters"
	"github.com/stretchr/testify/assert"
)

type TestKeepers struct {
	ClassificationsKeeper helpers.AuxiliaryKeeper
}

func createTestInput(t *testing.T) (context.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

	storeKey := storeTypes.NewKVStoreKey("test")
	Context := testutil.NewTestContext(t, storeKey)

	Mapper := mapper.Prototype().Initialize(storeKey)
	parameterManager := parameters.Prototype().Initialize(storeKey)
	parameterManager, _ = parameterManager.Set().Update(sdkTypes.WrapSDKContext(Context))

	keepers := TestKeepers{
		ClassificationsKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{bankKeeper.BaseKeeper{}, &stakingKeeper.Keeper{}}).(helpers.AuxiliaryKeeper),
	}

	return sdkTypes.WrapSDKContext(Context), keepers, Mapper, parameterManager
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.AuxiliaryKeeper
	}{
		{"valid", auxiliaryKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := keeperPrototype()
			assert.Equal(t, tt.want, got, "keeperPrototype()")
		})
	}
}

func Test_auxiliaryKeeper_Initialize(t *testing.T) {
	_, _, Mapper, parameterManager := createTestInput(t)
	testStakingKeeper := &stakingKeeper.Keeper{}
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		mapper helpers.Mapper
		in1    helpers.ParameterManager
		in2    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"valid", fields{Mapper}, args{Mapper, parameterManager, []interface{}{bankKeeper.BaseKeeper{}, testStakingKeeper}}, auxiliaryKeeper{Mapper, parameterManager, bankKeeper.BaseKeeper{}, testStakingKeeper}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := auxiliaryKeeper{
				mapper: tt.fields.mapper,
			}
			if got := au.Initialize(tt.args.mapper, tt.args.in1, tt.args.in2); got == nil {
				t.Errorf("Initialize() = nil, want non-nil")
			}
		})
	}
}
