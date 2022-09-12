// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/traits"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListData(t *testing.T) {
	//metaProperty := baseProperties.NewMetaProperty(baseIDs.NewID("id"), NewStringData("Data"))
	//metaPropertyList := base.NewMetaProperties([]properties.MetaProperty{metaProperty}...)
	type args struct {
		value []data.Data
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		// TODO: Add test cases.
		{"Test for some id", args{[]data.Data{NewStringData("Data")}}, listData{baseLists.NewDataList(NewStringData("Data"))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewListData(tt.args.value...), "NewListData(%v)", tt.args.value)
		})
	}
}

func TestReadListData(t *testing.T) {
	accAddress, _ := sdkTypes.AccAddressFromBech32("cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")
	//accAddress1, _ := sdkTypes.AccAddressFromBech32("mantle1x53dugvr4xvew442l9v2r5x7j8gfvged5xd3xr")
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.Data
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{"Test for empty string", args{""}, listData{}.ZeroValue(), assert.NoError},
		{"Test for wrong address string", args{"cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"}, NewListData(accAddressData{accAddress}), assert.NoError},
		//{"Test for correct address string", args{"mantle1x53dugvr4xvew442l9v2r5x7j8gfvged5xd3xr"}, NewListData(accAddressData{accAddress1}), assert.Error},
		{"Test for wrong address string format", args{"cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk51f"}, listData{}.ZeroValue(), assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadListData(tt.args.dataString)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadListData(%v)", tt.args.dataString)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadListData(%v)", tt.args.dataString)
		})
	}
}

func Test_listDataFromInterface(t *testing.T) {
	accAddress, _ := sdkTypes.AccAddressFromBech32("cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")
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
		{"Test for some id", args{NewListData(accAddressData{accAddress})}, listData{baseLists.NewDataList(accAddressData{accAddress})}, assert.NoError},
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

func Test_listData_Compare(t *testing.T) {
	accAddress, _ := sdkTypes.AccAddressFromBech32("cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")
	accAddress1, _ := sdkTypes.AccAddressFromBech32("cosmos1s9wtd9qw4y6udlck73m7g8ut3nnnpjnnwl7vmw")
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
		// TODO: Update test cases after fixes.
		{"Test for Equal case", fields{listData{baseLists.NewDataList(accAddressData{accAddress})}.Get()}, args{NewListData(accAddressData{accAddress})}, 0, false},
		{"Test for Not Equal case", fields{baseLists.NewDataList(accAddressData{accAddress1})}, args{NewListData(accAddressData{accAddress})}, 1, false},
		{"Test for Not Equal case", fields{baseLists.NewDataList(accAddressData{accAddress1})}, args{heightData{baseTypes.NewHeight(100)}}, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("Compare() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_listData_GenerateHash(t *testing.T) {
	accAddress, _ := sdkTypes.AccAddressFromBech32("cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")
	accAddress1, _ := sdkTypes.AccAddressFromBech32("invalidAddress")
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		//TODO: Update test cases after fixes.
		{"empty string", fields{baseLists.NewDataList()}, baseIDs.NewID("")},
		{"+ve case", fields{baseLists.NewDataList(accAddressData{accAddress})}, baseIDs.NewID("GVpq_tf8khitXl2MmMQfY-Ufu5DdATYNz3ZS9-wIl_U=")},
		{"-ve case", fields{baseLists.NewDataList(accAddressData{accAddress1})}, baseIDs.NewID("")},
		{"-ve case with empty datalist", fields{baseLists.NewDataList([]data.Data{}...)}, baseIDs.NewID("")},
		//{"-ve case with nil data", fields{nil}, baseIDs.NewID("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, listData.GenerateHash(), "GenerateHash()")
		})
	}
}

func Test_listData_Get(t *testing.T) {
	accAddress, _ := sdkTypes.AccAddressFromBech32("cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   lists.DataList
	}{
		// TODO: Add test cases.
		{"zero value", fields{baseLists.NewDataList()}, baseLists.NewDataList()},
		{"random data", fields{baseLists.NewDataList(accAddressData{accAddress})}, listData{baseLists.NewDataList(accAddressData{accAddress})}.Value},
		{"nil value", fields{baseLists.NewDataList(nil)}, baseLists.NewDataList(nil)},
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
	accAddress, _ := sdkTypes.AccAddressFromBech32("cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		// TODO: Add test cases.
		//{"nil value", fields{nil}, baseIDs.NewDataID(listData{nil})},
		{"some +ve case", fields{baseLists.NewDataList(accAddressData{accAddress})}, baseIDs.NewDataID(listData{baseLists.NewDataList(accAddressData{accAddress})})},
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
	accAddress, _ := sdkTypes.AccAddressFromBech32("cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"some +ve case", fields{baseLists.NewDataList(accAddressData{accAddress})}, idsConstants.ListDataID},
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

func Test_listData_String(t *testing.T) {
	accAddress, _ := sdkTypes.AccAddressFromBech32("cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"some +ve case", fields{baseLists.NewDataList(accAddressData{accAddress})}, "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"},
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
	accAddress, _ := sdkTypes.AccAddressFromBech32("cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")
	type fields struct {
		Value lists.DataList
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		// TODO: Add test cases.
		{"empty data", fields{baseLists.NewDataList([]data.Data{}...)}, NewListData([]data.Data{}...)},
		{"some +ve case", fields{baseLists.NewDataList(accAddressData{accAddress})}, NewListData([]data.Data{}...)},
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
