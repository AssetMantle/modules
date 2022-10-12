// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
	"reflect"
	"testing"
)

func TestNewImmutables(t *testing.T) {
	testImmutablePropertyList := base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("MutableData1")))
	type args struct {
		propertyList lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want qualified.Immutables
	}{
		// TODO: Add test cases.
		{"+ve", args{}, immutables{}},
		{"+ve", args{testImmutablePropertyList}, immutables{testImmutablePropertyList}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImmutables(tt.args.propertyList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImmutables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_immutables_GenerateHashID(t *testing.T) {
	testMesaProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("ImmutableData"))
	testImmutablePropertyList := base.NewPropertyList(testMesaProperty)
	metaList2 := make([][]byte, len(testImmutablePropertyList.GetList()))
	metaList2[0] = testMesaProperty.GetDataID().GetHashID().Bytes()
	type fields struct {
		PropertyList lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{
		// TODO: Add test cases.
		{"+ve", fields{base.NewPropertyList()}, baseIDs.GenerateHashID([][]byte{}...)},
		{"+ve", fields{testImmutablePropertyList}, baseIDs.GenerateHashID(metaList2...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testImmutables := immutables{
				PropertyList: tt.fields.PropertyList,
			}
			if got := testImmutables.GenerateHashID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHashID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_immutables_GetImmutablePropertyList(t *testing.T) {
	testImmutablePropertyList := base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("MutableData1")))
	type fields struct {
		PropertyList lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.PropertyList
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{base.NewPropertyList()}, base.NewPropertyList()},
		{"+ve", fields{testImmutablePropertyList}, immutables{testImmutablePropertyList}.PropertyList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testImmutables := immutables{
				PropertyList: tt.fields.PropertyList,
			}
			if got := testImmutables.GetImmutablePropertyList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImmutablePropertyList() = %v, want %v", got, tt.want)
			}
		})
	}
}
