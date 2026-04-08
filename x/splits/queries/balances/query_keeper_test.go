// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package balances

import (
	storeTypes "cosmossdk.io/store/types"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/x/splits/mapper"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/splits/parameters"
)

type TestKeepers struct {
	QueryKeeper helpers.QueryKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

	storeKey := storeTypes.NewKVStoreKey("test")
	Context := testutil.NewTestContext(t, storeKey)

	Mapper := mapper.Prototype().Initialize(storeKey)
	parameterManager := parameters.Prototype().Initialize(storeKey)
	parameterManager, _ = parameterManager.Set().Update(sdkTypes.WrapSDKContext(Context))

	keepers := TestKeepers{
		QueryKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.QueryKeeper),
	}

	return Context, keepers, Mapper, parameterManager
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryKeeper
	}{
		{"valid", queryKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := keeperPrototype()
			assert.Equal(t, tt.want, got, "keeperPrototype()")
		})
	}
}

func Test_queryKeeper_Initialize(t *testing.T) {
	_, _, Mapper, parameterManager := createTestInput(t)
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
		{"valid", fields{Mapper}, args{Mapper, parameterManager, []interface{}{}}, queryKeeper{Mapper}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryKeeper := queryKeeper{
				mapper: tt.fields.mapper,
			}
			if got := queryKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.in2); got == nil {
				t.Errorf("Initialize() = nil, want non-nil")
			}
		})
	}
}

func Test_queryKeeper_Enquire(t *testing.T) {
	Context, keepers, _, _ := createTestInput(t)

	t.Run("nil identity panics in validate", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = keepers.QueryKeeper.Enquire(sdkTypes.WrapSDKContext(Context), newQueryRequest(baseIDs.PrototypeIdentityID()))
		})
	})
}
