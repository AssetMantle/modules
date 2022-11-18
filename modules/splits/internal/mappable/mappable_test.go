// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/modules/splits/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

var (
	immutables          = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables            = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID    = baseIds.NewClassificationID(immutables, mutables)
	testOwnerIdentityID = baseIds.NewIdentityID(classificationID, immutables)
	testOwnableID       = baseIds.NewOwnableID(baseIds.NewStringID("ownerid"))
	splitID             = baseIds.NewSplitID(testOwnerIdentityID, testOwnableID)
	testRate            = sdkTypes.NewDec(1)
	split               = base.NewSplit(testOwnerIdentityID, testOwnableID, testRate)
)

func TestNewMappable(t *testing.T) {
	type args struct {
		split types.Split
	}
	tests := []struct {
		name string
		args args
		want helpers.Mappable
	}{
		// TODO: Add test cases.
		{"+ve", args{split}, mappable{split}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMappable(tt.args.split); !reflect.DeepEqual(got, tt.want) {
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
		// TODO: Add test cases.
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
		Split types.Split
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.Key
	}{
		// TODO: Add test cases.
		{"+ve", fields{split}, key.NewKey(splitID)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mappable := mappable{
				Split: tt.fields.Split,
			}
			if got := mappable.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mappable_RegisterCodec(t *testing.T) {
	type fields struct {
		Split types.Split
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"+ve", fields{split}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := mappable{
				Split: tt.fields.Split,
			}
			ma.RegisterCodec(tt.args.codec)
		})
	}
}
