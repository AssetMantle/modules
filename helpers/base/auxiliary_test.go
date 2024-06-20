// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	context2 "context"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"

	"github.com/AssetMantle/modules/helpers"
)

type dummyAuxiliaryKeeper struct{}

func (dummyAuxiliaryKeeper *dummyAuxiliaryKeeper) Initialize(_ helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return dummyAuxiliaryKeeper
}
func (dummyAuxiliaryKeeper *dummyAuxiliaryKeeper) Help(_ context.Context, _ helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	return nil, nil
}

type dummyMapper struct{}

func (dummyMapper dummyMapper) NewCollection(_ context2.Context) helpers.Collection {
	return nil
}
func (dummyMapper dummyMapper) Upsert(_ context2.Context, _ helpers.Record) {
}
func (dummyMapper dummyMapper) Read(_ context2.Context, _ helpers.Key) helpers.Record {
	return nil
}
func (dummyMapper dummyMapper) Delete(_ context2.Context, _ helpers.Key) {
}
func (dummyMapper dummyMapper) FetchAll(_ context2.Context) []helpers.Record {
	return nil
}
func (dummyMapper dummyMapper) Iterate(_ context2.Context, _ helpers.Key, _ func(helpers.Record) bool) {
}

func (dummyMapper dummyMapper) IterateAll(_ context2.Context, _ func(helpers.Record) bool) {
}

func (dummyMapper dummyMapper) IteratePaginated(_ context2.Context, _ helpers.Key, _ int32, _ func(helpers.Record) bool) {
}
func (dummyMapper dummyMapper) StoreDecoder(_ kv.Pair, _ kv.Pair) string {
	return ""
}
func (dummyMapper dummyMapper) Initialize(_ *storeTypes.KVStoreKey) helpers.Mapper {
	return nil
}
func (dummyMapper dummyMapper) Get(_ string) (interface{}, error) {
	return nil, nil
}
func (dummyMapper dummyMapper) Set(_ string, _ interface{}) error {
	return nil
}

func TestGetName(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		expect string
		panic  bool
	}{
		{
			name:   "Valid name",
			input:  "sampleName",
			expect: "sampleName",
		},
		{
			name:  "Empty name",
			input: "",
			panic: true,
		},
		{
			name:   "Whitespace name",
			input:  "   ",
			expect: "   ",
		},
		{
			name:   "Special character in name",
			input:  "@#$-_=+!",
			expect: "@#$-_=+!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {

			}
			if tt.panic {
				assert.Panics(t, func() { NewAuxiliary(tt.input, func() helpers.AuxiliaryKeeper { return nil }) })
			} else {
				if NewAuxiliary(tt.input, func() helpers.AuxiliaryKeeper { return nil }).GetName() != tt.expect {
					t.Errorf("GetName() = %v, want = %v", NewAuxiliary(tt.input, func() helpers.AuxiliaryKeeper { return nil }).GetName(), tt.expect)
				}
			}
		})
	}
}

func TestGetKeeper(t *testing.T) {
	tests := []struct {
		name        string
		auxiliary   auxiliary
		expectedObj helpers.AuxiliaryKeeper
	}{
		{
			name: "Normal case",
			auxiliary: auxiliary{
				auxiliaryKeeper: &dummyAuxiliaryKeeper{},
				name:            "testCase1",
			},
			expectedObj: &dummyAuxiliaryKeeper{},
		},
		{
			name: "Nil case",
			auxiliary: auxiliary{
				auxiliaryKeeper: nil,
				name:            "testCase2",
			},
			expectedObj: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKeeper := tt.auxiliary.GetKeeper(); gotKeeper != tt.expectedObj {
				t.Errorf("GetKeeper() = %v, want %v", gotKeeper, tt.expectedObj)
			}
		})
	}
}

func Test_auxiliary_Initialize(t *testing.T) {
	type fields struct {
		name            string
		auxiliaryKeeper helpers.AuxiliaryKeeper
		keeperPrototype func() helpers.AuxiliaryKeeper
	}
	type args struct {
		mapper           helpers.Mapper
		parameterManager helpers.ParameterManager
		auxiliaryKeepers []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Auxiliary
	}{
		{
			name: "Test valid input",
			fields: fields{
				name:            "testName",
				auxiliaryKeeper: &dummyAuxiliaryKeeper{},
				keeperPrototype: func() helpers.AuxiliaryKeeper { return &dummyAuxiliaryKeeper{} },
			},
			args: args{
				mapper:           dummyMapper{},
				parameterManager: nil,
				auxiliaryKeepers: []interface{}{&dummyAuxiliaryKeeper{}},
			},
			want: auxiliary{
				name:            "testName",
				auxiliaryKeeper: &dummyAuxiliaryKeeper{},
				keeperPrototype: func() helpers.AuxiliaryKeeper { return &dummyAuxiliaryKeeper{} },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliary := auxiliary{
				name:            tt.fields.name,
				auxiliaryKeeper: tt.fields.auxiliaryKeeper,
				keeperPrototype: tt.fields.keeperPrototype,
			}
			var auxiliaryKeepers []interface{}
			for _, v := range tt.args.auxiliaryKeepers {
				auxiliaryKeepers = append(auxiliaryKeepers, v)
			}
			assert.Equalf(t, tt.want.GetKeeper(), auxiliary.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaryKeepers...).GetKeeper(), "Initialize(%v, %v, %v)", tt.args.mapper, tt.args.parameterManager, auxiliaryKeepers)
			assert.Equalf(t, tt.want.GetName(), auxiliary.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaryKeepers...).GetName(), "Initialize(%v, %v, %v)", tt.args.mapper, tt.args.parameterManager, auxiliaryKeepers)
		})
	}
}

func TestNewAuxiliary(t *testing.T) {
	type args struct {
		name            string
		keeperPrototype func() helpers.AuxiliaryKeeper
	}
	tests := []struct {
		name   string
		args   args
		want   helpers.Auxiliary
		panics bool
	}{
		{
			name: "Test valid input",
			args: args{
				name:            "testName",
				keeperPrototype: func() helpers.AuxiliaryKeeper { return &dummyAuxiliaryKeeper{} },
			},
			want: auxiliary{
				name:            "testName",
				keeperPrototype: func() helpers.AuxiliaryKeeper { return &dummyAuxiliaryKeeper{} },
			},
		},
		{
			name: "Test empty name",
			args: args{
				name:            "",
				keeperPrototype: func() helpers.AuxiliaryKeeper { return &dummyAuxiliaryKeeper{} },
			},
			panics: true,
		},
		{
			name: "Test nil keeper prototype",
			args: args{
				name:            "testName",
				keeperPrototype: nil,
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { NewAuxiliary(tt.args.name, tt.args.keeperPrototype) })
			} else {
				assert.Equalf(t, tt.want.GetName(), NewAuxiliary(tt.args.name, tt.args.keeperPrototype).GetName(), "NewAuxiliary(%v, %v)", tt.args.name, tt.args.keeperPrototype)
				assert.Equalf(t, tt.want.GetKeeper(), NewAuxiliary(tt.args.name, tt.args.keeperPrototype).GetKeeper(), "NewAuxiliary(%v, %v)", tt.args.name, tt.args.keeperPrototype)
			}
		})
	}
}
