// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/documents"
	baseDocuments "github.com/AssetMantle/schema/x/documents/base"
	"github.com/AssetMantle/schema/x/helpers"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/x/orders/internal/key"
)

var (
	immutables       = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables         = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID = baseIDs.NewClassificationID(immutables, mutables)
	testOrder        = baseDocuments.NewOrder(classificationID, immutables, mutables)
	// testIdentity     = baseDocuments.NewIdentity(classificationID, immutables, mutables)
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
		{"+ve", args{testOrder}, &Mappable{baseDocuments.NewOrder(classificationID, immutables, mutables).Get().(*baseDocuments.Document)}},
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
		{"+ve", &Mappable{}},
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
		{"+ve", fields{testOrder}, key.NewKey(baseIDs.NewOrderID(classificationID, immutables)), false},
		{"panic case nil", fields{nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mappable := &Mappable{
				Order: tt.fields.Order.Get().(*baseDocuments.Document),
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
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testOrder}, args{codec.NewLegacyAmino()}},
		{"+ve nil", fields{nil}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := &Mappable{
				Order: tt.fields.Order.Get().(*baseDocuments.Document),
			}
			ma.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
