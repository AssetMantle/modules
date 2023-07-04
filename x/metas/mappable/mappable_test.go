// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/data"
	"github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/metas/key"
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
		{"+ve", args{base.NewStringData("data")}, &Mappable{base.NewStringData("data").ToAnyData().(*base.AnyData)}},
		{"+ve with nil", args{}, &Mappable{}},
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
		{"+ve", &Mappable{}},
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
		{"+ve", fields{base.NewStringData("Data")}, key.NewKey(baseIDs.GenerateDataID(base.NewStringData("Data")))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mappable := &Mappable{
				Data: tt.fields.Data.ToAnyData().(*base.AnyData),
			}
			if got := mappable.GenerateKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateKey() = %v, want %v", got, tt.want)
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
			ma := &Mappable{
				Data: tt.fields.Data.ToAnyData().(*base.AnyData),
			}
			ma.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
