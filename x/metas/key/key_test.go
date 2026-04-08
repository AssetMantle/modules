// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/stretchr/testify/assert"
)

var (
	testDataID  = baseIDs.GenerateDataID(baseData.NewStringData("Data")).(*baseIDs.DataID)
	testDataID1 = baseIDs.PrototypeDataID().(*baseIDs.DataID)
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
		{"valid", args{testDataID}, &Key{testDataID}},
		{"valid", args{testDataID1}, &Key{testDataID1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewKey(tt.args.dataID)
			assert.Equal(t, tt.want, got, "NewKey()")
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Key
	}{
		{"valid", &Key{baseIDs.PrototypeDataID().(*baseIDs.DataID)}},
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
		{"valid", fields{testDataID}, args{NewKey(testDataID)}, true},
		{"valid", fields{testDataID}, args{NewKey(testDataID1)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				DataID: tt.fields.DataID,
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
		{"valid", fields{testDataID}, (&Key{testDataID}).GeneratePrefixedStoreKeyBytes()},
		{"nil inputs", fields{testDataID1}, (&Key{testDataID1}).GeneratePrefixedStoreKeyBytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				DataID: tt.fields.DataID,
			}
			got := key.GeneratePrefixedStoreKeyBytes()
			assert.Equal(t, tt.want, got, "GenerateStoreKeyGenerateStoreKeyBytes()")
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
		{"valid", fields{testDataID}, false},
		{"invalid", fields{testDataID1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				DataID: tt.fields.DataID,
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}
