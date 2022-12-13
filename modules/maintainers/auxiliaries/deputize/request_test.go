// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
)

func TestNewAuxiliaryRequest(t *testing.T) {
	type args struct {
		fromID                     ids.IdentityID
		toID                       ids.IdentityID
		maintainedClassificationID ids.ClassificationID
		maintainedProperties       lists.PropertyList
		canMintAsset               bool
		canBurnAsset               bool
		canRenumerateAsset         bool
		canAddMaintainer           bool
		canRemoveMaintainer        bool
		canMutateMaintainer        bool
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryRequest
	}{
		{"+ve", args{testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true, true, true, true}, auxiliaryRequest{testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true, true, true, true}},
		{"+ve with nil", args{}, auxiliaryRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliaryRequest(tt.args.fromID, tt.args.toID, tt.args.maintainedClassificationID, tt.args.maintainedProperties, tt.args.canMintAsset, tt.args.canBurnAsset, tt.args.canRenumerateAsset, tt.args.canAddMaintainer, tt.args.canRemoveMaintainer, tt.args.canMutateMaintainer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuxiliaryRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryRequestFromInterface(t *testing.T) {
	type args struct {
		request helpers.AuxiliaryRequest
	}
	tests := []struct {
		name string
		args args
		want auxiliaryRequest
	}{
		{"+ve", args{NewAuxiliaryRequest(testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true, true, true, true)}, auxiliaryRequest{testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true, true, true, true}},
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
	type fields struct {
		FromID                     ids.IdentityID
		ToID                       ids.IdentityID
		MaintainedClassificationID ids.ClassificationID
		MaintainedProperties       lists.PropertyList
		CanMintAsset               bool
		CanBurnAsset               bool
		CanRenumerateAsset         bool
		CanAddMaintainer           bool
		CanRemoveMaintainer        bool
		CanMutateMaintainer        bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true, true, true, true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryRequest := auxiliaryRequest{
				FromID:                     tt.fields.FromID,
				ToID:                       tt.fields.ToID,
				MaintainedClassificationID: tt.fields.MaintainedClassificationID,
				MaintainedProperties:       tt.fields.MaintainedProperties,
				CanMintAsset:               tt.fields.CanMintAsset,
				CanBurnAsset:               tt.fields.CanBurnAsset,
				CanRenumerateAsset:         tt.fields.CanRenumerateAsset,
				CanAddMaintainer:           tt.fields.CanAddMaintainer,
				CanRemoveMaintainer:        tt.fields.CanRemoveMaintainer,
				CanMutateMaintainer:        tt.fields.CanMutateMaintainer,
			}
			if err := auxiliaryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
