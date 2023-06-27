// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/constants"
)

func createTestInput() *baseIDs.ClassificationID {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	return classificationID.(*baseIDs.ClassificationID)
}

func TestNewKey(t *testing.T) {
	type args struct {
		classificationID ids.ClassificationID
	}
	tests := []struct {
		name string
		args args
		want helpers.Key
	}{
		{"+ve", args{createTestInput()}, &Key{createTestInput()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKey(tt.args.classificationID); !reflect.DeepEqual(got, tt.want) {
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
		want    *Key
		wantErr bool
	}{
		{"+ve", args{NewKey(createTestInput())}, &Key{createTestInput()}, false},
		{"+ve MetaDataError", args{}, &Key{}, true},
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
		ClassificationID *baseIDs.ClassificationID
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
		{"+ve", fields{createTestInput()}, args{&Key{createTestInput()}}, true},
		{"+ve", fields{createTestInput()}, args{KeyPrototype()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := key.Equals(tt.args.compareKey); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_GenerateStoreKeyBytes(t *testing.T) {
	type fields struct {
		ClassificationID *baseIDs.ClassificationID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{createTestInput()}, constants.ModuleStoreKeyPrefix.GenerateStoreKey((&Key{createTestInput()}).GenerateStoreKeyBytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := key.GenerateStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStoreKeyBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_IsPartial(t *testing.T) {
	type fields struct {
		ClassificationID *baseIDs.ClassificationID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{createTestInput()}, false},
		{"-ve", fields{baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_RegisterCodec(t *testing.T) {
	type fields struct {
		ClassificationID *baseIDs.ClassificationID
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{createTestInput()}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ke := &Key{
				ClassificationID: tt.fields.ClassificationID,
			}
			ke.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
