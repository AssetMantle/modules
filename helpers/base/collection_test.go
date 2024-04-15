// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockRecord struct {
	key helpers.Key
}

func (mr MockRecord) GetKey() helpers.Key {
	return mr.key
}

func (mr MockRecord) GetMappable() helpers.Mappable {
	return nil
}

func (mr MockRecord) WithKey(helpers.Key) helpers.Record {
	return nil
}

func (mr MockRecord) ReadFromIterator(sdkTypes.Iterator) helpers.Record {
	return nil
}

func (mr MockRecord) Read(types.KVStore) helpers.Record {
	return nil
}

func (mr MockRecord) Write(types.KVStore) helpers.Record {
	return nil
}

func (mr MockRecord) Delete(types.KVStore) {
}

type MockMapper struct {
}

func (mm MockMapper) NewCollection(_ context.Context) helpers.Collection {
	return nil
}

func (mm MockMapper) StoreDecoder(_ kv.Pair, _ kv.Pair) string {
	return ""
}

func (mm MockMapper) Initialize(_ *sdkTypes.KVStoreKey) helpers.Mapper {
	return nil
}

func (mm MockMapper) Read(context.Context, helpers.Key) helpers.Record {
	return nil
}

func (mm MockMapper) Upsert(context.Context, helpers.Record) {
}

func (mm MockMapper) IterateAll(context.Context, func(helpers.Record) bool) {
}

func (mm MockMapper) FetchAll(context.Context) []helpers.Record {
	return nil
}

func (mm MockMapper) Delete(context.Context, helpers.Key) {
}

func (mm MockMapper) Iterate(context.Context, helpers.Key, func(helpers.Record) bool) {
}

func (mm MockMapper) IteratePaginated(context.Context, helpers.Key, int32, func(helpers.Record) bool) {
}

func TestMutate(t *testing.T) {
	tt := []struct {
		name     string
		mapper   helpers.Mapper
		input    helpers.Record
		expected helpers.Collection
	}{
		{
			name:     "Empty records",
			mapper:   MockMapper{},
			input:    MockRecord{},
			expected: collection{records: []helpers.Record{}, mapper: MockMapper{}, context: nil},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := collection{
				records: []helpers.Record{},
				mapper:  tc.mapper,
				context: nil,
			}

			output := c.Mutate(tc.input)

			for i := range output.Get() {
				if output.Get()[i].GetKey() != tc.expected.Get()[i].GetKey() {
					t.Errorf("Record key not match. Got %v, wants %v", output.Get()[i].GetKey(), tc.expected.Get()[i].GetKey())
				}
			}
		})
	}
}

func Test_collection_Initialize(t *testing.T) {
	type fields struct {
		records []helpers.Record
		mapper  helpers.Mapper
		context context.Context
	}
	type args struct {
		context context.Context
		mapper  helpers.Mapper
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Collection
	}{
		{
			name: "Test collection Initialize",
			fields: fields{
				records: []helpers.Record{},
				mapper:  nil,
				context: context.Background(),
			},
			args: args{
				context: context.Background(),
				mapper:  nil,
			},
			want: collection{
				records: []helpers.Record{},
				mapper:  nil,
				context: context.Background(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collection := collection{
				records: tt.fields.records,
				mapper:  tt.fields.mapper,
				context: tt.fields.context,
			}
			assert.Equalf(t, tt.want, collection.Initialize(tt.args.context, tt.args.mapper), "Initialize(%v, %v)", tt.args.context, tt.args.mapper)
		})
	}
}
