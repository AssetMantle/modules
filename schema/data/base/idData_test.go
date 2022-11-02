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
	"github.com/AssetMantle/modules/schema/traits"
)

func TestNewIDData(t *testing.T) {
	type args struct {
		value ids.ID
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		// TODO: Add test cases.
		{"+ve", args{NewStringData("Data")}, idData{NewStringData("Data")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewIDData(tt.args.value), "NewIDData(%v)", tt.args.value)
		})
	}
}

func Test_idDataFromInterface(t *testing.T) {
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		args    args
		want    idData
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{"+ve", args{NewIDData(NewStringData("Data"))}, idData{NewStringData("Data")}, assert.NoError},
		{"+ve", args{NewIDData(NewStringData(""))}, idData{NewStringData("")}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := idDataFromInterface(tt.args.listable)
			if !tt.wantErr(t, err, fmt.Sprintf("idDataFromInterface(%v)", tt.args.listable)) {
				return
			}
			assert.Equalf(t, tt.want, got, "idDataFromInterface(%v)", tt.args.listable)
		})
	}
}

func Test_idData_Bytes(t *testing.T) {
	type fields struct {
		Value ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{NewStringData("")}, []byte{}},
		{"+ve", fields{NewStringData("Data")}, NewStringData("Data").Bytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.Bytes(), "Bytes()")
		})
	}
}

func Test_idData_Compare(t *testing.T) {
	type fields struct {
		Value ids.ID
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
		{"+ve", fields{NewStringData("Data")}, args{idData{NewStringData("Data")}}, 0},
		{"+ve", fields{NewStringData("Data")}, args{idData{NewStringData("0")}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_idData_GenerateHashID(t *testing.T) {
	type fields struct {
		Value ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{
		// TODO: Add test cases.
		{"+ve", fields{NewStringData("Data")}, baseIDs.GenerateHashID(idData{NewStringData("Data")}.Bytes())},
		{"+ve with empty String", fields{NewStringData("")}, baseIDs.GenerateHashID(idData{NewStringData("")}.Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.GenerateHashID(), "GenerateHashID()")
		})
	}
}

func Test_idData_Get(t *testing.T) {
	type fields struct {
		Value ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", fields{NewStringData("Data")}, NewStringData("Data")},
		{"+ve", fields{NewStringData("")}, NewStringData("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.Get(), "Get()")
		})
	}
}

func Test_idData_GetID(t *testing.T) {
	type fields struct {
		Value ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		// TODO: Add test cases.
		{"+ve", fields{NewStringData("Data")}, baseIDs.NewDataID(idData{NewStringData("Data")})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.GetID(), "GetID()")
		})
	}
}

func Test_idData_GetType(t *testing.T) {
	type fields struct {
		Value ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.StringID
	}{
		// TODO: Add test cases.
		{"+ve", fields{NewStringData("Data")}, dataConstants.IDDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.GetType(), "GetType()")
		})
	}
}

func Test_idData_String(t *testing.T) {
	type fields struct {
		Value ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{NewStringData("Data")}, "Data"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.String(), "String()")
		})
	}
}

func Test_idData_ZeroValue(t *testing.T) {
	type fields struct {
		Value ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		// TODO: Add test cases.
		{"+ve", fields{NewStringData("Data")}, NewIDData(baseIDs.NewStringID(""))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.ZeroValue(), "ZeroValue()")
		})
	}
}
