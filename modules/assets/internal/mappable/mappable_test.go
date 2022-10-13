// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func createTestInput() (ids.ClassificationID, qualified.Immutables, qualified.Mutables, qualified.Document) {
	immutables := baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testDocument := baseQualified.NewDocument(classificationID, immutables, mutables)
	return classificationID, immutables, mutables, testDocument
}

func TestNewAsset(t *testing.T) {
	classificationID, immutables, mutables, _ := createTestInput()
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
		mutables         qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want mappables.Asset
	}{
		// TODO: Add test cases.
		{"+ve", args{classificationID: classificationID, immutables: immutables, mutables: mutables}, mappable{Document: baseQualified.NewDocument(classificationID, immutables, mutables)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAsset(tt.args.classificationID, tt.args.immutables, tt.args.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAsset() = %v, want %v", got, tt.want)
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

func Test_asset_GetBurn(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithBurn := baseQualified.NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(constants.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))))))

	type fields struct {
		Document qualified.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		// TODO: Add test cases.
		{"+ve", fields{Document: testDocumentWithBurn}, baseProperties.NewMesaProperty(constants.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1)))},
		{"+ve", fields{Document: testDocument}, constants.BurnHeightProperty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := mappable{
				Document: tt.fields.Document,
			}
			if got := asset.GetBurn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBurn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asset_GetKey(t *testing.T) {
	_, _, _, testDocument := createTestInput()
	type fields struct {
		Document qualified.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.Key
	}{
		// TODO: Add test cases.
		{"+ve", fields{testDocument}, key.NewKey(baseIDs.NewAssetID(mappable{testDocument}.GetClassificationID(), mappable{testDocument}.GetImmutables()))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := mappable{
				Document: tt.fields.Document,
			}
			if got := asset.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asset_GetLock(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithLock := baseQualified.NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(constants.LockProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))))))

	type fields struct {
		Document qualified.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		// TODO: Add test cases.
		{"+ve with default lock", fields{testDocument}, constants.LockProperty},
		{"+ve with mutated", fields{testDocumentWithLock}, baseProperties.NewMesaProperty(constants.LockProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := mappable{
				Document: tt.fields.Document,
			}
			if got := asset.GetLock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asset_GetSupply(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithSupply := baseQualified.NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(constants.SupplyProperty.GetKey(), baseData.NewDecData(types.NewDec(1))))))
	type fields struct {
		Document qualified.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		// TODO: Add test cases.
		{"+ve", fields{testDocument}, constants.SupplyProperty},
		{"+ve", fields{testDocumentWithSupply}, baseProperties.NewMesaProperty(constants.SupplyProperty.GetKey(), baseData.NewDecData(types.NewDec(1)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := mappable{
				Document: tt.fields.Document,
			}
			if got := asset.GetSupply(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSupply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asset_RegisterCodec(t *testing.T) {
	_, _, _, testDocument := createTestInput()
	type fields struct {
		Document qualified.Document
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
		{"+ve", fields{testDocument}, args{codec: codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := mappable{
				Document: tt.fields.Document,
			}
			as.RegisterCodec(tt.args.codec)
		})
	}
}
