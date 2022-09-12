// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
	"github.com/stretchr/testify/assert"
	"testing"
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

		{"Test for some id", args{baseIDs.NewID("100")}, idData{baseIDs.NewID("100")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewIDData(tt.args.value), "NewIDData(%v)", tt.args.value)
		})
	}
}

func TestReadIDData(t *testing.T) {
	type args struct {
		idData string
	}
	tests := []struct {
		name    string
		args    args
		want    data.Data
		wantErr assert.ErrorAssertionFunc
	}{

		{"Test for some +ve int id", args{"100"}, NewIDData(baseIDs.NewID("100")), assert.NoError},
		{"Test for some -ve int id", args{"-100"}, NewIDData(baseIDs.NewID("-100")), assert.NoError},
		{"Test for zero id", args{"0"}, NewIDData(baseIDs.NewID("0")), assert.NoError},
		{"Test for string with special char id", args{"10-10"}, NewIDData(baseIDs.NewID("10-10")), assert.NoError},   // wrong
		{"Test for string with special char id", args{"10%10"}, NewIDData(baseIDs.NewID("10%10")), assert.NoError},   // wrong
		{"Test for string with special char id", args{"10%`10"}, NewIDData(baseIDs.NewID("10%`10")), assert.NoError}, // wrong
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadIDData(tt.args.idData)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadIDData(%v)", tt.args.idData)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadIDData(%v)", tt.args.idData)
		})
	}
}

func Test_idDataFromInterface(t *testing.T) {
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		args    args
		want    idData
		wantErr assert.ErrorAssertionFunc
	}{

		{"+ve Unit test", args{idData{baseIDs.NewID("100")}}, idData{baseIDs.NewID("100")}, assert.NoError},
		{"-ve Unit test", args{heightData{baseTypes.NewHeight(100)}}, idData{}, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := idDataFromInterface(tt.args.listable)
			if !tt.wantErr(t, err, fmt.Sprintf("idDataFromInterface(%v)", tt.args.listable)) {
				return
			}
			assert.Equalf(t, tt.want, got, "idDataFromInterface(%v)", tt.args.listable)
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

		{"Test for Equal case", fields{baseIDs.NewID("100")}, args{idData{baseIDs.NewID("100")}}, 0},
		{"Test for LT case", fields{baseIDs.NewID("100")}, args{idData{baseIDs.NewID("0")}}, 1},
		{"Test for GT case", fields{baseIDs.NewID("0")}, args{idData{baseIDs.NewID("100")}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_idData_GenerateHash(t *testing.T) {
	type fields struct {
		Value ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{

		{"ZeroValue Test", fields{baseIDs.NewID("")}, baseIDs.NewID("")},
		{"+ve Value Test", fields{baseIDs.NewID("100")}, baseIDs.NewID(stringUtilities.Hash(idData{baseIDs.NewID("100")}.Value.String()))},
		{"-ve Value Test", fields{baseIDs.NewID("-100")}, baseIDs.NewID(stringUtilities.Hash(idData{baseIDs.NewID("-100")}.Value.String()))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.GenerateHash(), "GenerateHash()")
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

		{"Test for zero value", fields{baseIDs.NewID("0")}, idData{baseIDs.NewID("0")}.Value},
		{"Test for +ve value", fields{baseIDs.NewID("100")}, idData{baseIDs.NewID("100")}.Value},
		{"Test for -ve value", fields{baseIDs.NewID("-100")}, idData{baseIDs.NewID("-100")}.Value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.Get(), "Get()")
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

		{"Test for zero value", fields{baseIDs.NewID("0")}, baseIDs.NewDataID(idData{baseIDs.NewID("0")})},
		{"Test for +ve value", fields{baseIDs.NewID("100")}, baseIDs.NewDataID(idData{baseIDs.NewID("100")})},
		{"Test for -ve value", fields{baseIDs.NewID("-100")}, baseIDs.NewDataID(idData{baseIDs.NewID("-100")})},
		{"Test for special char value", fields{baseIDs.NewID("%#100")}, baseIDs.NewDataID(idData{baseIDs.NewID("%#100")})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
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
		want   ids.ID
	}{

		{"Test for zero value", fields{baseIDs.NewID("0")}, idsConstants.IDDataID},
		{"Test for +ve value", fields{baseIDs.NewID("100")}, idsConstants.IDDataID},
		{"Test for -ve value", fields{baseIDs.NewID("-100")}, idsConstants.IDDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.GetType(), "GetType()")
		})
	}
}

func Test_idData_String(t *testing.T) {
	type fields struct {
		Value ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{"Test for zero value", fields{baseIDs.NewID("0")}, "0"},
		{"Test for +ve value", fields{baseIDs.NewID("100")}, "100"},
		{"Test for -ve value", fields{baseIDs.NewID("-100")}, "-100"},
		{"Test for special char value", fields{baseIDs.NewID("%#100")}, "%#100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.String(), "String()")
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

		{"Test for zero value", fields{baseIDs.NewID("0")}, NewIDData(baseIDs.NewID(""))},
		{"Test for +ve value", fields{baseIDs.NewID("100")}, NewIDData(baseIDs.NewID(""))},
		{"Test for -ve value", fields{baseIDs.NewID("-100")}, NewIDData(baseIDs.NewID(""))},
		{"Test for special char value", fields{baseIDs.NewID("%#100")}, NewIDData(baseIDs.NewID(""))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idData := idData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, idData.ZeroValue(), "ZeroValue()")
		})
	}
}
