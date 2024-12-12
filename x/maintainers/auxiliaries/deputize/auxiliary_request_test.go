// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"

	"github.com/AssetMantle/modules/helpers"
)

func TestNewAuxiliaryRequest(t *testing.T) {
	type args struct {
		fromID                     ids.IdentityID
		toID                       ids.IdentityID
		maintainedClassificationID ids.ClassificationID
		maintainedProperties       lists.PropertyList
		canAddMaintainer           bool
		canRemoveMaintainer        bool
		canMutateMaintainer        bool
		permissionIDs              []ids.StringID
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryRequest
	}{
		{"+ve", args{testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true, []ids.StringID{}}, auxiliaryRequest{testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true, []ids.StringID{}}},
		{"+ve with nil", args{}, auxiliaryRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliaryRequest(tt.args.fromID, tt.args.toID, tt.args.maintainedClassificationID, tt.args.maintainedProperties, tt.args.canAddMaintainer, tt.args.canRemoveMaintainer, tt.args.canMutateMaintainer, tt.args.permissionIDs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuxiliaryRequest() = %v, want %v", got, tt.want)
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
		CanAddMaintainer           bool
		CanRemoveMaintainer        bool
		CanMutateMaintainer        bool
		PermissionIDs              []ids.StringID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true, []ids.StringID{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryRequest := auxiliaryRequest{
				tt.fields.FromID,
				tt.fields.ToID,
				tt.fields.MaintainedClassificationID,
				tt.fields.MaintainedProperties,
				tt.fields.CanAddMaintainer,
				tt.fields.CanRemoveMaintainer,
				tt.fields.CanMutateMaintainer,
				tt.fields.PermissionIDs,
			}
			if err := auxiliaryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
