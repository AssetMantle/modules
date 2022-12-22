// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/internal/module"
	base2 "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

var (
	testDataID  = baseIDs.GenerateDataID(base2.NewStringData("Data"))
	testDataID1 = baseIDs.PrototypeDataID()
)

func TestNewKey(t *testing.T) {
	type args struct {
		dataID ids.DataID
	}
	tests := []struct {
		name string
		args args
		want helpers.Key
	}{
		{"+ve", args{testDataID.(*baseIDs.DataID)}, &Key{testDataID.(*baseIDs.DataID)}},
		{"+ve", args{testDataID1}, &Key{testDataID1.(*baseIDs.DataID)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKey(tt.args.dataID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Key
	}{
		{"+ve", &Key{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keyFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    Key
		wantErr bool
	}{
		{"+ve", args{NewKey(testDataID)}, Key{testDataID.(*baseIDs.DataID)}, false},
		{"+ve with nil", args{NewKey(testDataID1)}, Key{testDataID1.(*baseIDs.DataID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := keyFromInterface(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("keyFromInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keyFromInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_Equals(t *testing.T) {
	type fields struct {
		DataID *baseIDs.DataID
	}
	type args struct {
		compareKey helpers.Key
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"+ve", fields{testDataID.(*baseIDs.DataID)}, args{NewKey(testDataID)}, true},
		{"+ve", fields{testDataID.(*baseIDs.DataID)}, args{NewKey(testDataID1)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				DataId: tt.fields.DataID,
			}
			if got := key.Equals(tt.args.compareKey); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_GenerateStoreKeyBytes(t *testing.T) {
	type fields struct {
		DataID *baseIDs.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{testDataID.(*baseIDs.DataID)}, module.StoreKeyPrefix.GenerateStoreKey(testDataID.Bytes())},
		{"+ve with nil", fields{testDataID1.(*baseIDs.DataID)}, module.StoreKeyPrefix.GenerateStoreKey(testDataID1.Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				DataId: tt.fields.DataID,
			}
			if got := key.GenerateStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStoreKeyBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_IsPartial(t *testing.T) {
	type fields struct {
		DataID *baseIDs.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testDataID.(*baseIDs.DataID)}, false},
		{"-ve", fields{testDataID1.(*baseIDs.DataID)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				DataId: tt.fields.DataID,
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_RegisterCodec(t *testing.T) {
	type fields struct {
		DataID *baseIDs.DataID
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testDataID.(*baseIDs.DataID)}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				DataId: tt.fields.DataID,
			}
			key.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
