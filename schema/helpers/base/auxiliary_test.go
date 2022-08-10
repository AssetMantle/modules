package base

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
	"reflect"
	"testing"
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
		// TODO: Add test cases.
		// TODO get it verified as it's failing
		{"+ve", args{"testAuxiliary", base.TestAuxiliaryKeeperPrototype}, auxiliary{name: "testAuxiliary", keeperPrototype: base.TestAuxiliaryKeeperPrototype}},
		{"nil", args{"nil", nil}, auxiliary{"nil", nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliary(tt.args.name, tt.args.keeperPrototype); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuxiliary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliary_GetKeeper(t *testing.T) {
	context, _, _ := base.SetupTest(t)
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
		// TODO: Add test cases.
		{"+ve", fields{"testAuxiliary", base.TestAuxiliaryKeeperPrototype(), base.TestAuxiliaryKeeperPrototype}, auxiliary{name: "testAuxiliary", keeperPrototype: base.TestAuxiliaryKeeperPrototype}.auxiliaryKeeper},
		{"nil", fields{"nil", nil, nil}, auxiliary{}.GetKeeper()},
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
		// TODO: Add test cases.
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
	_, storeKey, _ := base.SetupTest(t)

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
		parameters       helpers.Parameters
		auxiliaryKeepers []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Auxiliary
	}{
		// TODO: Add test cases.
		{"+ve", fields{"testAuxiliary", base.TestAuxiliaryKeeperPrototype(), base.TestAuxiliaryKeeperPrototype}, args{mapper: Mapper, parameters: nil}, Auxiliary},
		//{"nil", fields{"", nil, nil}, args{mapper: Mapper, parameters: nil, auxiliaryKeepers: nil}, auxiliary{"", nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliary := auxiliary{
				name:            tt.fields.name,
				auxiliaryKeeper: tt.fields.auxiliaryKeeper,
				keeperPrototype: tt.fields.keeperPrototype,
			}
			if got := auxiliary.Initialize(tt.args.mapper, tt.args.parameters, tt.args.auxiliaryKeepers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
