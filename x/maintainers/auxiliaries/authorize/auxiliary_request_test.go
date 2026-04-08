// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authorize

import (
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/stretchr/testify/assert"
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
		{"valid", args{classificationID, identityID, []ids.StringID{}}, auxiliaryRequest{classificationID, identityID, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAuxiliaryRequest(tt.args.maintainedClassificationID, tt.args.maintainedIdentityID)
			assert.Equal(t, tt.want, got, "NewAuxiliaryRequest()")
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
		{"valid", fields{classificationID, identityID}, false},
		{"nil inputs error", fields{}, true},
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
