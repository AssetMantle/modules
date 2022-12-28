// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
		{"+ve with nil", args{}, &DecData{}},
		{"+ve with zero dec", args{types.ZeroDec()}, &DecData{types.ZeroDec()}},
		{"+ve", args{types.NewDec(100)}, &DecData{types.NewDec(100)}},
		{"+ve with -ve Dec", args{types.NewDec(-100)}, &DecData{types.NewDec(-100)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewDecData(tt.args.value), "NewDecData(%v)", tt.args.value)
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
		{"+ve with nil", fields{}, []byte{0x1}}, // TODO: Update test after fixing the bug
		{"+ve with zero dec", fields{types.ZeroDec()}, DecData{types.ZeroDec()}.Bytes()},
		{"+ve", fields{types.NewDec(100)}, DecData{types.NewDec(100)}.Bytes()},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, DecData{types.NewDec(-100)}.Bytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := DecData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, decData.Bytes(), "Bytes()")
		})
	}
}

func Test_decData_Compare(t *testing.T) {
	require.Panics(t, func() {
		DecData{}.Compare(nil)
	})
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
		{"+ve with nil", fields{}, args{&DecData{}}, 0},
		{"+ve with nil", fields{types.Dec{}}, args{&DecData{types.Dec{}}}, 0},
		{"+ve with zero dec", fields{types.ZeroDec()}, args{&DecData{types.ZeroDec()}}, 0},
		{"+ve", fields{types.NewDec(100)}, args{&DecData{types.NewDec(100)}}, 0},
		{"-ve", fields{types.NewDec(-100)}, args{&DecData{types.NewDec(100)}}, -1},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, args{&DecData{types.NewDec(-100)}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := DecData{
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
		name      string
		fields    fields
		want      ids.HashID
		wantPanic bool
	}{
		{"panic case with nil", fields{types.Dec{}}, baseIDs.GenerateHashID(), false},
		{"+ve with zero dec", fields{types.ZeroDec()}, baseIDs.GenerateHashID(), false},
		{"+ve", fields{types.NewDec(100)}, baseIDs.GenerateHashID(DecData{types.NewDec(100)}.Bytes()), false},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, baseIDs.GenerateHashID(DecData{types.NewDec(-100)}.Bytes()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := DecData{
				Value: tt.fields.Value,
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					decData.GenerateHashID()
				})
			} else {
				assert.Equalf(t, tt.want, decData.GenerateHashID(), "GenerateHashID()")
			}

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
		{"+ve with nil", fields{}, DecData{}.Value},
		{"+ve with zero dec", fields{types.ZeroDec()}, DecData{types.ZeroDec()}.Value},
		{"+ve", fields{types.NewDec(100)}, DecData{types.NewDec(100)}.Value},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, DecData{types.NewDec(-100)}.Value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := DecData{
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
		name      string
		fields    fields
		want      ids.DataID
		wantPanic bool
	}{
		{"panic case with nil", fields{types.Dec{}}, nil, true}, // TODO: Check whether planned panic in GenerateDataID is expected behaviour
		{"+ve with zero dec", fields{types.ZeroDec()}, baseIDs.GenerateDataID(&DecData{types.ZeroDec()}), false},
		{"+ve", fields{types.NewDec(100)}, baseIDs.GenerateDataID(&DecData{types.NewDec(100)}), false},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, baseIDs.GenerateDataID(&DecData{types.NewDec(-100)}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := DecData{
				Value: tt.fields.Value,
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					decData.GetID()
				})
			} else if got := decData.GetID(); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve with nil", fields{}, dataConstants.DecDataID},
		{"+ve with zero dec", fields{types.ZeroDec()}, dataConstants.DecDataID},
		{"+ve", fields{types.NewDec(100)}, dataConstants.DecDataID},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, dataConstants.DecDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := DecData{
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
		{"+ve with nil", fields{}, DecData{}.Value.String()},
		{"+ve with zero dec", fields{types.ZeroDec()}, DecData{types.ZeroDec()}.Value.String()},
		{"+ve", fields{types.NewDec(100)}, DecData{types.NewDec(100)}.Value.String()},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, DecData{types.NewDec(-100)}.Value.String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := DecData{
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
		{"+ve with nil", fields{}, &DecData{types.ZeroDec()}},
		{"+ve with zero dec", fields{types.ZeroDec()}, &DecData{types.ZeroDec()}},
		{"+ve", fields{types.NewDec(100)}, &DecData{types.ZeroDec()}},
		{"+ve with -ve Dec", fields{types.NewDec(-100)}, &DecData{types.ZeroDec()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := DecData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, decData.ZeroValue(), "ZeroValue()")
		})
	}
}
