// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"fmt"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	baseTypes "github.com/AssetMantle/modules/schema/types/base"

	"github.com/stretchr/testify/assert"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var fromAddress = "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"

var accAddress = NewAccAddressData(sdkTypes.AccAddress(fromAddress)).String()

func TestListDataPrototype(t *testing.T) {
	type args struct {
		value lists.DataList
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		{"+ve", args{}, listData{}.ZeroValue()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ListDataPrototype(), "Prototype(%v)", tt.args.value)
		})
	}
}

func TestNewListData(t *testing.T) {
	type args struct {
		value lists.DataList
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		{"+ve for some id", args{baseLists.NewDataList(NewStringData("Data"))}, listData{baseLists.NewDataList(NewStringData("Data"))}},
		{"+ve for empty String", args{baseLists.NewDataList(NewStringData(""))}, listData{baseLists.NewDataList(NewStringData(""))}},

		{"+ve empty datalist", args{baseLists.NewDataList([]data.Data{}...)}, listData{}.ZeroValue()},
		{"+ve address string", args{baseLists.NewDataList(NewStringData(fromAddress))}, listData{baseLists.NewDataList(NewStringData(fromAddress))}},
		// TODO: Check address format
		//{"-ve wrong address string format", args{baseLists.NewDataList(NewStringData(fromAddress))}, listData{}.ZeroValue()},
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
		{"+ve for some id", args{listData{baseLists.NewDataList(NewStringData("Data"))}}, listData{baseLists.NewDataList(NewStringData("Data"))}, assert.NoError},
		{"+ve for empty String", args{listData{baseLists.NewDataList(NewStringData(""))}}, listData{baseLists.NewDataList(NewStringData(""))}, assert.NoError},
		{"-ve for empty String", args{heightData{}}, listData{}, assert.Error},
		{"-ve for empty String", args{NewStringData("")}, listData{}, assert.Error},
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
		name        string
		fields      fields
		args        args
		want        data.ListData
		wantFailure bool
	}{
		{"+ve for multiple ids", fields{baseLists.NewDataList(NewStringData("Data"), NewStringData("Data"), NewStringData("Data"))}, args{}, listData{baseLists.NewDataList(NewStringData("Data"), NewStringData("Data"), NewStringData("Data"))}, false},
		{"+ve for multiple ids/nils", fields{baseLists.NewDataList(NewStringData("Data"), NewStringData(""), NewStringData("Data"))}, args{}, listData{baseLists.NewDataList(NewStringData("Data"), NewStringData("Data"), NewStringData(""))}, false},
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, args{}, listData{baseLists.NewDataList(NewStringData("Data"))}, false},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, args{}, listData{baseLists.NewDataList(NewStringData(""))}, false},
		{"-ve for value inequality", fields{baseLists.NewDataList(NewStringData("Data"))}, args{}, listData{baseLists.NewDataList(NewStringData("Data1"))}, true},
		{"-ve for occurrence inequality", fields{baseLists.NewDataList(NewStringData("Data"), NewStringData("Data"), NewStringData("Data"))}, args{}, listData{baseLists.NewDataList(NewStringData("Data"), NewStringData("Data"))}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			if tt.wantFailure {
				assert.NotEqualf(t, tt.want, listData.Add(tt.args.data...), "Add(%v)", tt.args.data)
			} else {
				assert.Equalf(t, tt.want, listData.Add(tt.args.data...), "Add(%v)", tt.args.data)
			}
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
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, NewStringData("Data").Bytes()}, // for a single data no loop iteration is required, so directly it's byte should match
		{"+ve for multiple ids", fields{baseLists.NewDataList(NewStringData("Data"), NewStringData("Data1"))}, bytes.Join([][]byte{NewStringData("Data").Bytes(), NewStringData("Data1").Bytes()}, nil)},
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
		name      string
		fields    fields
		args      args
		want      int
		wantPanic bool
	}{
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, args{listData{baseLists.NewDataList(NewStringData("Data"))}}, 0, false},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, args{listData{baseLists.NewDataList(NewStringData("Data"))}}, -1, false},
		{"Test for Equal case", fields{baseLists.NewDataList(NewStringData(fromAddress))}, args{listData{baseLists.NewDataList(NewStringData(fromAddress))}}, 0, false},
		{"Test for Not Equal case", fields{baseLists.NewDataList(NewStringData(fromAddress))}, args{listData{baseLists.NewDataList(NewStringData(accAddress))}}, 1, false},
		{"panic test for Not Equal case", fields{baseLists.NewDataList(NewStringData(accAddress))}, args{heightData{baseTypes.NewHeight(100)}}, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					listData.Compare(tt.args.listable)
				})
			} else {
				assert.Equalf(t, tt.want, listData.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
			}
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
		want   string
	}{
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, baseIDs.GenerateHashID(listData{baseLists.NewDataList(NewStringData("Data"))}.Bytes()).String()},
		{"+ve for empty String", fields{baseLists.NewDataList(NewStringData(""))}, baseIDs.GenerateHashID(listData{baseLists.NewDataList(NewStringData(""))}.Bytes()).String()},
		{"empty string", fields{baseLists.NewDataList()}, baseIDs.NewStringID("").String()},
		{"+ve case", fields{baseLists.NewDataList(NewStringData(accAddress))}, baseIDs.NewStringID("xrHmURH4R458qdPeDW8kU9eO3a3bvQRE0W6CAoZ8yCw=").String()},
		{"-ve case", fields{baseLists.NewDataList(NewStringData(""))}, baseIDs.NewStringID("").String()},
		{"-ve case with empty datalist", fields{baseLists.NewDataList([]data.Data{}...)}, baseIDs.NewStringID("").String()},
		{"-ve case with nil data", fields{baseLists.NewDataList()}, baseIDs.NewStringID("").String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.GenerateHashID().String(), "GenerateHashID()")
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
		{"+ve for some id", fields{baseLists.NewDataList(NewStringData("Data"))}, args{NewStringData("Data")}, 0, true},
		{"+ve for empty String", fields{baseLists.NewDataList([]data.Data{NewStringData("Data"), NewStringData("")}...)}, args{NewStringData("")}, 0, true},
		{"-ve", fields{baseLists.NewDataList([]data.Data{NewStringData("Data"), NewStringData("")}...)}, args{NewStringData("test")}, 2, false},
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
