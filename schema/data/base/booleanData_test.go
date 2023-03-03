// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
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
		{"+ve", args{true}, &BooleanData{true}},
		{"+ve", args{false}, &BooleanData{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewBooleanData(tt.args.value), "NewBooleanData(%v)", tt.args.value)
		})
	}
}

//
// func TestBooleanDataFromInterface(t *testing.T) {
//	type args struct {
//		dataString data.Data
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    data.Data
//		wantErr assert.ErrorAssertionFunc
//	}{
//		{"-ve", args{NewBooleanData(false)}, &BooleanData{false}, assert.NoError},
//		{"+ve", args{NewBooleanData(true)}, &BooleanData{true}, assert.NoError},
//		{"-ve", args{NewStringData("test")}, &BooleanData{false}, assert.Error},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := booleanDataFromInterface(tt.args.dataString)
//			if !tt.wantErr(t, err, fmt.Sprintf("booleanDataFromInterface(%v)", tt.args.dataString)) {
//				return
//			}
//			assert.Equalf(t, tt.want, got, "booleanDataFromInterface(%v)", tt.args.dataString)
//		})
//	}
// }
//
// func Test_booleanDataFromInterface(t *testing.T) {
//	type args struct {
//		listable traits.Listable
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *BooleanData
//		wantErr assert.ErrorAssertionFunc
//	}{
//		{"+ve with empty string", args{&BooleanData{}}, &BooleanData{}, assert.NoError},
//		{"+ve", args{&BooleanData{true}}, &BooleanData{true}, assert.NoError},
//		{"+ve", args{&BooleanData{false}}, &BooleanData{false}, assert.NoError},
//		{"-ve", args{NewStringData("test")}, &BooleanData{false}, assert.Error},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := booleanDataFromInterface(tt.args.listable)
//			if !tt.wantErr(t, err, fmt.Sprintf("booleanDataFromInterface(%v)", tt.args.listable)) {
//				return
//			}
//			assert.Equalf(t, tt.want, got, "booleanDataFromInterface(%v)", tt.args.listable)
//		})
//	}
// }

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
		{"+ve", fields{false}, args{&BooleanData{true}}, -1},
		{"+ve", fields{true}, args{&BooleanData{false}}, 1},
		{"+ve", fields{false}, args{&BooleanData{false}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := &BooleanData{
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
		{"+ve with nil", fields{}, baseIDs.GenerateHashID()},
		{"+ve", fields{true}, baseIDs.GenerateHashID((&BooleanData{true}).Bytes())},
		{"+ve", fields{false}, baseIDs.GenerateHashID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := &BooleanData{
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
		{"+ve", fields{}, (&BooleanData{}).Value},
		{"+ve", fields{true}, (&BooleanData{true}).Value},
		{"+ve", fields{false}, (&BooleanData{false}).Value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := &BooleanData{
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
		{"+ve", fields{}, baseIDs.GenerateDataID(&BooleanData{})},
		{"+ve", fields{true}, baseIDs.GenerateDataID(&BooleanData{true})},
		{"+ve", fields{false}, baseIDs.GenerateDataID(&BooleanData{false})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := &BooleanData{
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
		{"+ve", fields{}, idsConstants.BooleanDataTypeID},
		{"+ve", fields{true}, idsConstants.BooleanDataTypeID},
		{"+ve", fields{false}, idsConstants.BooleanDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := &BooleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.GetTypeID(), "GetTypeID()")
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
		{"+ve", fields{}, "false"},
		{"+ve", fields{true}, "true"},
		{"+ve", fields{false}, "false"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := &BooleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.AsString(), "String()")
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
		{"+ve", fields{}, &BooleanData{}},
		{"+ve", fields{true}, NewBooleanData(false)},
		{"+ve", fields{false}, NewBooleanData(false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := &BooleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.ZeroValue(), "ZeroValue()")
		})
	}
}

func Test_booleanData_Bytes(t *testing.T) {
	type fields struct {
		Value bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{true}, []byte{0x1}},
		{"+ve", fields{false}, []byte{0x0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booleanData := &BooleanData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, booleanData.Bytes(), "Bytes()")
		})
	}
}
