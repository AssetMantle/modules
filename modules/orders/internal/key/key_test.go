// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/orders/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
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
		{"+ve", args{testOrderID}, &Key{OrderID: testOrderID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKey(tt.args.orderID); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", args{}, &Key{}, true},
		{"+ve", args{NewKey(testOrderID)}, &Key{testOrderID}, false},
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
		{"+ve", fields{testOrderID}, args{&Key{testOrderID}}, true},
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
		{"+ve", fields{testOrderID}, module.StoreKeyPrefix.GenerateStoreKey((&Key{testOrderID}).GenerateStoreKeyBytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				OrderID: tt.fields.OrderID,
			}
			if got := key.GenerateStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStoreKeyGenerateStoreKeyBytes() = %v, want %v", got, tt.want)
			}
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
		{"+ve", fields{testOrderID}, false},
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

func Test_key_RegisterCodec(t *testing.T) {
	type fields struct {
		OrderID *baseIDs.OrderID
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testOrderID}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ke := &Key{
				OrderID: tt.fields.OrderID,
			}
			ke.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
