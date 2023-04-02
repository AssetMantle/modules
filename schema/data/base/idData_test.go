// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
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
		{"+ve", args{baseIDs.NewStringID("Data")}, &IDData{baseIDs.NewStringID("Data").ToAnyID().(*baseIDs.AnyID)}},
		{"+ve empty string", args{baseIDs.NewStringID("")}, &IDData{baseIDs.NewStringID("").ToAnyID().(*baseIDs.AnyID)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewIDData(tt.args.value), "NewIDData(%v)", tt.args.value)
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
		{"+ve", fields{baseIDs.NewStringID("")}, []byte{}},
		{"+ve", fields{baseIDs.NewStringID("Data")}, baseIDs.NewStringID("Data").Bytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := IDData{
				Value: tt.fields.Value.ToAnyID().(*baseIDs.AnyID),
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
		{"+ve", fields{baseIDs.NewStringID("Data")}, args{&IDData{baseIDs.NewStringID("Data").ToAnyID().(*baseIDs.AnyID)}}, 0},
		{"+ve", fields{baseIDs.NewStringID("Data")}, args{&IDData{baseIDs.NewStringID("0").ToAnyID().(*baseIDs.AnyID)}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := IDData{
				Value: tt.fields.Value.ToAnyID().(*baseIDs.AnyID),
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
		{"+ve", fields{baseIDs.NewStringID("Data")}, baseIDs.GenerateHashID((&IDData{baseIDs.NewStringID("Data").ToAnyID().(*baseIDs.AnyID)}).Bytes())},
		{"+ve with empty String", fields{baseIDs.NewStringID("")}, baseIDs.GenerateHashID((&IDData{baseIDs.NewStringID("").ToAnyID().(*baseIDs.AnyID)}).Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := IDData{
				Value: tt.fields.Value.ToAnyID().(*baseIDs.AnyID),
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
		{"+ve", fields{baseIDs.NewStringID("Data")}, baseIDs.NewStringID("Data")},
		{"+ve", fields{baseIDs.NewStringID("")}, baseIDs.NewStringID("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := IDData{
				Value: tt.fields.Value.ToAnyID().(*baseIDs.AnyID),
			}
			assert.Equalf(t, tt.want, idData.Get().Get(), "Get()")
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
		{"+ve", fields{baseIDs.NewStringID("Data")}, baseIDs.GenerateDataID(&IDData{baseIDs.NewStringID("Data").ToAnyID().(*baseIDs.AnyID)})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := IDData{
				Value: tt.fields.Value.ToAnyID().(*baseIDs.AnyID),
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
		{"+ve", fields{baseIDs.NewStringID("Data")}, dataConstants.IDDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := IDData{
				Value: tt.fields.Value.ToAnyID().(*baseIDs.AnyID),
			}
			assert.Equalf(t, tt.want, idData.GetTypeID(), "GetTypeID()")
		})
	}
}

func Test_idData_AsString(t *testing.T) {
	type fields struct {
		Value ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{baseIDs.NewStringID("Data")}, "Data"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := IDData{
				Value: tt.fields.Value.ToAnyID().(*baseIDs.AnyID),
			}
			assert.Equalf(t, tt.want, idData.AsString(), "String()")
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
		{"+ve", fields{baseIDs.NewStringID("Data")}, NewIDData(baseIDs.NewStringID(""))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := IDData{
				Value: tt.fields.Value.ToAnyID().(*baseIDs.AnyID),
			}
			assert.Equalf(t, tt.want, idData.ZeroValue(), "ZeroValue()")
		})
	}
}
