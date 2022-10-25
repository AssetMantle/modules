// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

func TestNewDecData(t *testing.T) {
	type args struct {
		value types.Dec
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		// TODO: Add test cases.
		{"+ve with nil", args{}, decData{}},
		{"+ve with zero dec", args{types.ZeroDec()}, decData{types.ZeroDec()}},
		{"+ve", args{types.NewDec(100)}, decData{types.NewDec(100)}},
		{"+ve with -ve Dec", args{types.NewDec(-100)}, decData{types.NewDec(-100)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewDecData(tt.args.value), "NewDecData(%v)", tt.args.value)
		})
	}
}

func Test_decDataFromInterface(t *testing.T) {
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		args    args
		want    decData
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{"+ve with nil", args{decData{}}, decData{}, assert.NoError},
		{"+ve with zero dec", args{decData{types.ZeroDec()}}, decData{types.ZeroDec()}, assert.NoError},
		{"+ve", args{decData{types.NewDec(100)}}, decData{types.NewDec(100)}, assert.NoError},
		{"+ve with -ve Dec", args{decData{types.NewDec(-100)}}, decData{types.NewDec(-100)}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decDataFromInterface(tt.args.listable)
			if !tt.wantErr(t, err, fmt.Sprintf("decDataFromInterface(%v)", tt.args.listable)) {
				return
			}
			assert.Equalf(t, tt.want, got, "decDataFromInterface(%v)", tt.args.listable)
		})
	}
}

func Test_decData_Bytes(t *testing.T) {
	type fields struct {
		Value types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		// {"+ve with nil", fields{}}, []byte{}},
		{"+ve with zero dec", fields{types.ZeroDec()}, decData{types.ZeroDec()}.Value.Bytes()},
		{"+ve", fields{types.NewDec(100)}, decData{types.NewDec(100)}.Value.Bytes()},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, decData{types.NewDec(-100)}.Value.Bytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, decData.Bytes(), "Bytes()")
		})
	}
}

func Test_decData_Compare(t *testing.T) {
	type fields struct {
		Value types.Dec
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
		{"+ve with nil", fields{}, args{decData{}}, 0},
		{"+ve with zero dec", fields{types.ZeroDec()}, args{decData{types.ZeroDec()}}, 0},
		{"+ve", fields{types.NewDec(100)}, args{decData{types.NewDec(100)}}, 0},
		{"-ve", fields{types.NewDec(-100)}, args{decData{types.NewDec(100)}}, -1},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, args{decData{types.NewDec(-100)}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, decData.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_decData_GenerateHashID(t *testing.T) {
	type fields struct {
		Value types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.HashID
	}{
		// TODO: Update Unit tests after fix
		{"+ve with nil", fields{types.Dec{}}, baseIDs.GenerateHashID(decData{}.Bytes())}, // TODO: Panics for nil Dec
		{"+ve with zero dec", fields{types.ZeroDec()}, baseIDs.GenerateHashID()},
		{"+ve", fields{types.NewDec(100)}, baseIDs.GenerateHashID(decData{types.NewDec(100)}.Bytes())},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, baseIDs.GenerateHashID(decData{types.NewDec(-100)}.Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, decData.GenerateHashID(), "GenerateHashID()")
		})
	}
}

func Test_decData_Get(t *testing.T) {
	type fields struct {
		Value types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   types.Dec
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, decData{}.Value},
		{"+ve with zero dec", fields{types.ZeroDec()}, decData{types.ZeroDec()}.Value},
		{"+ve", fields{types.NewDec(100)}, decData{types.NewDec(100)}.Value},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, decData{types.NewDec(-100)}.Value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, decData.Get(), "Get()")
		})
	}
}

func Test_decData_GetID(t *testing.T) {
	type fields struct {
		Value types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{types.Dec{}}, baseIDs.NewDataID(decData{})}, // TODO: paincs for nil
		{"+ve with zero dec", fields{types.ZeroDec()}, baseIDs.NewDataID(decData{types.ZeroDec()})},
		{"+ve", fields{types.NewDec(100)}, baseIDs.NewDataID(decData{types.NewDec(100)})},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, baseIDs.NewDataID(decData{types.NewDec(-100)})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			if got := decData.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decData_GetType(t *testing.T) {
	type fields struct {
		Value types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.StringID
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, dataConstants.DecDataID},
		{"+ve with zero dec", fields{types.ZeroDec()}, dataConstants.DecDataID},
		{"+ve", fields{types.NewDec(100)}, dataConstants.DecDataID},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, dataConstants.DecDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, decData.GetType(), "GetType()")
		})
	}
}

func Test_decData_String(t *testing.T) {
	type fields struct {
		Value types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, decData{}.Value.String()},
		{"+ve with zero dec", fields{types.ZeroDec()}, decData{types.ZeroDec()}.Value.String()},
		{"+ve", fields{types.NewDec(100)}, decData{types.NewDec(100)}.Value.String()},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, decData{types.NewDec(-100)}.Value.String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, decData.String(), "String()")
		})
	}
}

func Test_decData_ZeroValue(t *testing.T) {
	type fields struct {
		Value types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, decData{types.ZeroDec()}},
		{"+ve with zero dec", fields{types.ZeroDec()}, decData{types.ZeroDec()}},
		{"+ve", fields{types.NewDec(100)}, decData{types.ZeroDec()}},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, decData{types.ZeroDec()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, decData.ZeroValue(), "ZeroValue()")
		})
	}
}
