// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

func TestNewBooleanData(t *testing.T) {
	type args struct {
		value bool
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		// TODO: Add test cases.
		{"+ve", args{true}, booleanData{true}},
		{"+ve", args{false}, booleanData{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewBooleanData(tt.args.value), "NewBooleanData(%v)", tt.args.value)
		})
	}
}

func TestBooleanDataFromInterface(t *testing.T) {
	type args struct {
		dataString data.Data
	}
	tests := []struct {
		name    string
		args    args
		want    data.Data
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{"-ve", args{NewBooleanData(false)}, booleanData{false}, assert.NoError},
		{"+ve", args{NewBooleanData(true)}, booleanData{true}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := booleanDataFromInterface(tt.args.dataString)
			if !tt.wantErr(t, err, fmt.Sprintf("booleanDataFromInterface(%v)", tt.args.dataString)) {
				return
			}
			assert.Equalf(t, tt.want, got, "booleanDataFromInterface(%v)", tt.args.dataString)
		})
	}
}

func Test_booleanDataFromInterface(t *testing.T) {
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		args    args
		want    booleanData
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{"+ve with empty string", args{booleanData{}}, booleanData{}, assert.NoError},
		{"+ve", args{booleanData{true}}, booleanData{true}, assert.NoError},
		{"-ve", args{booleanData{false}}, booleanData{false}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := booleanDataFromInterface(tt.args.listable)
			if !tt.wantErr(t, err, fmt.Sprintf("booleanDataFromInterface(%v)", tt.args.listable)) {
				return
			}
			assert.Equalf(t, tt.want, got, "booleanDataFromInterface(%v)", tt.args.listable)
		})
	}
}

func Test_booleanData_Compare(t *testing.T) {
	type fields struct {
		Value bool
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
		{"+ve", fields{false}, args{booleanData{true}}, -1},
		{"+ve", fields{true}, args{booleanData{false}}, 1},
		{"+ve", fields{false}, args{booleanData{false}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := booleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_booleanData_GenerateHashID(t *testing.T) {
	type fields struct {
		Value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, baseIDs.GenerateHashID()},
		{"+ve", fields{true}, baseIDs.GenerateHashID(booleanData{true}.Bytes())},
		{"+ve", fields{false}, baseIDs.GenerateHashID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := booleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.GenerateHashID(), "GenerateHashID()")
		})
	}
}

func Test_booleanData_Get(t *testing.T) {
	type fields struct {
		Value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{}, booleanData{}.Value},
		{"+ve", fields{true}, booleanData{true}.Value},
		{"+ve", fields{false}, booleanData{false}.Value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := booleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.Get(), "Get()")
		})
	}
}

func Test_booleanData_GetID(t *testing.T) {
	type fields struct {
		Value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		// TODO: Add test cases.
		{"+ve", fields{}, baseIDs.NewDataID(booleanData{})},
		{"+ve", fields{true}, baseIDs.NewDataID(booleanData{true})},
		{"+ve", fields{false}, baseIDs.NewDataID(booleanData{false})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := booleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.GetID(), "GetID()")
		})
	}
}

func Test_booleanData_GetType(t *testing.T) {
	type fields struct {
		Value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", fields{}, idsConstants.BooleanDataID},
		{"+ve", fields{true}, idsConstants.BooleanDataID},
		{"+ve", fields{false}, idsConstants.BooleanDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := booleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.GetType(), "GetType()")
		})
	}
}

func Test_booleanData_String(t *testing.T) {
	type fields struct {
		Value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{}, "false"},
		{"+ve", fields{true}, "true"},
		{"+ve", fields{false}, "false"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := booleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.String(), "String()")
		})
	}
}

func Test_booleanData_ZeroValue(t *testing.T) {
	type fields struct {
		Value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		// TODO: Add test cases.
		{"+ve", fields{}, booleanData{}},
		{"+ve", fields{true}, NewBooleanData(false)},
		{"+ve", fields{false}, NewBooleanData(false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := booleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.ZeroValue(), "ZeroValue()")
		})
	}
}
