// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/traits"
	"reflect"
	"testing"
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
		// TODO: Add test cases.
		{"+ve", args{nil}, propertyList{List: NewList(propertiesToListables([]properties.Property{}...)...)}},
		{"+ve", args{[]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}}, propertyList{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}},
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
		// TODO: Add test cases.
		{"+ve", args{nil}, propertiesToListables()},
		{"+ve", args{[]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}}, propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)},
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
		List lists.List
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
		// TODO: Add test cases.
		{"+ve with no property addition", fields{NewList()}, args{nil}, propertyList{List: NewList()}},
		{"+ve nil with property addition", fields{NewList(propertiesToListables([]properties.Property{}...)...)}, args{[]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}}, propertyList{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA")), base.NewProperty(baseIDs.NewID("b"), NewStringData("factB"))}...)...)}},                                  //TODO: Not being added
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, args{[]properties.Property{base.NewProperty(baseIDs.NewID("b"), NewStringData("factB"))}}, propertyList{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA")), base.NewProperty(baseIDs.NewID("b"), NewStringData("factB"))}...)...)}}, //TODO: Not being added
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.Add(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_GetList(t *testing.T) {
	type fields struct {
		List lists.List
	}
	tests := []struct {
		name   string
		fields fields
		want   []properties.Property
	}{
		// TODO: Add test cases.
		{"+ve", fields{NewList()}, []properties.Property{}}, //TODO: Type & value same but not matching
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, []properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_GetProperty(t *testing.T) {
	type fields struct {
		List lists.List
	}
	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   properties.Property
	}{
		// TODO: Add test cases.
		{"-ve", fields{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, args{base.NewProperty(baseIDs.NewID("b"), NewStringData("factB")).GetID()}, nil},
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, args{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA")).GetID()}, propertyList{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}.GetList()[0]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.GetProperty(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Mutate(t *testing.T) {
	type fields struct {
		List lists.List
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
		// TODO: Add test cases.
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, args{nil}, propertyList{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}}, //TODO: Type & value same but not matching
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, args{[]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factB"))}}, propertyList{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factB"))}...)...)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.Mutate(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propertyList_Remove(t *testing.T) {
	type fields struct {
		List lists.List
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
		// TODO: Add test cases.
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, args{nil}, propertyList{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}}, //TODO: Type & value same but not matching
		{"+ve", fields{NewList()}, args{[]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}}, propertyList{NewList()}},
		{"+ve", fields{NewList(propertiesToListables([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA")), base.NewProperty(baseIDs.NewID("b"), NewStringData("factB"))}...)...)}, args{[]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}}, propertyList{NewList()}}, // TODO: raise issue
		// TODO: raise issue
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			propertyList := propertyList{
				List: tt.fields.List,
			}
			if got := propertyList.Remove(tt.args.properties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
