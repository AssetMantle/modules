// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
	"reflect"
	"testing"
)

func TestNewDataList(t *testing.T) {
	type args struct {
		data []data.Data
	}
	tests := []struct {
		name string
		args args
		want lists.DataList
	}{
		// TODO: Add test cases.
		{"+ve with empty struct", args{[]data.Data{}}, dataList{List: NewList(dataToListables([]data.Data{}...)...)}},
		{"+ve", args{[]data.Data{NewStringData("Data")}}, dataList{List: NewList(dataToListables([]data.Data{NewStringData("Data")}...)...)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDataList(tt.args.data...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDataList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataList_Add(t *testing.T) {
	type fields struct {
		List lists.List
	}
	type args struct {
		data []data.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.DataList
	}{
		// TODO: Add test cases.
		{"+ve with no dataList", fields{List: NewList(dataToListables([]data.Data{}...)...)}, args{}, dataList{List: NewList(dataToListables([]data.Data{}...)...)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataList := dataList{
				List: tt.fields.List,
			}
			if got := dataList.Add(tt.args.data...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataList_GetList(t *testing.T) {
	type fields struct {
		List lists.List
	}
	tests := []struct {
		name   string
		fields fields
		want   []data.Data
	}{
		// TODO: Add test cases.
		{"+ve with empty struct", fields{list{}}, []data.Data{}}, //TODO: issue Panic for nil
		{"+ve", fields{NewList(dataToListables([]data.Data{NewStringData("Data")}...)...)}, []data.Data{NewStringData("Data")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataList := dataList{
				List: tt.fields.List,
			}
			if got := dataList.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataList_Remove(t *testing.T) {
	type fields struct {
		List lists.List
	}
	type args struct {
		data []data.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.DataList
	}{
		// TODO: Add test cases.
		{"+ve with empty struct", fields{NewList(dataToListables([]data.Data{}...)...)}, args{}, dataList{List: NewList(dataToListables([]data.Data{}...)...)}},
		{"+ve", fields{NewList(dataToListables([]data.Data{NewStringData("Data")}...)...)}, args{[]data.Data{NewStringData("Data")}}, dataList{List: NewList(dataToListables([]data.Data{}...)...)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataList := dataList{
				List: tt.fields.List,
			}
			if got := dataList.Remove(tt.args.data...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataList_Search(t *testing.T) {
	type fields struct {
		List lists.List
	}
	type args struct {
		data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  bool
	}{
		// TODO: Add test cases.
		{"+ve with empty struct", fields{NewList(dataToListables([]data.Data{}...)...)}, args{NewStringData("Data")}, 0, false}, //TODO: fix this
		{"+ve", fields{NewList(dataToListables([]data.Data{NewStringData("Data")}...)...)}, args{NewStringData("Data")}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataList := dataList{
				List: tt.fields.List,
			}
			got, got1 := dataList.Search(tt.args.data)
			if got != tt.want {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Search() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_dataToListables(t *testing.T) {
	type args struct {
		data []data.Data
	}
	tests := []struct {
		name string
		args args
		want []traits.Listable
	}{
		// TODO: Add test cases.
		{"+ve with empty struct", args{[]data.Data{}}, []traits.Listable{}},
		{"+ve", args{[]data.Data{NewStringData("Data")}}, []traits.Listable{NewStringData("Data")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dataToListables(tt.args.data...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataToListables() = %v, want %v", got, tt.want)
			}
		})
	}
}
