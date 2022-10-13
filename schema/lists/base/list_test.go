// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	base2 "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/traits"
)

func TestNewList(t *testing.T) {
	type args struct {
		listables []traits.Listable
	}
	tests := []struct {
		name string
		args args
		want lists.List
	}{
		// TODO: Add test cases.
		{"+ve", args{[]traits.Listable{}}, list{}},
		{"+ve", args{[]traits.Listable{base.NewStringID("ID")}}, list{base.NewStringID("ID")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.listables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Add(t *testing.T) {
	type args struct {
		listables []traits.Listable
	}
	tests := []struct {
		name string
		list list
		args args
		want lists.List
	}{
		// TODO: Add test cases.
		{"+ve with nil", list{}, args{[]traits.Listable{}}, list{}},
		{"+ve", []traits.Listable{}, args{[]traits.Listable{base.NewStringID("ID")}}, list{base.NewStringID("ID")}},
		{"+ve", list{base.NewStringID("ID")}, args{[]traits.Listable{base.NewStringID("ID1")}}, list{base.NewStringID("ID"), base.NewStringID("ID1")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Add(tt.args.listables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Get(t *testing.T) {
	tests := []struct {
		name string
		list list
		want []traits.Listable
	}{
		// TODO: Add test cases.
		{"+ve", []traits.Listable{base.NewStringID("ID")}, list{base.NewStringID("ID")}},
		{"+ve", list{}, list{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Mutate(t *testing.T) {
	type args struct {
		listables []traits.Listable
	}
	tests := []struct {
		name string
		list list
		args args
		want lists.List
	}{
		// TODO: Add test cases.
		{"+ve", []traits.Listable{base2.NewMetaProperty(base.NewStringID("ID1"), NewStringData("Data1"))}, args{[]traits.Listable{base2.NewMetaProperty(base.NewStringID("ID1"), NewStringData("Data2"))}}, list{base2.NewMetaProperty(base.NewStringID("ID1"), NewStringData("Data2"))}},
		{"+ve with nil", list{}, args{}, list{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Mutate(tt.args.listables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Remove(t *testing.T) {
	type args struct {
		listables []traits.Listable
	}
	tests := []struct {
		name string
		list list
		args args
		want lists.List
	}{
		// TODO: Add test cases.
		{"+ve for nil", []traits.Listable{}, args{}, list{}},
		{"+ve", []traits.Listable{base.NewStringID("ID")}, args{[]traits.Listable{base.NewStringID("ID2")}}, list{base.NewStringID("ID")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Remove(tt.args.listables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Search(t *testing.T) {
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name  string
		list  list
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
		{"+ve for nil", []traits.Listable{}, args{}, 0, false}, // TODO: panics if list is nil
		{"+ve", []traits.Listable{base.NewStringID("ID")}, args{base.NewStringID("ID")}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.list.Search(tt.args.listable)
			if got != tt.want {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Search() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_list_Size(t *testing.T) {
	tests := []struct {
		name string
		list list
		want int
	}{
		// TODO: Add test cases.
		{"+ve", []traits.Listable{}, 0},
		{"+ve", []traits.Listable{base.NewStringID("ID")}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
