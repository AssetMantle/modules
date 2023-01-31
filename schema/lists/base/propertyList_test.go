// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"reflect"
	"testing"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/traits"
)

func TestNewPropertyList(t *testing.T) {
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name string
		args args
		want lists.PropertyList
	}{
		{"+ve", args{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, &PropertyList{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")).ToAnyProperty().(*baseProperties.AnyProperty), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty)}}},
		//Duplicate property must not be added
		{"-ve duplicate supplied", args{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))}}, &PropertyList{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")).ToAnyProperty().(*baseProperties.AnyProperty)}}}, // TODO: Should fail for duplicate Property
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPropertyList(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPropertyList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertiesToListables(t *testing.T) {
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name string
		args args
		want []traits.Listable
	}{
		{"+ve", args{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, []traits.Listable{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := propertiesToListables(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("propertiesToListables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Add(t *testing.T) {
	type fields struct {
		List []*baseProperties.AnyProperty
	}
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.PropertyList
	}{
		{"+ve", fields{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")).ToAnyProperty().(*baseProperties.AnyProperty)}}, args{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}...)},
		{"-ve", fields{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")).ToAnyProperty().(*baseProperties.AnyProperty)}}, args{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewDecData(sdkTypes.NewDec(1)))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewDecData(sdkTypes.NewDec(1)))}...)}, // TODO: Should fail as add method should not be able to mutate/add property with existing key
		{"+ve with nil", fields{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")).ToAnyProperty().(*baseProperties.AnyProperty)}}, args{nil}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))}...)},                                                                                                                                                                                                                    // TODO: Should fail as add method should not be able to mutate/add property with existing key
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := &PropertyList{
				PropertyList: tt.fields.List,
			}
			if got := propertyList.Add(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_GetList(t *testing.T) {
	type fields struct {
		List []*baseProperties.AnyProperty
	}
	tests := []struct {
		name   string
		fields fields
		want   []properties.AnyProperty
	}{
		{"+ve", fields{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty)}}, []properties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := &PropertyList{
				PropertyList: tt.fields.List,
			}
			if got := propertyList.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_GetProperty(t *testing.T) {
	type fields struct {
		List []*baseProperties.AnyProperty
	}
	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    properties.AnyProperty
		wantErr bool
	}{
		{"+ve Meta", fields{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty)}}, args{baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D"))}, baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty(), false},
		{"+ve Mesa", fields{[]*baseProperties.AnyProperty{baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty)}}, args{baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D"))}, baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty(), false},
		{"panic nil propertyID", fields{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty)}}, args{propertyID: nil}, baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty(), true}, // TODO: panics if propertyID is nil
		{"-ve", fields{[]*baseProperties.AnyProperty{baseProperties.NewEmptyMetaPropertyFromID(baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D"))).ToAnyProperty().(*baseProperties.AnyProperty)}}, args{baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D"))}, baseProperties.NewEmptyMetaPropertyFromID(baseIDs.NewPropertyID(baseIDs.NewStringID("supply"), baseIDs.NewStringID("D"))).ToAnyProperty(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := &PropertyList{
				PropertyList: tt.fields.List,
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("&StringIDFromInterface() error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := propertyList.GetProperty(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_GetPropertyIDList(t *testing.T) {
	type fields struct {
		List []*baseProperties.AnyProperty
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.IDList
	}{
		{"+ve", fields{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty)}}, NewIDList().Add(baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).GetID())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := &PropertyList{
				PropertyList: tt.fields.List,
			}
			if got := propertyList.GetPropertyIDList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPropertyIDList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Mutate(t *testing.T) {
	type fields struct {
		List []*baseProperties.AnyProperty
	}
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.PropertyList
	}{
		{"+ve", fields{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty)}}, args{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(2)))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(2)))}...)},
		// {"+ve", fields{NewList(propertiesToListables([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}...)...)}, args{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), NewStringData("test"))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), NewStringData("test"))}...)}, //TODO: Should handle error if different property data is tried to mutate
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := &PropertyList{
				PropertyList: tt.fields.List,
			}
			if got := propertyList.Mutate(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Remove(t *testing.T) {
	type fields struct {
		List []*baseProperties.AnyProperty
	}
	type args struct {
		properties []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.PropertyList
	}{
		{"+ve", fields{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply1"), baseData.NewDecData(sdkTypes.NewDec(2))).ToAnyProperty().(*baseProperties.AnyProperty)}}, args{[]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply1"), baseData.NewDecData(sdkTypes.NewDec(2)))}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := &PropertyList{
				PropertyList: tt.fields.List,
			}
			if got := propertyList.Remove(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_ScrubData(t *testing.T) {
	type fields struct {
		List []*baseProperties.AnyProperty
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.PropertyList
	}{
		{"+ve", fields{[]*baseProperties.AnyProperty{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty)}}, NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ScrubData()}...)},
		{"+ve", fields{[]*baseProperties.AnyProperty{baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))).ToAnyProperty().(*baseProperties.AnyProperty), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Test")).ToAnyProperty().(*baseProperties.AnyProperty)}}, NewPropertyList([]properties.Property{baseProperties.NewMesaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1))), baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Test")).ScrubData()}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := &PropertyList{
				PropertyList: tt.fields.List,
			}
			if got := propertyList.ScrubData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScrubData() = %v, want %v", got, tt.want)
			}
		})
	}
}
