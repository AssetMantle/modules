// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/stretchr/testify/assert"
)

func TestNewHeightData(t *testing.T) {
	type args struct {
		value types.Height
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		{"+ve with nil", args{nil}, &HeightData{nil}},
		{"Test for +ve int", args{baseTypes.NewHeight(100)}, &HeightData{baseTypes.NewHeight(100).(*baseTypes.Height)}},
		{"Test for +ve int", args{baseTypes.NewHeight(-100)}, &HeightData{baseTypes.NewHeight(-100).(*baseTypes.Height)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeightData(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeightData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightData_Compare(t *testing.T) {
	type fields struct {
		Value types.Height
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

		{"Test for Equal case", fields{baseTypes.NewHeight(100)}, args{&HeightData{baseTypes.NewHeight(100).(*baseTypes.Height)}}, 0},
		{"Test for LT case", fields{baseTypes.NewHeight(0)}, args{&HeightData{baseTypes.NewHeight(100).(*baseTypes.Height)}}, -1},
		{"Test for GT case", fields{baseTypes.NewHeight(100)}, args{&HeightData{baseTypes.NewHeight(0).(*baseTypes.Height)}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := HeightData{
				Value: tt.fields.Value.(*baseTypes.Height),
			}
			if got := heightData.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightData_GenerateHashID(t *testing.T) {
	type fields struct {
		Value types.Height
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{

		{"Test for zero value", fields{baseTypes.NewHeight(-1)}, baseIDs.GenerateHashID()},
		{"Test for -ve value", fields{baseTypes.NewHeight(-100)}, baseIDs.GenerateHashID()},
		{"Test for +ve value", fields{baseTypes.NewHeight(100)}, baseIDs.GenerateHashID(HeightData{baseTypes.NewHeight(100).(*baseTypes.Height)}.Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := HeightData{
				Value: tt.fields.Value.(*baseTypes.Height),
			}
			if got := heightData.GenerateHashID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHashID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightData_Get(t *testing.T) {
	type fields struct {
		Value types.Height
	}
	tests := []struct {
		name   string
		fields fields
		want   types.Height
	}{

		{"Test for zero value", fields{baseTypes.NewHeight(0)}, HeightData{baseTypes.NewHeight(0).(*baseTypes.Height)}.Value},
		{"Test for +ve value", fields{baseTypes.NewHeight(100)}, HeightData{baseTypes.NewHeight(100).(*baseTypes.Height)}.Value},
		{"Test for -ve value", fields{baseTypes.NewHeight(-100)}, HeightData{baseTypes.NewHeight(-100).(*baseTypes.Height)}.Value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := HeightData{
				Value: tt.fields.Value.(*baseTypes.Height),
			}
			if got := heightData.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightData_GetID(t *testing.T) {
	type fields struct {
		Value types.Height
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{

		{"Test for zero value", fields{baseTypes.NewHeight(0)}, baseIDs.GenerateDataID(&HeightData{baseTypes.NewHeight(0).(*baseTypes.Height)})},
		{"Test for +ve value", fields{baseTypes.NewHeight(100)}, baseIDs.GenerateDataID(&HeightData{baseTypes.NewHeight(100).(*baseTypes.Height)})},
		{"Test for -ve value", fields{baseTypes.NewHeight(-100)}, baseIDs.GenerateDataID(&HeightData{baseTypes.NewHeight(-100).(*baseTypes.Height)})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := HeightData{
				Value: tt.fields.Value.(*baseTypes.Height),
			}
			if got := heightData.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightData_GetType(t *testing.T) {
	type fields struct {
		Value types.Height
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{

		{"Test for an integer value", fields{baseTypes.NewHeight(100)}, idsConstants.HeightDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := HeightData{
				Value: tt.fields.Value.(*baseTypes.Height),
			}
			if got := heightData.GetType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightData_String(t *testing.T) {
	type fields struct {
		Value types.Height
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{"Test for zero value", fields{baseTypes.NewHeight(0)}, strconv.FormatInt(HeightData{baseTypes.NewHeight(0).(*baseTypes.Height)}.Value.Get(), 10)},
		{"Test for +ve value", fields{baseTypes.NewHeight(100)}, strconv.FormatInt(HeightData{baseTypes.NewHeight(100).(*baseTypes.Height)}.Value.Get(), 10)},
		{"Test for -ve value", fields{baseTypes.NewHeight(-100)}, strconv.FormatInt(HeightData{baseTypes.NewHeight(-100).(*baseTypes.Height)}.Value.Get(), 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := HeightData{
				Value: tt.fields.Value.(*baseTypes.Height),
			}
			if got := heightData.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightData_ZeroValue(t *testing.T) {
	type fields struct {
		Value types.Height
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{

		{"Test for zero value", fields{baseTypes.NewHeight(0)}, &HeightData{baseTypes.NewHeight(-1).(*baseTypes.Height)}},
		{"Test for +ve Int value", fields{baseTypes.NewHeight(100)}, &HeightData{baseTypes.NewHeight(-1).(*baseTypes.Height)}},
		{"Test for -ve Int value", fields{baseTypes.NewHeight(-100)}, &HeightData{baseTypes.NewHeight(-1).(*baseTypes.Height)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := HeightData{
				Value: tt.fields.Value.(*baseTypes.Height),
			}
			if got := heightData.ZeroValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightData_Bytes(t *testing.T) {
	type fields struct {
		Value types.Height
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve with ZeroHeight", fields{baseTypes.NewHeight(-1)}, baseTypes.NewHeight(-1).Bytes()},
		{"+ve with nil", fields{nil}, []byte{}},
		{"+ve", fields{baseTypes.NewHeight(100)}, baseTypes.NewHeight(100).Bytes()},
		{"+ve with max int", fields{baseTypes.NewHeight(int64(^uint(0) >> 1))}, baseTypes.NewHeight(int64(^uint(0) >> 1)).Bytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := HeightData{
				Value: tt.fields.Value.(*baseTypes.Height),
			}
			assert.Equalf(t, tt.want, heightData.Bytes(), "Bytes()")
		})
	}
}
