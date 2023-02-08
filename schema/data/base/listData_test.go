// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"fmt"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var fromAddress = "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"

var accAddress = NewAccAddressData(sdkTypes.AccAddress(fromAddress)).AsString()

func TestListDataPrototype(t *testing.T) {
	type args struct {
		value data.ListData
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		{"+ve", args{}, &ListData{[]*AnyData{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PrototypeListData(), "Prototype(%v)", tt.args.value)
		})
	}
}

func TestNewListData(t *testing.T) {
	type args struct {
		value data.ListData
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		{"+ve for some id", args{NewListData(NewStringData("Data"))}, &ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}}},
		{"+ve for empty String", args{NewListData(NewStringData(""))}, &ListData{[]*AnyData{NewStringData("").ToAnyData().(*AnyData)}}},

		{"+ve empty datalist", args{NewListData([]data.Data{}...)}, (&ListData{}).ZeroValue()},
		{"+ve address string", args{NewListData(NewStringData(fromAddress))}, &ListData{[]*AnyData{NewStringData(fromAddress).ToAnyData().(*AnyData)}}},
		// TODO: Check address format
		// {"-ve wrong address string format", args{NewListData(NewStringData(fromAddress))}, &ListData{}.ZeroValue()},
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
		want    *ListData
		wantErr assert.ErrorAssertionFunc
	}{
		{"+ve for some id", args{&ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}}}, &ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}}, assert.NoError},
		{"+ve for empty String", args{&ListData{[]*AnyData{NewStringData("").ToAnyData().(*AnyData)}}}, &ListData{[]*AnyData{NewStringData("").ToAnyData().(*AnyData)}}, assert.NoError},
		{"-ve for empty String", args{&HeightData{}}, &ListData{}, assert.Error},
		{"-ve for empty String", args{NewStringData("").ToAnyData().(*AnyData)}, &ListData{}, assert.Error},
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
		Value []*AnyData
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
		{"+ve for multiple ids", fields{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData), NewStringData("Data").ToAnyData().(*AnyData), NewStringData("Data").ToAnyData().(*AnyData)}}, args{}, &ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData), NewStringData("Data").ToAnyData().(*AnyData), NewStringData("Data").ToAnyData().(*AnyData)}}, false},
		{"+ve for multiple ids/nils", fields{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData), NewStringData("").ToAnyData().(*AnyData), NewStringData("Data").ToAnyData().(*AnyData)}}, args{}, &ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData), NewStringData("Data").ToAnyData().(*AnyData), NewStringData("").ToAnyData().(*AnyData)}}, false},
		{"+ve for some id", fields{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}}, args{}, &ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}}, false},
		{"+ve for empty String", fields{[]*AnyData{NewStringData("").ToAnyData().(*AnyData)}}, args{}, &ListData{[]*AnyData{NewStringData("").ToAnyData().(*AnyData)}}, false},
		{"-ve for value inequality", fields{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}}, args{}, &ListData{[]*AnyData{NewStringData("Data1").ToAnyData().(*AnyData)}}, true},
		{"-ve for occurrence inequality", fields{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData), NewStringData("Data").ToAnyData().(*AnyData), NewStringData("Data").ToAnyData().(*AnyData)}}, args{}, &ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData), NewStringData("Data").ToAnyData().(*AnyData)}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := &ListData{
				DataList: tt.fields.Value,
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
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data").ToAnyData().(*AnyData))}, NewStringData("Data").Bytes()}, // for a single data no loop iteration is required, so directly it's byte should match
		{"+ve for multiple ids", fields{NewListData(NewStringData("Data").ToAnyData().(*AnyData), NewStringData("Data1"))}, bytes.Join([][]byte{NewStringData("Data").Bytes(), NewStringData("Data1").Bytes()}, nil)},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, []byte(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.Bytes(), "Bytes()")
		})
	}
}

func Test_listData_Compare(t *testing.T) {
	type fields struct {
		Value data.ListData
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
		{"+ve for some id", fields{NewListData(NewStringData("Data").ToAnyData().(*AnyData))}, args{&ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}}}, 0, false},
		{"+ve for empty String", fields{NewListData(NewStringData("").ToAnyData().(*AnyData))}, args{&ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}}}, -1, false},
		{"Test for Equal case", fields{NewListData(NewStringData(fromAddress).ToAnyData().(*AnyData))}, args{&ListData{[]*AnyData{NewStringData(fromAddress).ToAnyData().(*AnyData)}}}, 0, false},
		{"Test for Not Equal case", fields{NewListData(NewStringData(fromAddress))}, args{&ListData{[]*AnyData{NewStringData(accAddress).ToAnyData().(*AnyData)}}}, 1, false},
		{"panic test for Not Equal case", fields{NewListData(NewStringData(accAddress))}, args{&HeightData{baseTypes.NewHeight(100).(*baseTypes.Height)}}, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
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
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, baseIDs.GenerateHashID((&ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}}).Bytes()).AsString()},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, baseIDs.GenerateHashID((&ListData{[]*AnyData{NewStringData("").ToAnyData().(*AnyData)}}).Bytes()).AsString()},
		{"empty string", fields{NewListData()}, baseIDs.NewStringID("").AsString()},
		{"+ve case", fields{NewListData(NewStringData(accAddress))}, baseIDs.NewStringID("xrHmURH4R458qdPeDW8kU9eO3a3bvQRE0W6CAoZ8yCw=").AsString()},
		{"-ve case", fields{NewListData(NewStringData(""))}, baseIDs.NewStringID("").AsString()},
		{"-ve case with empty datalist", fields{NewListData([]data.Data{}...)}, baseIDs.NewStringID("").AsString()},
		{"-ve case with nil data", fields{NewListData()}, baseIDs.NewStringID("").AsString()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.GenerateHashID().AsString(), "GenerateHashID()")
		})
	}
}

func Test_listData_Get(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   []*AnyData
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, (&ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}}).DataList},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, (&ListData{[]*AnyData{NewStringData("").ToAnyData().(*AnyData)}}).DataList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.Get(), "Get()")
		})
	}
}

func Test_listData_GetID(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, baseIDs.GenerateDataID(&ListData{[]*AnyData{NewStringData("Data").ToAnyData().(*AnyData)}})},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, baseIDs.GenerateDataID(&ListData{[]*AnyData{NewStringData("").ToAnyData().(*AnyData)}})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.GetID(), "GetID()")
		})
	}
}

func Test_listData_GetType(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.StringID
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, dataConstants.ListDataID},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, dataConstants.ListDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.GetType(), "GetType()")
		})
	}
}

func Test_listData_Remove(t *testing.T) {
	type fields struct {
		Value data.ListData
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
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, args{[]data.Data{}}, &ListData{[]*AnyData{NewStringData("").ToAnyData().(*AnyData)}}},
		{"+ve for empty String & removing it", fields{NewListData(NewStringData(""))}, args{[]data.Data{NewStringData("").ToAnyData().(*AnyData)}}, &ListData{[]*AnyData{}}},
		{"+ve ", fields{NewListData(NewStringData("data"))}, args{[]data.Data{NewStringData("data")}}, &ListData{[]*AnyData{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.Remove(tt.args.data...), "Remove(%v)", tt.args.data)
		})
	}
}

func Test_listData_Search(t *testing.T) {
	type fields struct {
		Value data.ListData
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
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, args{NewStringData("Data").ToAnyData().(*AnyData)}, 0, true},
		{"+ve for empty String", fields{NewListData([]data.Data{NewStringData("Data").ToAnyData().(*AnyData), NewStringData("").ToAnyData().(*AnyData)}...)}, args{NewStringData("").ToAnyData().(*AnyData)}, 0, true},
		{"-ve", fields{NewListData([]data.Data{NewStringData("Data").ToAnyData().(*AnyData), NewStringData("").ToAnyData().(*AnyData)}...)}, args{NewStringData("test")}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			got, got1 := listData.Search(tt.args.data)
			assert.Equalf(t, tt.want, got, "Search(%v)", tt.args.data)
			assert.Equalf(t, tt.want1, got1, "Search(%v)", tt.args.data)
		})
	}
}

func Test_listData_String(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, "Data"},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.AsString(), "String()")
		})
	}
}

func Test_listData_ZeroValue(t *testing.T) {
	type fields struct {
		Value data.ListData
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		{"+ve for some id", fields{NewListData(NewStringData("Data"))}, NewListData(NewListData([]data.Data{}...))},
		{"+ve for empty String", fields{NewListData(NewStringData(""))}, NewListData(NewListData([]data.Data{}...))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := tt.fields.Value
			assert.Equalf(t, tt.want, listData.ZeroValue(), "ZeroValue()")
		})
	}
}
