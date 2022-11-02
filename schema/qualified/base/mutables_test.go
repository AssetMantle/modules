// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
)

func TestNewMutables(t *testing.T) {
	testMutableProperties := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))
	type args struct {
		propertyList lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want qualified.Mutables
	}{
		// TODO: Add test cases.
		{"+ve", args{}, mutables{}},
		{"+ve", args{base.NewPropertyList(testMutableProperties)}, mutables{base.NewPropertyList(testMutableProperties)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMutables(tt.args.propertyList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMutables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutables_GetMutablePropertyList(t *testing.T) {
	testMutableProperties := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))
	type fields struct {
		PropertyList lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.PropertyList
	}{
		// TODO: Add test cases.
		{"+ve", fields{}, base.NewPropertyList()},
		{"+ve", fields{base.NewPropertyList(testMutableProperties)}, mutables{base.NewPropertyList(testMutableProperties)}.MesaPropertyList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutables := mutables{
				MesaPropertyList: tt.fields.PropertyList,
			}
			if got := mutables.GetMutablePropertyList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMutablePropertyList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutables_Mutate(t *testing.T) {
	testMutableProperties := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))
	testMutableProperties1 := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("MutableData"))

	type fields struct {
		PropertyList lists.PropertyList
	}
	type args struct {
		propertyList []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   qualified.Mutables
	}{
		// TODO: Add test cases.
		{"+ve", fields{}, args{}, mutables{}},
		{"+ve mutate", fields{base.NewPropertyList(testMutableProperties)}, args{[]properties.Property{testMutableProperties1}}, mutables{base.NewPropertyList(testMutableProperties).Mutate(testMutableProperties1)}}, // TODO: it seems incorrect, not failing for wrong check
		{"+ve nil mutate", fields{base.NewPropertyList(testMutableProperties)}, args{}, mutables{base.NewPropertyList(testMutableProperties)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutables := mutables{
				MesaPropertyList: tt.fields.PropertyList,
			}
			if got := mutables.Mutate(tt.args.propertyList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}
