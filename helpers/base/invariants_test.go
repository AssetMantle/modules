package base

import (
	"github.com/AssetMantle/modules/helpers"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInvariants(t *testing.T) {
	type args struct {
		moduleName    string
		route         string
		invariantList []sdkTypes.Invariant
	}
	tests := []struct {
		name string
		args args
		want helpers.Invariants
	}{
		{
			name: "Test NewInvariants",
			args: args{
				moduleName:    "test",
				route:         "test",
				invariantList: []sdkTypes.Invariant{},
			},
			want: invariants{
				moduleName:    "test",
				route:         "test",
				invariantList: []sdkTypes.Invariant{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var interfaceList []interface{}
			for _, invariant := range tt.args.invariantList {
				interfaceList = append(interfaceList, invariant)
			}
			assert.Equalf(t, tt.want, NewInvariants(tt.args.moduleName, tt.args.route, tt.args.invariantList...), "NewInvariants(%v, %v, %v)", tt.args.moduleName, tt.args.route, interfaceList)
		})
	}
}

func Test_invariants_Register(t *testing.T) {
	type fields struct {
		moduleName    string
		route         string
		invariantList []sdkTypes.Invariant
	}
	type args struct {
		invariantRegistry sdkTypes.InvariantRegistry
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test invariants Register",
			fields: fields{
				moduleName:    "test",
				route:         "test",
				invariantList: []sdkTypes.Invariant{},
			},
			args: args{
				invariantRegistry: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invariants := invariants{
				moduleName:    tt.fields.moduleName,
				route:         tt.fields.route,
				invariantList: tt.fields.invariantList,
			}
			invariants.Register(tt.args.invariantRegistry)
		})
	}
}
