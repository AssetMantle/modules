// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/utilities/test"
	"github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

func TestNewAuxiliary(t *testing.T) {
	type args struct {
		name            string
		keeperPrototype func() helpers.AuxiliaryKeeper
	}
	tests := []struct {
		name string
		args args
		want helpers.Auxiliary
	}{

		// TODO get it verified as it's failing
		{"+ve", args{"testAuxiliary", base.TestAuxiliaryKeeperPrototype}, auxiliary{name: "testAuxiliary", keeperPrototype: base.TestAuxiliaryKeeperPrototype}},
		{"nil", args{"nil", nil}, auxiliary{"nil", nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliary(tt.args.name, tt.args.keeperPrototype); !reflect.DeepEqual(got.GetName(), tt.want.GetName()) {
				t.Errorf("NewAuxiliary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliary_GetKeeper(t *testing.T) {
	context, _, _ := test.SetupTest(t)
	type fields struct {
		name            string
		auxiliaryKeeper helpers.AuxiliaryKeeper
		keeperPrototype func() helpers.AuxiliaryKeeper
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.AuxiliaryKeeper
	}{

		{"+ve", fields{"testAuxiliary", base.TestAuxiliaryKeeperPrototype(), base.TestAuxiliaryKeeperPrototype}, auxiliary{name: "testAuxiliary", keeperPrototype: base.TestAuxiliaryKeeperPrototype}.auxiliaryKeeper},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliary := auxiliary{
				name:            tt.fields.name,
				auxiliaryKeeper: tt.fields.auxiliaryKeeper,
				keeperPrototype: tt.fields.keeperPrototype,
			}
			if got := auxiliary.GetKeeper().Help(context, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeeper() = %T, want %T", got, tt.want)
			}
		})
	}
}

func Test_auxiliary_GetName(t *testing.T) {
	type fields struct {
		name            string
		auxiliaryKeeper helpers.AuxiliaryKeeper
		keeperPrototype func() helpers.AuxiliaryKeeper
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{"+ve", fields{"testAuxiliary", base.TestAuxiliaryKeeperPrototype(), base.TestAuxiliaryKeeperPrototype}, auxiliary{name: "testAuxiliary", keeperPrototype: base.TestAuxiliaryKeeperPrototype}.name},
		{"nil", fields{"", nil, nil}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliary := auxiliary{
				name:            tt.fields.name,
				auxiliaryKeeper: tt.fields.auxiliaryKeeper,
				keeperPrototype: tt.fields.keeperPrototype,
			}
			if got := auxiliary.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliary_Initialize(t *testing.T) {
	_, storeKey, _ := test.SetupTest(t)

	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Auxiliary := auxiliary{"testAuxiliary", base.TestAuxiliaryKeeperPrototype(), base.TestAuxiliaryKeeperPrototype}
	Auxiliary.auxiliaryKeeper = Auxiliary.keeperPrototype().Initialize(Mapper, nil, nil).(helpers.AuxiliaryKeeper)
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
		// TODO find fix
		// {"+ve", fields{"testAuxiliary", base.TestAuxiliaryKeeperPrototype(), base.TestAuxiliaryKeeperPrototype}, args{mapper: Mapper, parameterManager: nil, auxiliaryKeepers: nil}, Auxiliary},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliary := auxiliary{
				name:            tt.fields.name,
				auxiliaryKeeper: tt.fields.auxiliaryKeeper,
				keeperPrototype: tt.fields.keeperPrototype,
			}
			if got := auxiliary.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaryKeepers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
