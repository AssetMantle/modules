// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
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
		{"+ve with nil", args{}, &DecData{}},
		{"+ve with zero dec", args{sdkTypes.ZeroDec()}, &DecData{sdkTypes.ZeroDec()}},
		{"+ve", args{sdkTypes.NewDec(100)}, &DecData{sdkTypes.NewDec(100)}},
		{"+ve with -ve Dec", args{sdkTypes.NewDec(-100)}, &DecData{sdkTypes.NewDec(-100)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewDecData(tt.args.value), "NewDecData(%v)", tt.args.value)
		})
	}
}

//func Test_decDataFromInterface(t *testing.T) {
//	type args struct {
//		listable traits.Listable
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *DecData
//		wantErr assert.ErrorAssertionFunc
//	}{
//		{"+ve with nil", args{&DecData{}}, &DecData{}, assert.NoError},
//		{"+ve with nil", args{&DecData{sdkTypes.Dec{}}}, &DecData{}, assert.NoError},
//		{"+ve with zero dec", args{&DecData{sdkTypes.ZeroDec()}}, &DecData{sdkTypes.ZeroDec()}, assert.NoError},
//		{"+ve", args{&DecData{sdkTypes.NewDec(100)}}, &DecData{sdkTypes.NewDec(100)}, assert.NoError},
//		{"+ve with -ve Dec", args{&DecData{sdkTypes.NewDec(-100)}}, &DecData{sdkTypes.NewDec(-100)}, assert.NoError},
//		{"-ve with nil", args{nil}, &DecData{}, assert.Error},
//		{"-ve stringData", args{&StringData{"testData"}}, &DecData{}, assert.Error},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := decDataFromInterface(tt.args.listable)
//			if !tt.wantErr(t, err, fmt.Sprintf("decDataFromInterface(%v)", tt.args.listable)) {
//				return
//			}
//			assert.Equalf(t, tt.want, got, "decDataFromInterface(%v)", tt.args.listable)
//		})
//	}
//}

func Test_decData_Bytes(t *testing.T) {
	type fields struct {
		Value sdkTypes.Dec
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"+ve with nil", fields{}, []byte{0x1}, true}, // TODO: Update test after fixing the bug
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, (&DecData{sdkTypes.ZeroDec()}).Bytes(), false},
		{"+ve", fields{sdkTypes.NewDec(100)}, (&DecData{sdkTypes.NewDec(100)}).Bytes(), false},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, (&DecData{sdkTypes.NewDec(-100)}).Bytes(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := &DecData{
				Value: tt.fields.Value,
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			assert.Equalf(t, tt.want, decData.Bytes(), "Bytes()")
		})
	}
}

func Test_decData_Compare(t *testing.T) {
	require.Panics(t, func() {
		(&DecData{}).Compare(nil)
	})
	type fields struct {
		Value sdkTypes.Dec
	}
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"panic with nil", fields{}, args{&DecData{}}, 0, true},
		{"MetaDataError with nil", fields{sdkTypes.Dec{}}, args{&DecData{sdkTypes.Dec{}}}, 0, true},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, args{&DecData{sdkTypes.ZeroDec()}}, 0, false},
		{"+ve", fields{sdkTypes.NewDec(100)}, args{&DecData{sdkTypes.NewDec(100)}}, 0, false},
		{"-ve", fields{sdkTypes.NewDec(-100)}, args{&DecData{sdkTypes.NewDec(100)}}, -1, false},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, args{&DecData{sdkTypes.NewDec(-100)}}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := &DecData{
				Value: tt.fields.Value,
			}
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", r, tt.wantErr)
				}
			}()
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
		{"panic case with nil", fields{sdkTypes.Dec{}}, baseIDs.GenerateHashID(), true},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, baseIDs.GenerateHashID(), false},
		{"+ve", fields{sdkTypes.NewDec(100)}, baseIDs.GenerateHashID((&DecData{sdkTypes.NewDec(100)}).Bytes()), false},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, baseIDs.GenerateHashID((&DecData{sdkTypes.NewDec(-100)}).Bytes()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := &DecData{
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
		{"+ve with nil", fields{}, (&DecData{}).Value},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, (&DecData{sdkTypes.ZeroDec()}).Value},
		{"+ve", fields{sdkTypes.NewDec(100)}, (&DecData{sdkTypes.NewDec(100)}).Value},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, (&DecData{sdkTypes.NewDec(-100)}).Value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := &DecData{
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
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, baseIDs.GenerateDataID(&DecData{sdkTypes.ZeroDec()}), false},
		{"+ve", fields{sdkTypes.NewDec(100)}, baseIDs.GenerateDataID(&DecData{sdkTypes.NewDec(100)}), false},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, baseIDs.GenerateDataID(&DecData{sdkTypes.NewDec(-100)}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := &DecData{
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
			decData := &DecData{
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
		{"+ve with nil", fields{}, (&DecData{}).Value.String()},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, (&DecData{sdkTypes.ZeroDec()}).Value.String()},
		{"+ve", fields{sdkTypes.NewDec(100)}, (&DecData{sdkTypes.NewDec(100)}).Value.String()},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, (&DecData{sdkTypes.NewDec(-100)}).Value.String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := &DecData{
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
		{"+ve with nil", fields{}, &DecData{sdkTypes.ZeroDec()}},
		{"+ve with zero dec", fields{sdkTypes.ZeroDec()}, &DecData{sdkTypes.ZeroDec()}},
		{"+ve", fields{sdkTypes.NewDec(100)}, &DecData{sdkTypes.ZeroDec()}},
		{"+ve with -ve Dec", fields{sdkTypes.NewDec(-100)}, &DecData{sdkTypes.ZeroDec()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := &DecData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, decData.ZeroValue(), "ZeroValue()")
		})
	}
}
