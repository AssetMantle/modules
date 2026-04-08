// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"context"
	storeTypes "cosmossdk.io/store/types"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/metas/mapper"
	"github.com/AssetMantle/modules/x/metas/parameters"
	"github.com/AssetMantle/modules/x/metas/record"
	"github.com/AssetMantle/schema/data/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestKeepers struct {
	MetasKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (context.Context, TestKeepers) {

	storeKey := storeTypes.NewKVStoreKey("test")
	Context := testutil.NewTestContext(t, storeKey)

	Mapper := mapper.Prototype().Initialize(storeKey)
	parameterManager := parameters.Prototype().Initialize(storeKey)
	parameterManager, _ = parameterManager.Set().Update(sdkTypes.WrapSDKContext(Context))

	keepers := TestKeepers{
		MetasKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.TransactionKeeper),
	}

	return sdkTypes.WrapSDKContext(Context), keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	Context, keepers := CreateTestInput(t)
	defaultAddr := sdkTypes.AccAddress("addr")
	data, err := base.PrototypeAnyData().FromString("S|default")
	require.Equal(t, nil, err)
	newFact, err := base.PrototypeAnyData().FromString("S|newFact")
	require.Equal(t, nil, err)
	keepers.MetasKeeper.(transactionKeeper).mapper.NewCollection(Context).Add(record.NewRecord(data))
	type fields struct {
		mapper           helpers.Mapper
		parameterManager helpers.ParameterManager
	}
	type args struct {
		context context.Context
		message helpers.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionResponse
		wantErr bool
	}{
		{"valid", fields{keepers.MetasKeeper.(transactionKeeper).mapper, keepers.MetasKeeper.(transactionKeeper).parameterManager}, args{Context, NewMessage(defaultAddr, newFact).(*Message)}, newTransactionResponse(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:           tt.fields.mapper,
				parameterManager: tt.fields.parameterManager,
			}
			got, err := transactionKeeper.Transact(tt.args.context, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "Transact() got")
		})
	}
}

func Test_keeperPrototype(t *testing.T) {
	got := keeperPrototype()
	assert.Equal(t, transactionKeeper{}, got)
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	Mapper := mapper.Prototype().Initialize(storeKey)
	parameterManager := parameters.Prototype().Initialize(storeKey)

	keeper := keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{})
	require.NotNil(t, keeper)
}
