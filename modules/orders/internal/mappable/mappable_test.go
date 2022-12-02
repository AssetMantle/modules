// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var (
	immutables       = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables         = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIds.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID = baseIds.NewClassificationID(immutables, mutables)
	testOrder        = baseDocuments.NewOrder(classificationID, immutables, mutables)
	//testIdentity     = baseDocuments.NewIdentity(classificationID, immutables, mutables)
)

func TestNewMappable(t *testing.T) {
	type args struct {
		order documents.Order
	}
	tests := []struct {
		name string
		args args
		want helpers.Mappable
	}{
		{"+ve", args{testOrder}, mappable{baseDocuments.NewOrder(classificationID, immutables, mutables)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMappable(tt.args.order); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", mappable{}},
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
		Order documents.Order
	}
	tests := []struct {
		name      string
		fields    fields
		want      helpers.Key
		wantPanic bool
	}{
		{"+ve", fields{testOrder}, key.NewKey(baseIds.NewOrderID(classificationID, immutables)), false},
		{"panic case nil", fields{nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mappable := mappable{
				Order: tt.fields.Order,
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					mappable.GetKey()
				})
			} else if got := mappable.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mappable_RegisterCodec(t *testing.T) {
	type fields struct {
		Order documents.Order
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testOrder}, args{codec.New()}},
		{"+ve nil", fields{nil}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := mappable{
				Order: tt.fields.Order,
			}
			ma.RegisterCodec(tt.args.codec)
		})
	}
}
