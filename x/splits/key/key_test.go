// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/stretchr/testify/assert"
)

var (
	immutables          = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables            = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID    = baseIDs.NewClassificationID(immutables, mutables)
	testOwnerIdentityID = baseIDs.NewIdentityID(classificationID, immutables)
	testAssetID         = baseDocuments.NewCoinAsset("ownerID").GetCoinAssetID().(*baseIDs.AssetID)
	splitID             = baseIDs.NewSplitID(testAssetID, testOwnerIdentityID).(*baseIDs.SplitID)
)

func TestNewKey(t *testing.T) {
	type args struct {
		splitID ids.SplitID
	}
	tests := []struct {
		name string
		args args
		want helpers.Key
	}{
		{"valid", args{splitID}, &Key{splitID}},
		{"nil inputs", args{baseIDs.PrototypeSplitID()}, &Key{baseIDs.PrototypeSplitID().(*baseIDs.SplitID)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewKey(tt.args.splitID)
			assert.Equal(t, tt.want, got, "NewKey()")
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Key
	}{
		{"valid", &Key{baseIDs.PrototypeSplitID().(*baseIDs.SplitID)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prototype()
			assert.Equal(t, tt.want, got, "Prototype()")
		})
	}
}

func Test_key_Equals(t *testing.T) {
	type fields struct {
		SplitID ids.SplitID
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
		{"valid", fields{splitID}, args{NewKey(splitID)}, true},
		{"valid", fields{splitID}, args{NewKey(baseIDs.PrototypeSplitID())}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				SplitID: tt.fields.SplitID.(*baseIDs.SplitID),
			}
			if got := key.Equals(tt.args.compareKey); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_GenerateStoreKeyBytes(t *testing.T) {
	type fields struct {
		SplitID ids.SplitID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"valid", fields{splitID}, (&Key{splitID}).GeneratePrefixedStoreKeyBytes()},
		{"valid", fields{baseIDs.PrototypeSplitID()}, (&Key{baseIDs.PrototypeSplitID().(*baseIDs.SplitID)}).GeneratePrefixedStoreKeyBytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				SplitID: tt.fields.SplitID.(*baseIDs.SplitID),
			}
			got := key.GeneratePrefixedStoreKeyBytes()
			assert.Equal(t, tt.want, got, "GeneratePrefixedStoreKeyBytes()")
		})
	}
}

func Test_key_IsPartial(t *testing.T) {
	type fields struct {
		SplitID ids.SplitID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"valid", fields{splitID}, false},
		{"valid", fields{baseIDs.PrototypeSplitID()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				SplitID: tt.fields.SplitID.(*baseIDs.SplitID),
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}
