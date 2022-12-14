// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	constanstsProperties "github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	sdkType "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

var (
	testOrder = NewOrder(testClassificationID, testImmutables, testMutables)
)

func TestNewOrder(t *testing.T) {
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
		mutables         qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want documents.Order
	}{
		{"+ve with nil immutables", args{testClassificationID, nil, testMutables}, order{NewDocument(testClassificationID, nil, testMutables)}},
		{"+ve with nil mutables", args{testClassificationID, testImmutables, nil}, order{NewDocument(testClassificationID, testImmutables, nil)}},
		{"+ve with nil immutables and mutables", args{testClassificationID, nil, nil}, order{NewDocument(testClassificationID, nil, nil)}},
		{"+ve with nil classificationID", args{nil, testImmutables, testMutables}, order{NewDocument(nil, testImmutables, testMutables)}},
		{"+ve", args{testClassificationID, testImmutables, testMutables}, order{NewDocument(testClassificationID, testImmutables, testMutables)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrder(tt.args.classificationID, tt.args.immutables, tt.args.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_order_GetCreationHeight(t *testing.T) {
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   types.Height
	}{
		{"+ve", fields{testOrder}, baseTypes.NewHeight(int64(1))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := order{
				Document: tt.fields.Document,
			}
			if got := order.GetCreationHeight(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCreationHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_order_GetExchangeRate(t *testing.T) {
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   sdkType.Dec
	}{
		{"+ve", fields{testOrder}, sdkType.NewDec(int64(10))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := order{
				Document: tt.fields.Document,
			}
			if got := order.GetExchangeRate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExchangeRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_order_GetExpiryHeight(t *testing.T) {
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   types.Height
	}{
		{"+ve", fields{testOrder}, baseTypes.NewHeight(int64(100))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := order{
				Document: tt.fields.Document,
			}
			if got := order.GetExpiryHeight(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpiryHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_order_GetMakerID(t *testing.T) {
	testOrder.Mutate(baseProperties.NewMetaProperty(constanstsProperties.MakerIDProperty.GetKey(), baseData.NewIDData(testIdentityID)))
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.IdentityID
	}{
		{"+ve", fields{testOrder}, testIdentityID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := order{
				Document: tt.fields.Document,
			}
			if got := order.GetMakerID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMakerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_order_GetMakerOwnableID(t *testing.T) {
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.OwnableID
	}{
		{"+ve", fields{testOrder}, base.NewOwnableID(base.NewStringID("MakerOwnableID"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := order{
				Document: tt.fields.Document,
			}
			if got := order.GetMakerOwnableID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMakerOwnableID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_order_GetMakerOwnableSplit(t *testing.T) {
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   sdkType.Dec
	}{
		{"+ve", fields{testOrder}, sdkType.NewDec(int64(10))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := order{
				Document: tt.fields.Document,
			}
			if got := order.GetMakerOwnableSplit(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMakerOwnableSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_order_GetTakerID(t *testing.T) {
	testOrder.Mutate(baseProperties.NewMetaProperty(constanstsProperties.TakerIDProperty.GetKey(), baseData.NewIDData(testIdentityID)))
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.IdentityID
	}{
		{"+ve with IdentityID", fields{testOrder}, testIdentityID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := order{
				Document: tt.fields.Document,
			}
			if got := order.GetTakerID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTakerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_order_GetTakerOwnableID(t *testing.T) {
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.OwnableID
	}{
		{"+ve", fields{testOrder}, base.NewOwnableID(base.NewStringID("TakerOwnableID"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := order{
				Document: tt.fields.Document,
			}
			if got := order.GetTakerOwnableID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTakerOwnableID() = %v, want %v", got, tt.want)
			}
		})
	}
}
