// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/AssetMantle/modules/helpers"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"reflect"
	"testing"
)

func createTestData() *baseIDs.MaintainerID {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	// testIdentityID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testMaintainerID := baseIDs.NewMaintainerID(testClassificationID, immutables)
	return testMaintainerID.(*baseIDs.MaintainerID)
}

func TestNewKey(t *testing.T) {
	type args struct {
		maintainerID ids.MaintainerID
	}
	tests := []struct {
		name string
		args args
		want helpers.Key
	}{
		{"+ve with nil", args{}, &Key{}},
		{"+ve", args{createTestData()}, &Key{createTestData()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKey(tt.args.maintainerID); !reflect.DeepEqual(got, tt.want) {
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

func Test_key_Equals(t *testing.T) {
	type fields struct {
		MaintainerID *baseIDs.MaintainerID
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
		{"+ve", fields{createTestData()}, args{NewKey(createTestData())}, true},
		{"-ve", fields{createTestData()}, args{NewKey(baseIDs.PrototypeMaintainerID())}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				MaintainerID: tt.fields.MaintainerID,
			}
			if got := key.Equals(tt.args.compareKey); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_GenerateStoreKeyBytes(t *testing.T) {
	type fields struct {
		MaintainerID *baseIDs.MaintainerID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{createTestData()}, (&Key{createTestData()}).GeneratePrefixedStoreKeyBytes()},
		{"-ve", fields{baseIDs.PrototypeMaintainerID().(*baseIDs.MaintainerID)}, (&Key{baseIDs.PrototypeMaintainerID().(*baseIDs.MaintainerID)}).GeneratePrefixedStoreKeyBytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				MaintainerID: tt.fields.MaintainerID,
			}
			if got := key.GeneratePrefixedStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStoreKeyGenerateStoreKeyBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_IsPartial(t *testing.T) {
	type fields struct {
		MaintainerID *baseIDs.MaintainerID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{createTestData()}, false},
		{"-ve", fields{baseIDs.PrototypeMaintainerID().(*baseIDs.MaintainerID)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				MaintainerID: tt.fields.MaintainerID,
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}
