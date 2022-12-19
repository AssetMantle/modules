// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	base2 "github.com/AssetMantle/modules/schema/ids/base"
)

func TestNewMappable(t *testing.T) {
	type args struct {
		data data.Data
	}
	tests := []struct {
		name string
		args args
		want helpers.Mappable
	}{
		{"+ve", args{base.NewStringData("data")}, mappable{base.NewStringData("data")}},
		{"+ve with nil", args{}, mappable{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMappable(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMappable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Mappable
	}{
		{"+ve", mappable{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mappable_GetKey(t *testing.T) {
	type fields struct {
		Data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.Key
	}{
		{"+ve", fields{base.NewStringData("Data")}, key.NewKey(base2.NewDataID(base.NewStringData("Data")))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mappable := mappable{
				Data: tt.fields.Data,
			}
			if got := mappable.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mappable_RegisterCodec(t *testing.T) {
	type fields struct {
		Data data.Data
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{base.NewStringData("Data")}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := mappable{
				Data: tt.fields.Data,
			}
			ma.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
