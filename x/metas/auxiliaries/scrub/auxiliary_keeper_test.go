// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"context"
	storeTypes "cosmossdk.io/store/types"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/metas/mapper"
	"github.com/AssetMantle/modules/x/metas/parameters"
	"github.com/stretchr/testify/assert"
)

type TestKeepers struct {
	MetasKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {

	storeKey := storeTypes.NewKVStoreKey("test")
	Context := testutil.NewTestContext(t, storeKey)

	Mapper := mapper.Prototype().Initialize(storeKey)
	parameterManager := parameters.Prototype().Initialize(storeKey)

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
			assert.Equal(t, tt.want, got, "Help() got")
		})
	}
}
