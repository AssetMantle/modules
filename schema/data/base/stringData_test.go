// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStringData(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		// TODO: Add test cases.
		{"+ve data", args{"data"}, stringData{"data"}},
		{"special char data", args{"data%/@1!"}, stringData{"data%/@1!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewStringData(tt.args.value), "NewStringData(%v)", tt.args.value)
		})
	}
}

func Test_stringDataFromInterface(t *testing.T) {
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		args    args
		want    stringData
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{"+ve data", args{stringData{"data"}}, stringData{"data"}, assert.NoError},
		{"data with special char", args{stringData{"data_!@#$%^&*("}}, stringData{"data_!@#$%^&*("}, assert.NoError},
		{"empty string", args{stringData{""}}, stringData{""}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringDataFromInterface(tt.args.listable)
			if !tt.wantErr(t, err, fmt.Sprintf("stringDataFromInterface(%v)", tt.args.listable)) {
				return
			}
			assert.Equalf(t, tt.want, got, "stringDataFromInterface(%v)", tt.args.listable)
		})
	}
}

func Test_stringData_Bytes(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringData := stringData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, stringData.Bytes(), "Bytes()")
		})
	}
}

func Test_stringData_Compare(t *testing.T) {
	type fields struct {
		Value string
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
		{"+ve data", fields{"data"}, args{stringData{"data"}}, 0},
		{"data with special char", fields{"data"}, args{stringData{"data_!@#$%^&*("}}, -1},
		{"empty string", fields{"data"}, args{stringData{""}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringData := stringData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, stringData.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_stringData_GenerateHashID(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{
		// TODO: Add test cases.
		{"+ve data", fields{"data"}, baseIDs.GenerateHashID(stringData{"data"}.Bytes())},
		{"data with special char", fields{"data_!@#$%^&*("}, baseIDs.GenerateHashID(stringData{"data_!@#$%^&*("}.Bytes())},
		{"empty string", fields{""}, baseIDs.GenerateHashID(stringData{""}.Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringData := stringData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, stringData.GenerateHashID(), "GenerateHashID()")
		})
	}
}

func Test_stringData_Get(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve data", fields{"data"}, "data"},
		{"data with special char", fields{"data_!@#$%^&*("}, "data_!@#$%^&*("},
		{"empty string", fields{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringData := stringData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, stringData.Get(), "Get()")
		})
	}
}

func Test_stringData_GetID(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		// TODO: Add test cases.
		{"+ve data", fields{"data"}, baseIDs.NewDataID(stringData{"data"})},
		{"data with special char", fields{"data_!@#$%^&*("}, baseIDs.NewDataID(stringData{"data_!@#$%^&*("})},
		{"empty string", fields{""}, baseIDs.NewDataID(stringData{""})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringData := stringData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, stringData.GetID(), "GetID()")
		})
	}
}

func Test_stringData_GetType(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.StringID
	}{
		// TODO: Add test cases.
		{"+ve data", fields{"data"}, idsConstants.StringDataID},
		{"data with special char", fields{"data_!@#$%^&*("}, idsConstants.StringDataID},
		{"empty string", fields{""}, idsConstants.StringDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringData := stringData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, stringData.GetType(), "GetType()")
		})
	}
}

func Test_stringData_String(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve data", fields{"data"}, "data"},
		{"data with special char", fields{"data_!@#$%^&*("}, "data_!@#$%^&*("},
		{"empty string", fields{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringData := stringData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, stringData.String(), "String()")
		})
	}
}

func Test_stringData_ZeroValue(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		// TODO: Add test cases.
		{"+ve data", fields{"data"}, stringData{""}},
		{"data with special char", fields{"data_!@#$%^&*("}, stringData{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringData := stringData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, stringData.ZeroValue(), "ZeroValue()")
		})
	}
}
