// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authorize

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"

	"github.com/AssetMantle/modules/helpers"
)

var (
	immutableProperty = baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))
	mutableProperty   = baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))
	immutables        = baseQualified.NewImmutables(baseLists.NewPropertyList(immutableProperty))
	mutables          = baseQualified.NewMutables(baseLists.NewPropertyList(mutableProperty))
	classificationID  = baseIDs.NewClassificationID(immutables, mutables)
	identityID        = baseIDs.NewIdentityID(classificationID, immutables)
)

func TestNewAuxiliaryRequest(t *testing.T) {
	type args struct {
		maintainedClassificationID ids.ClassificationID
		maintainedIdentityID       ids.IdentityID
		permissionIDs              []ids.StringID
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryRequest
	}{
		{"+ve", args{classificationID, identityID, []ids.StringID{}}, auxiliaryRequest{classificationID, identityID, []ids.StringID{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliaryRequest(tt.args.maintainedClassificationID, tt.args.maintainedIdentityID); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", args{NewAuxiliaryRequest(classificationID, identityID)}, auxiliaryRequest{classificationID, identityID, []ids.StringID{}}},
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
		MaintainedClassificationID ids.ClassificationID
		MaintainerIdentityID       ids.IdentityID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{classificationID, identityID}, false},
		{"+ve", fields{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryRequest := auxiliaryRequest{
				MaintainedClassificationID: tt.fields.MaintainedClassificationID,
				MaintainerIdentityID:       tt.fields.MaintainerIdentityID,
			}
			if err := auxiliaryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
