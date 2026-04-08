// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"cosmossdk.io/math"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/AssetMantle/modules/helpers"
	"github.com/stretchr/testify/assert"
)

func createTestInput1() (ids.IdentityID, ids.AssetID, math.Int) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testOwnerID := baseIDs.NewIdentityID(classificationID, immutables)
	testAssetID := baseDocuments.NewCoinAsset("OwnerID").GetCoinAssetID()
	testValue := math.OneInt()
	return testOwnerID, testAssetID, testValue
}

func TestNewAuxiliaryRequest(t *testing.T) {
	testOwnerID, testAssetID, testValue := createTestInput1()
	type args struct {
		ownerID ids.IdentityID
		assetID ids.AssetID
		value   math.Int
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryRequest
	}{
		{"valid", args{testOwnerID, testAssetID, testValue}, NewAuxiliaryRequest(testOwnerID, testAssetID, testValue)},
		{"nil inputs", args{baseIDs.PrototypeIdentityID(), baseIDs.PrototypeAssetID(), testValue}, NewAuxiliaryRequest(baseIDs.PrototypeIdentityID(), baseIDs.PrototypeAssetID(), testValue)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAuxiliaryRequest(tt.args.ownerID, tt.args.assetID, tt.args.value)
			assert.Equal(t, tt.want, got, "NewAuxiliaryRequest()")
		})
	}
}

func Test_auxiliaryRequest_Validate(t *testing.T) {
	testOwnerID, testAssetID, testValue := createTestInput1()
	type fields struct {
		OwnerID ids.IdentityID
		AssetID ids.AssetID
		Value   math.Int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{testOwnerID, testAssetID, testValue}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryRequest := auxiliaryRequest{
				OwnerID: tt.fields.OwnerID,
				AssetID: tt.fields.AssetID,
				Supply:  tt.fields.Value,
			}
			if err := auxiliaryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
