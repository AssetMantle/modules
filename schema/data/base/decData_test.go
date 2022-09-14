// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
	"github.com/cosmos/cosmos-sdk/types"
	"math"
	"reflect"
	"strconv"
	"testing"
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
		{"Test for zeroDec", args{
			types.ZeroDec(),
		}, decData{
			types.ZeroDec(),
		}},
		{"Test for NewDec with precision", args{
			types.NewDecWithPrec(int64(10), 2),
		}, decData{
			types.NewDecWithPrec(int64(10), 2),
		}},
		{"Test for NewDec", args{
			types.NewDec(10),
		}, decData{
			types.NewDec(10),
		}},
		{"Test for NewDec with -ve int", args{
			types.NewDec(-10),
		}, decData{
			types.NewDec(-10),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDecData(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDecData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadDecData(t *testing.T) {
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.Data
		wantErr bool
	}{
		{"Empty String", args{""}, decData{}.ZeroValue(), false},
		{"Empty String with -ve", args{"-"}, decData{}.ZeroValue(), true},
		{"+ve case with +ve int", args{"10"}, NewDecData(types.NewDec(10)), false},
		{"+ve case with -ve int", args{"-10"}, NewDecData(types.NewDec(-10)), false},
		{"+ve case with 0", args{"0"}, NewDecData(types.NewDec(0)), false},
		{"+ve case with Max +ve int64", args{strconv.FormatInt(math.MaxInt64, 10)}, NewDecData(types.NewDec(math.MaxInt64)), false},
		{"+ve case with Min -ve int64", args{strconv.FormatInt(math.MinInt64, 10)}, NewDecData(types.NewDec(math.MinInt64)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDecData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDecData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDecData() got = %v, want %v", got, tt.want)
			}
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
		wantErr bool
	}{
		{"ZeroDecData", args{decData{
			types.ZeroDec(),
		}}, decData{
			types.ZeroDec(),
		}, false},
		{"NewDec with +ve int", args{decData{
			types.NewDec(100),
		}}, decData{
			types.NewDec(100),
		}, false},
		{"NewDec with -ve int", args{decData{
			types.NewDec(-100),
		}}, decData{
			types.NewDec(-100),
		}, false},
		{"-ve case", args{baseIDs.NewStringID("ID")}, decData{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decDataFromInterface(tt.args.listable)
			if (err != nil) != tt.wantErr {
				t.Errorf("decDataFromInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decDataFromInterface() got = %v, want %v", got, tt.want)
			}
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
		{"Test for ZeroDec", fields{
			types.ZeroDec(),
		}, args{decData{
			types.ZeroDec(),
		}.ZeroValue()}, 0},
		{"Test for GT case", fields{
			types.NewDec(100),
		}, args{decData{
			types.NewDec(-100),
		}.ZeroValue()}, 1},
		{"Test for LT case", fields{
			types.NewDec(-100),
		}, args{decData{
			types.NewDec(100),
		}.ZeroValue()}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			if got := decData.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decData_GenerateHash(t *testing.T) {
	type fields struct {
		Value types.Dec
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		{"ZeroValue Test", fields{types.ZeroDec()}, baseIDs.NewStringID("")},
		{"+ve Value Test", fields{types.NewDec(100)}, baseIDs.NewStringID(stringUtilities.Hash(decData{types.NewDec(100)}.Value.String()))},
		{"-ve Value Test", fields{types.NewDec(-100)}, baseIDs.NewStringID(stringUtilities.Hash(decData{types.NewDec(-100)}.Value.String()))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			if got := decData.GenerateHash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHash() = %v, want %v", got, tt.want)
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
		{"Test for ZeroDec", fields{types.ZeroDec()}, types.ZeroDec()},
		{"Test for some +ve Data", fields{types.NewDec(100)}, types.NewDec(100)},
		{"Test for some -ve Data", fields{types.NewDec(-100)}, types.NewDec(-100)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			if got := decData.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
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
		{"Test with ZeroDec data", fields{types.ZeroDec()}, baseIDs.NewDataID(decData{types.ZeroDec()})},
		{"Test for some +ve Data", fields{types.NewDec(100)}, baseIDs.NewDataID(decData{types.NewDec(100)})},
		{"Test for some -ve Data", fields{types.NewDec(-100)}, baseIDs.NewDataID(decData{types.NewDec(-100)})},
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
		want   ids.ID
	}{
		{"Test with ZeroDec data", fields{types.ZeroDec()}, idsConstants.DecDataID},
		{"Test for some +ve Data", fields{types.NewDec(100)}, idsConstants.DecDataID},
		{"Test for some -ve Data", fields{types.NewDec(-100)}, idsConstants.DecDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			if got := decData.GetType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
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
		{"Test with ZeroDec data", fields{types.ZeroDec()}, decData{types.ZeroDec()}.Value.String()},
		{"Test for some +ve Data", fields{types.NewDec(100)}, decData{types.NewDec(100)}.Value.String()},
		{"Test for some -ve Data", fields{types.NewDec(-100)}, decData{types.NewDec(-100)}.Value.String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			if got := decData.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
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
		{"Test with +ve Value", fields{types.NewDec(100)}, NewDecData(types.ZeroDec())},
		{"Test with -ve Value", fields{types.NewDec(-100)}, NewDecData(types.ZeroDec())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decData := decData{
				Value: tt.fields.Value,
			}
			if got := decData.ZeroValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
