// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
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
	immutables       = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables         = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID = baseIDs.NewClassificationID(immutables, mutables)
	testOrderID      = baseIDs.NewOrderID(classificationID, immutables).(*baseIDs.OrderID)
)

func TestNewKey(t *testing.T) {
	type args struct {
		orderID ids.OrderID
	}
	tests := []struct {
		name string
		args args
		want helpers.Key
	}{
		{"valid", args{testOrderID}, &Key{OrderID: testOrderID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewKey(tt.args.orderID)
			assert.Equal(t, tt.want, got, "NewKey()")
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Key
	}{
		{"valid", &Key{baseIDs.PrototypeOrderID().(*baseIDs.OrderID)}},
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
		OrderID *baseIDs.OrderID
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
		{"valid", fields{testOrderID}, args{&Key{testOrderID}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				OrderID: tt.fields.OrderID,
			}
			if got := key.Equals(tt.args.compareKey); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_GenerateStoreKeyBytes(t *testing.T) {
	type fields struct {
		OrderID *baseIDs.OrderID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"valid", fields{testOrderID}, (&Key{testOrderID}).GeneratePrefixedStoreKeyBytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				OrderID: tt.fields.OrderID,
			}
			got := key.GeneratePrefixedStoreKeyBytes()
			assert.Equal(t, tt.want, got, "GenerateStoreKeyGenerateStoreKeyBytes()")
		})
	}
}

func Test_key_IsPartial(t *testing.T) {
	type fields struct {
		OrderID *baseIDs.OrderID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"valid", fields{testOrderID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				OrderID: tt.fields.OrderID,
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}
