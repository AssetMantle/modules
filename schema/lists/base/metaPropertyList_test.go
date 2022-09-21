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

func TestNewMetaProperties(t *testing.T) {
	type args struct {
		metaProperties []properties.MetaProperty
	}
	tests := []struct {
		name string
		args args
		want lists.MetaPropertyList
	}{
		// TODO: Add test cases.
		{"+ve with nil", args{[]properties.MetaProperty{}}, metaPropertyList{NewList(metaPropertiesToListables([]properties.MetaProperty{}...)...)}},
		{"+ve", args{[]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}}, metaPropertyList{NewList(metaPropertiesToListables([]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMetaProperties(tt.args.metaProperties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMetaProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaPropertiesToListables(t *testing.T) {
	type args struct {
		metaProperties []properties.MetaProperty
	}
	tests := []struct {
		name string
		args args
		want []traits.Listable
	}{
		// TODO: Add test cases.
		{"+ve with nil", args{}, metaPropertiesToListables(metaPropertyList{NewList()}.GetList()...)},
		{"+ve", args{[]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}}, metaPropertiesToListables(metaPropertyList{NewList(metaPropertiesToListables([]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}.GetList()...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := metaPropertiesToListables(tt.args.metaProperties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("metaPropertiesToListables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaPropertyList_Add(t *testing.T) {
	type fields struct {
		List lists.List
	}
	type args struct {
		metaProperties []properties.MetaProperty
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.MetaPropertyList
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{NewList()}, args{}, metaPropertyList{NewList()}},
		{"+ve", fields{NewList(metaPropertiesToListables([]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, args{[]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}}, metaPropertyList{NewList(metaPropertiesToListables([]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaPropertyList := metaPropertyList{
				List: tt.fields.List,
			}
			if got := metaPropertyList.Add(tt.args.metaProperties...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaPropertyList_GetList(t *testing.T) {
	type fields struct {
		List lists.List
	}
	tests := []struct {
		name   string
		fields fields
		want   []properties.MetaProperty
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{NewList()}, []properties.MetaProperty{}},
		{"+ve", fields{NewList(metaPropertiesToListables([]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, []properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaPropertyList := metaPropertyList{
				List: tt.fields.List,
			}
			if got := metaPropertyList.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaPropertyList_GetMetaProperty(t *testing.T) {
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
		want   properties.MetaProperty
	}{
		// TODO: Add test cases.
		{"-ve", fields{NewList(metaPropertiesToListables([]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, args{base.NewMetaProperty(baseIDs.NewID("b"), NewStringData("factB")).GetID()}, nil},
		{"+ve", fields{NewList(metaPropertiesToListables([]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, args{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA")).GetID()}, base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaPropertyList := metaPropertyList{
				List: tt.fields.List,
			}
			if got := metaPropertyList.GetMetaProperty(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMetaProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaPropertyList_ToPropertyList(t *testing.T) {
	type fields struct {
		List lists.List
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.PropertyList
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{NewList()}, NewPropertyList()},
		{"+ve", fields{NewList(metaPropertiesToListables([]properties.MetaProperty{base.NewMetaProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)...)}, NewPropertyList([]properties.Property{base.NewProperty(baseIDs.NewID("a"), NewStringData("factA"))}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metaPropertyList := metaPropertyList{
				List: tt.fields.List,
			}
			if got := metaPropertyList.ToPropertyList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToPropertyList() = %v, want %v", got, tt.want)
			}
		})
	}
}
