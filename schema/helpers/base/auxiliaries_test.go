package base

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
	"reflect"
	"testing"
)

func TestNewAuxiliaries(t *testing.T) {
	_, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Auxilaries := NewAuxiliaries(NewAuxiliary("testAuxilary", base.TestAuxiliaryKeeperPrototype).Initialize(Mapper, nil))
	type args struct {
		auxiliaryList []helpers.Auxiliary
	}
	tests := []struct {
		name string
		args args
		want helpers.Auxiliaries
	}{
		// TODO: Add test cases.
		{"nil", args{nil}, auxiliaries{nil}},
		{"+ve", args{Auxilaries.GetList()}, Auxilaries.(auxiliaries)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliaries(tt.args.auxiliaryList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuxiliaries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaries_Get(t *testing.T) {
	_, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Auxilaries := NewAuxiliaries(NewAuxiliary("testAuxilary", base.TestAuxiliaryKeeperPrototype).Initialize(Mapper, nil))
	type fields struct {
		auxiliaryList []helpers.Auxiliary
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Auxiliary
	}{
		// TODO: Add test cases.
		{"+ve", fields{Auxilaries.GetList()}, args{"testAuxilary"}, Auxilaries.GetList()[0]},
		{"nil", fields{nil}, args{""}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaries := auxiliaries{
				auxiliaryList: tt.fields.auxiliaryList,
			}
			if got := auxiliaries.Get(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaries_GetList(t *testing.T) {
	_, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	type fields struct {
		auxiliaryList []helpers.Auxiliary
	}
	tests := []struct {
		name   string
		fields fields
		want   []helpers.Auxiliary
	}{
		// TODO: Add more test cases.
		{"+ve", fields{[]helpers.Auxiliary{NewAuxiliary("testAuxiliary", base.TestAuxiliaryKeeperPrototype).Initialize(Mapper, nil)}}, []helpers.Auxiliary{NewAuxiliary("testAuxiliary", base.TestAuxiliaryKeeperPrototype).Initialize(Mapper, nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaries := auxiliaries{
				auxiliaryList: tt.fields.auxiliaryList,
			}
			if got := auxiliaries.GetList(); !reflect.DeepEqual(got[0].GetName(), tt.want[0].GetName()) {
				t.Errorf("GetList() = %v, want %v", got[0].GetName(), tt.want[0].GetName())
			}
		})
	}

}
