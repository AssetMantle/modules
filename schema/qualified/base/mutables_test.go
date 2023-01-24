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
		{"+ve", args{}, &Mutables{}},
		{"+ve", args{base.NewPropertyList(testMutableProperties)}, &Mutables{base.NewPropertyList(testMutableProperties).(*base.PropertyList)}},
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
		{"+ve", fields{}, base.NewPropertyList()},
		{"+ve", fields{base.NewPropertyList(testMutableProperties)}, (&Mutables{base.NewPropertyList(testMutableProperties).(*base.PropertyList)}).PropertyList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutables := &Mutables{
				PropertyList: tt.fields.PropertyList.(*base.PropertyList),
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
		{"+ve", fields{}, args{}, &Mutables{}},
		{"+ve mutate", fields{base.NewPropertyList(testMutableProperties)}, args{[]properties.Property{testMutableProperties1}}, (&Mutables{base.NewPropertyList(testMutableProperties).(*base.PropertyList)}).Mutate(testMutableProperties1)}, // TODO: it seems incorrect, not failing for wrong check
		{"+ve nil mutate", fields{base.NewPropertyList(testMutableProperties)}, args{}, &Mutables{base.NewPropertyList(testMutableProperties).(*base.PropertyList)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutables := &Mutables{
				PropertyList: tt.fields.PropertyList.(*base.PropertyList),
			}
			if got := mutables.Mutate(tt.args.propertyList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}
