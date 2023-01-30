// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/splits/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

var (
	immutables          = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables            = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID    = baseIds.NewClassificationID(immutables, mutables)
	testOwnerIdentityID = baseIds.NewIdentityID(classificationID, immutables)
	testOwnableID       = baseIds.NewCoinID(baseIds.NewStringID("ownerid"))
	splitID             = baseIds.NewSplitID(testOwnerIdentityID, testOwnableID).(*baseIds.SplitID)
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
		{"+ve", args{splitID}, &Key{splitID}},
		{"+ve with nil", args{baseIds.PrototypeSplitID()}, &Key{baseIds.PrototypeSplitID().(*baseIds.SplitID)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKey(tt.args.splitID); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", args{NewKey(splitID)}, &Key{splitID}, false},
		{"+ve", args{NewKey(baseIds.PrototypeSplitID())}, &Key{baseIds.PrototypeSplitID().(*baseIds.SplitID)}, false},
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
		{"+ve", fields{splitID}, args{NewKey(splitID)}, true},
		{"+ve", fields{splitID}, args{NewKey(baseIds.PrototypeSplitID())}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				SplitID: tt.fields.SplitID.(*baseIds.SplitID),
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
		{"+ve", fields{splitID}, module.StoreKeyPrefix.GenerateStoreKey((&Key{splitID}).GenerateStoreKeyBytes())},
		{"+ve", fields{baseIds.PrototypeSplitID()}, module.StoreKeyPrefix.GenerateStoreKey((&Key{baseIds.PrototypeSplitID().(*baseIds.SplitID)}).GenerateStoreKeyBytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				SplitID: tt.fields.SplitID.(*baseIds.SplitID),
			}
			if got := key.GenerateStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStoreKeyBytes() = %v, want %v", got, tt.want)
			}
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
		{"+ve", fields{splitID}, false},
		{"+ve", fields{baseIds.PrototypeSplitID()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				SplitID: tt.fields.SplitID.(*baseIds.SplitID),
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_RegisterCodec(t *testing.T) {
	type fields struct {
		SplitID ids.SplitID
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{splitID}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ke := &Key{
				SplitID: tt.fields.SplitID.(*baseIds.SplitID),
			}
			ke.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
