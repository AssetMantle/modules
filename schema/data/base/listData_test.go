// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/traits"
)

func TestNewListData(t *testing.T) {
	type args struct {
		value lists.DataList
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		// TODO: Add test cases.
		{"+ve for some id", args{baseLists.NewDataList(NewStringData("Data"))}, listData{baseLists.NewDataList(NewStringData("Data"))}},
		{"+ve for empty String", args{baseLists.NewDataList(NewStringData(""))}, listData{baseLists.NewDataList(NewStringData(""))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewListData(tt.args.value), "NewListData(%v)", tt.args.value)
		})
	}
}

func Test_listDataFromInterface(t *testing.T) {
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		args    args
		want    listData
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{"+ve for some id", args{listData{baseLists.NewDataList(NewStringData("Data"))}}, listData{baseLists.NewDataList(NewStringData("Data"))}, assert.NoError},
		{"+ve for empty String", args{listData{baseLists.NewDataList(NewStringData(""))}}, listData{baseLists.NewDataList(NewStringData(""))}, assert.NoError},
		{"+ve for empty String", args{heightData{}}, listData{}, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := listDataFromInterface(tt.args.listable)
			if !tt.wantErr(t, err, fmt.Sprintf("listDataFromInterface(%v)", tt.args.listable)) {
				return
			}
			assert.Equalf(t, tt.want, got, "listDataFromInterface(%v)", tt.args.listable)
		})
	}
}

func Test_listData_Add(t *testing.T) {
	type fields struct {
		Value lists.DataList
	}
	type args struct {
		data []data.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   data.ListData
	}{
		// TODO: Add test cases.
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, args{}, listData{baseLists.NewDataList(NewStringData("Data"))}},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, args{}, listData{baseLists.NewDataList(NewStringData(""))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.Add(tt.args.data...), "Add(%v)", tt.args.data)
		})
	}
}

func Test_listData_Bytes(t *testing.T) {
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, NewStringData("Data").Bytes()}, //for a single data no loop iteration is required so directly it's byte should match
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, []byte(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.Bytes(), "Bytes()")
		})
	}
}

func Test_listData_Compare(t *testing.T) {
	type fields struct {
		Value lists.DataList
	}
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, args{listData{baseLists.NewDataList(NewStringData("Data"))}}, 0},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, args{listData{baseLists.NewDataList(NewStringData("Data"))}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_listData_GenerateHashID(t *testing.T) {
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{
		// TODO: Add test cases.
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, baseIDs.GenerateHashID(listData{baseLists.NewDataList(NewStringData("Data"))}.Bytes())},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, baseIDs.GenerateHashID(listData{baseLists.NewDataList(NewStringData(""))}.Bytes())},
		{"empty string", fields{baseLists.NewDataList()}, baseIDs.GenerateHashID()},
		// {"-ve case with nil data", fields{nil}, baseIDs.NewStringID("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.GenerateHashID(), "GenerateHashID()")
		})
	}
}

func Test_listData_Get(t *testing.T) {
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   []data.Data
	}{
		// TODO: Add test cases.
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, listData{baseLists.NewDataList(NewStringData("Data"))}.Value.GetList()},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, listData{baseLists.NewDataList(NewStringData(""))}.Value.GetList()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.Get(), "Get()")
		})
	}
}

func Test_listData_GetID(t *testing.T) {
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		// TODO: Add test cases.
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, baseIDs.NewDataID(listData{baseLists.NewDataList(NewStringData("Data"))})},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, baseIDs.NewDataID(listData{baseLists.NewDataList(NewStringData(""))})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.GetID(), "GetID()")
		})
	}
}

func Test_listData_GetType(t *testing.T) {
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.StringID
	}{
		// TODO: Add test cases.
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, dataConstants.ListDataID},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, dataConstants.ListDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.GetType(), "GetType()")
		})
	}
}

func Test_listData_Remove(t *testing.T) {
	type fields struct {
		Value lists.DataList
	}
	type args struct {
		data []data.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   data.ListData
	}{
		// TODO: Add test cases.
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, args{[]data.Data{}}, listData{baseLists.NewDataList(NewStringData(""))}},
		{"+ve for empty String & removing it", fields{baseLists.NewDataList(NewStringData(""))}, args{[]data.Data{NewStringData("")}}, listData{baseLists.NewDataList()}},
		{"+ve ", fields{baseLists.NewDataList(NewStringData("data"))}, args{[]data.Data{NewStringData("data")}}, listData{baseLists.NewDataList()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.Remove(tt.args.data...), "Remove(%v)", tt.args.data)
		})
	}
}

func Test_listData_Search(t *testing.T) {
	type fields struct {
		Value lists.DataList
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
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, args{NewStringData("Data")}, 0, true},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, args{NewStringData("")}, 0, true},
		{"-ve", fields{baseLists.NewDataList(NewStringData("Data"))}, args{NewStringData("")}, 1, false}, //TODO: Report this issue
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			got, got1 := listData.Search(tt.args.data)
			assert.Equalf(t, tt.want, got, "Search(%v)", tt.args.data)
			assert.Equalf(t, tt.want1, got1, "Search(%v)", tt.args.data)
		})
	}
}

func Test_listData_String(t *testing.T) {
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, "Data"},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.String(), "String()")
		})
	}
}

func Test_listData_ZeroValue(t *testing.T) {
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		// TODO: Add test cases.
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, NewListData(baseLists.NewDataList([]data.Data{}...))},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, NewListData(baseLists.NewDataList([]data.Data{}...))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.ZeroValue(), "ZeroValue()")
		})
	}
}
