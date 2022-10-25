// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	base2 "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func createTestInput() ids.IdentityID {
	immutables := baseQualified.NewImmutables(base2.NewPropertyList(baseProperties.NewMesaProperty(base.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(base2.NewPropertyList(baseProperties.NewMesaProperty(base.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := base.NewClassificationID(immutables, mutables)
	testIdentity := base.NewIdentityID(classificationID, immutables)

	return testIdentity
}

func TestNewKey(t *testing.T) {
	testIdentity := createTestInput()
	type args struct {
		IdentityID ids.IdentityID
	}
	tests := []struct {
		name string
		args args
		want helpers.Key
	}{
		// TODO: Add test cases.
		{"+ve", args{testIdentity}, key{IdentityID: testIdentity}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKey(tt.args.IdentityID); !reflect.DeepEqual(got, tt.want) {
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
		// TODO: Add test cases.
		{"+ve", key{}},
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
	testIdentity := createTestInput()

	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    key
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", args{}, key{}, true},
		{"+ve", args{NewKey(testIdentity)}, key{testIdentity}, false},
		{"-ve", args{base.NewStringID("StringID")}, key{}, true},
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
	testIdentity := createTestInput()
	type fields struct {
		IdentityID ids.IdentityID
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
		// TODO: Add test cases.
		{"+ve", fields{testIdentity}, args{key{testIdentity}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := key{
				IdentityID: tt.fields.IdentityID,
			}
			if got := key.Equals(tt.args.compareKey); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_GenerateStoreKeyBytes(t *testing.T) {
	testIdentity := createTestInput()

	type fields struct {
		IdentityID ids.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{testIdentity}, module.StoreKeyPrefix.GenerateStoreKey(key{testIdentity}.Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := key{
				IdentityID: tt.fields.IdentityID,
			}
			if got := key.GenerateStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStoreKeyBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_IsPartial(t *testing.T) {
	testIdentity := createTestInput()

	type fields struct {
		IdentityID ids.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{testIdentity}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := key{
				IdentityID: tt.fields.IdentityID,
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_RegisterCodec(t *testing.T) {
	testIdentity := createTestInput()

	type fields struct {
		IdentityID ids.IdentityID
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"+ve", fields{testIdentity}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ke := key{
				IdentityID: tt.fields.IdentityID,
			}
			ke.RegisterCodec(tt.args.codec)
		})
	}
}
