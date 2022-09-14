// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
	types2 "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"

	"strconv"
	"testing"
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

		{"Test for +ve int", args{baseTypes.NewHeight(100)}, heightData{baseTypes.NewHeight(100)}},
		{"Test for +ve int", args{baseTypes.NewHeight(-100)}, heightData{baseTypes.NewHeight(-100)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeightData(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeightData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadHeightData(t *testing.T) {
	type args struct {
		dataString string
	}
	tests := []struct {
		name        string
		args        args
		want        data.Data
		wantErr     bool
		errorString string
	}{

		{"Test for empty String", args{""}, heightData{}.ZeroValue(), false, "nil"},
		{"Test for some +ve integer", args{"100"}, heightData{baseTypes.NewHeight(100)}, false, ""},
		{"Test for some -ve integer", args{"-100"}, heightData{baseTypes.NewHeight(-100)}, false, ""},
		{"Test for some float", args{"100.5"}, nil, true, fmt.Sprintf("strconv.ParseInt: parsing \"%v\": invalid syntax", 100.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadHeightData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadHeightData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, tt.errorString, err.Error())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadHeightData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightDataFromInterface(t *testing.T) {
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name        string
		args        args
		want        heightData
		wantErr     bool
		errorString string
	}{

		{"Test for empty height data", args{heightData{}}, heightData{}, false, ""},
		{"Test for +ve int height data", args{heightData{baseTypes.NewHeight(100)}}, heightData{baseTypes.NewHeight(100)}, false, ""},
		{"Test for -ve int height data", args{heightData{baseTypes.NewHeight(-100)}}, heightData{baseTypes.NewHeight(-100)}, false, ""},
		{"Test for Other listable Type", args{decData{types2.ZeroDec()}.ZeroValue()}, heightData{}, true, constants.MetaDataError.Error()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := heightDataFromInterface(tt.args.listable)
			if (err != nil) != tt.wantErr {
				t.Errorf("heightDataFromInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, tt.errorString, err.Error())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heightDataFromInterface() got = %v, want %v", got, tt.want)
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

		{"Test for Equal case", fields{baseTypes.NewHeight(100)}, args{heightData{baseTypes.NewHeight(100)}}, 0},
		{"Test for LT case", fields{baseTypes.NewHeight(0)}, args{heightData{baseTypes.NewHeight(100)}}, -1},
		{"Test for GT case", fields{baseTypes.NewHeight(100)}, args{heightData{baseTypes.NewHeight(0)}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := heightData{
				Value: tt.fields.Value,
			}
			if got := heightData.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heightData_GenerateHash(t *testing.T) {
	type fields struct {
		Value types.Height
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{

		{"Test for zero value", fields{baseTypes.NewHeight(0)}, baseIDs.NewStringID("")},
		{"Test for -ve value", fields{baseTypes.NewHeight(-100)}, baseIDs.NewStringID(stringUtilities.Hash(strconv.FormatInt(heightData{baseTypes.NewHeight(-100)}.Value.Get(), 10)))},
		{"Test for +ve value", fields{baseTypes.NewHeight(100)}, baseIDs.NewStringID(stringUtilities.Hash(strconv.FormatInt(heightData{baseTypes.NewHeight(100)}.Value.Get(), 10)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := heightData{
				Value: tt.fields.Value,
			}
			if got := heightData.GenerateHash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHash() = %v, want %v", got, tt.want)
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

		{"Test for zero value", fields{baseTypes.NewHeight(0)}, heightData{baseTypes.NewHeight(0)}.Value},
		{"Test for +ve value", fields{baseTypes.NewHeight(100)}, heightData{baseTypes.NewHeight(100)}.Value},
		{"Test for -ve value", fields{baseTypes.NewHeight(-100)}, heightData{baseTypes.NewHeight(-100)}.Value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := heightData{
				Value: tt.fields.Value,
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

		{"Test for zero value", fields{baseTypes.NewHeight(0)}, baseIDs.NewDataID(heightData{baseTypes.NewHeight(0)})},
		{"Test for +ve value", fields{baseTypes.NewHeight(100)}, baseIDs.NewDataID(heightData{baseTypes.NewHeight(100)})},
		{"Test for -ve value", fields{baseTypes.NewHeight(-100)}, baseIDs.NewDataID(heightData{baseTypes.NewHeight(-100)})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := heightData{
				Value: tt.fields.Value,
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
			heightData := heightData{
				Value: tt.fields.Value,
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

		{"Test for zero value", fields{baseTypes.NewHeight(0)}, strconv.FormatInt(heightData{baseTypes.NewHeight(0)}.Value.Get(), 10)},
		{"Test for +ve value", fields{baseTypes.NewHeight(100)}, strconv.FormatInt(heightData{baseTypes.NewHeight(100)}.Value.Get(), 10)},
		{"Test for -ve value", fields{baseTypes.NewHeight(-100)}, strconv.FormatInt(heightData{baseTypes.NewHeight(-100)}.Value.Get(), 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := heightData{
				Value: tt.fields.Value,
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

		{"Test for zero value", fields{baseTypes.NewHeight(0)}, heightData{baseTypes.NewHeight(0)}},
		{"Test for +ve Int value", fields{baseTypes.NewHeight(100)}, heightData{baseTypes.NewHeight(0)}},
		{"Test for -ve Int value", fields{baseTypes.NewHeight(-100)}, heightData{baseTypes.NewHeight(0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heightData := heightData{
				Value: tt.fields.Value,
			}
			if got := heightData.ZeroValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
