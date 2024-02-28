// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
)

func createTestInput() *baseIDs.IdentityID {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testIdentity := baseIDs.NewIdentityID(classificationID, immutables)

	return testIdentity.(*baseIDs.IdentityID)
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
		{"+ve", args{testIdentity}, &Key{IdentityID: testIdentity}},
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
	testIdentity := createTestInput()

	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *Key
		wantErr bool
	}{
		{"+ve", args{nil}, &Key{nil}, true},
		{"-ve", args{NewKey(nil)}, &Key{nil}, false},
		{"-ve", args{testIdentity}, &Key{nil}, true},
		{"+ve", args{NewKey(testIdentity)}, &Key{testIdentity}, false},
		{"-ve", args{baseIDs.NewStringID("StringID")}, &Key{}, true},
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
		IdentityID *baseIDs.IdentityID
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
		{"+ve", fields{nil}, args{&Key{nil}}, true},
		{"-ve", fields{nil}, args{&Key{testIdentity}}, false},
		{"-ve", fields{testIdentity}, args{&Key{nil}}, false},
		{"+ve", fields{testIdentity}, args{&Key{testIdentity}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
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
		IdentityID *baseIDs.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{testIdentity}, (&Key{testIdentity}).GeneratePrefixedStoreKeyBytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				IdentityID: tt.fields.IdentityID,
			}
			if got := key.GeneratePrefixedStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratePrefixedStoreKeyBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_IsPartial(t *testing.T) {
	testIdentity := createTestInput()

	type fields struct {
		IdentityID *baseIDs.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testIdentity}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				IdentityID: tt.fields.IdentityID,
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}
