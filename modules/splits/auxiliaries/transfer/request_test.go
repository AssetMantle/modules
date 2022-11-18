// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

func createTestInput1() (ids.IdentityID, ids.OwnableID, types.Dec) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIds.NewClassificationID(immutables, mutables)
	testOwnerID := baseIds.NewIdentityID(classificationID, immutables)
	testOwnableID := baseIds.NewOwnableID(baseIds.NewStringID("OwnerID"))
	testValue := types.NewDec(1)
	return testOwnerID, testOwnableID, testValue
}

func TestNewAuxiliaryRequest(t *testing.T) {
	testOwnerID, testOwnableID, testValue := createTestInput1()
	type args struct {
		fromID    ids.IdentityID
		toID      ids.IdentityID
		ownableID ids.OwnableID
		value     types.Dec
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryRequest
	}{
		// TODO: Add test cases.
		{"+ve", args{testOwnerID, testOwnerID, testOwnableID, testValue}, NewAuxiliaryRequest(testOwnerID, testOwnerID, testOwnableID, testValue)},
		{"+ve with nil", args{testOwnerID, baseIds.PrototypeIdentityID(), baseIds.PrototypeOwnableID(), testValue}, NewAuxiliaryRequest(testOwnerID, baseIds.PrototypeIdentityID(), baseIds.PrototypeOwnableID(), testValue)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliaryRequest(tt.args.fromID, tt.args.toID, tt.args.ownableID, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuxiliaryRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryRequestFromInterface(t *testing.T) {
	testOwnerID, testOwnableID, testValue := createTestInput1()
	type args struct {
		request helpers.AuxiliaryRequest
	}
	tests := []struct {
		name string
		args args
		want auxiliaryRequest
	}{
		// TODO: Add test cases.
		{"+ve", args{NewAuxiliaryRequest(testOwnerID, testOwnerID, testOwnableID, testValue)}, auxiliaryRequest{testOwnerID, testOwnerID, testOwnableID, testValue}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := auxiliaryRequestFromInterface(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("auxiliaryRequestFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryRequest_Validate(t *testing.T) {
	testOwnerID, testOwnableID, testValue := createTestInput1()
	type fields struct {
		FromID    ids.IdentityID
		ToID      ids.IdentityID
		OwnableID ids.OwnableID
		Value     types.Dec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{testOwnerID, testOwnerID, testOwnableID, testValue}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryRequest := auxiliaryRequest{
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			if err := auxiliaryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
