// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"reflect"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
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
		value sdkTypes.Dec
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		{"+ve with nil", args{}, decData{}},
		{"+ve with zero dec", args{sdkTypes.ZeroDec()}, decData{sdkTypes.ZeroDec()}},
		{"+ve", args{sdkTypes.NewDec(100)}, decData{sdkTypes.NewDec(100)}},
		{"+ve with -ve Dec", args{sdkTypes.NewDec(-100)}, decData{sdkTypes.NewDec(-100)}},
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
		{"+ve with nil", args{decData{}}, decData{}, assert.NoError},
		{"+ve with nil", args{decData{sdkTypes.Dec{}}}, decData{}, assert.NoError},
		{"+ve with zero dec", args{decData{sdkTypes.ZeroDec()}}, decData{sdkTypes.ZeroDec()}, assert.NoError},
		{"+ve", args{decData{sdkTypes.NewDec(100)}}, decData{sdkTypes.NewDec(100)}, assert.NoError},
		{"+ve with -ve Dec", args{decData{sdkTypes.NewDec(-100)}}, decData{sdkTypes.NewDec(-100)}, assert.NoError},
		{"-ve with nil", args{nil}, decData{}, assert.Error},
		{"-ve stringData", args{stringData{"testData"}}, decData{}, assert.Error},
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
		Value sdkTypes.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve with nil", fields{}, []byte{0x1}}, // TODO: Update test after fixing the bug
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, decData{sdkTypes.ZeroDec()}.Bytes()},
		{"+ve", fields{sdkTypes.NewDec(100)}, decData{sdkTypes.NewDec(100)}.Bytes()},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, decData{sdkTypes.NewDec(-100)}.Bytes()},
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
	require.Panics(t, func() {
		decData{}.Compare(nil)
	})
	type fields struct {
		Value sdkTypes.Dec
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
		{"+ve with nil", fields{}, args{decData{}}, 0},
		{"+ve with nil", fields{sdkTypes.Dec{}}, args{decData{sdkTypes.Dec{}}}, 0},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, args{decData{sdkTypes.ZeroDec()}}, 0},
		{"+ve", fields{sdkTypes.NewDec(100)}, args{decData{sdkTypes.NewDec(100)}}, 0},
		{"-ve", fields{sdkTypes.NewDec(-100)}, args{decData{sdkTypes.NewDec(100)}}, -1},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, args{decData{sdkTypes.NewDec(-100)}}, 0},
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
		Value sdkTypes.Dec
	}
	tests := []struct {
		name      string
		fields    fields
		want      ids.HashID
		wantPanic bool
	}{
		{"panic case with nil", fields{sdkTypes.Dec{}}, baseIDs.GenerateHashID(), false},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, baseIDs.GenerateHashID(), false},
		{"+ve", fields{sdkTypes.NewDec(100)}, baseIDs.GenerateHashID(decData{sdkTypes.NewDec(100)}.Bytes()), false},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, baseIDs.GenerateHashID(decData{sdkTypes.NewDec(-100)}.Bytes()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
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
		Value sdkTypes.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   sdkTypes.Dec
	}{
		{"+ve with nil", fields{}, decData{}.Value},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, decData{sdkTypes.ZeroDec()}.Value},
		{"+ve", fields{sdkTypes.NewDec(100)}, decData{sdkTypes.NewDec(100)}.Value},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, decData{sdkTypes.NewDec(-100)}.Value},
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
		Value sdkTypes.Dec
	}
	tests := []struct {
		name      string
		fields    fields
		want      ids.DataID
		wantPanic bool
	}{
		{"panic case with nil", fields{sdkTypes.Dec{}}, nil, true}, // TODO: Check whether planned panic in GenerateDataID is expected behaviour
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, baseIDs.GenerateDataID(decData{sdkTypes.ZeroDec()}), false},
		{"+ve", fields{sdkTypes.NewDec(100)}, baseIDs.GenerateDataID(decData{sdkTypes.NewDec(100)}), false},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, baseIDs.GenerateDataID(decData{sdkTypes.NewDec(-100)}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
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
		Value sdkTypes.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.StringID
	}{
		{"+ve with nil", fields{}, dataConstants.DecDataID},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, dataConstants.DecDataID},
		{"+ve", fields{sdkTypes.NewDec(100)}, dataConstants.DecDataID},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, dataConstants.DecDataID},
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
		Value sdkTypes.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve with nil", fields{}, decData{}.Value.String()},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, decData{sdkTypes.ZeroDec()}.Value.String()},
		{"+ve", fields{sdkTypes.NewDec(100)}, decData{sdkTypes.NewDec(100)}.Value.String()},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, decData{sdkTypes.NewDec(-100)}.Value.String()},
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
		Value sdkTypes.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		{"+ve with nil", fields{}, decData{sdkTypes.ZeroDec()}},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, decData{sdkTypes.ZeroDec()}},
		{"+ve", fields{sdkTypes.NewDec(100)}, decData{sdkTypes.ZeroDec()}},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, decData{sdkTypes.ZeroDec()}},
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
