// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"testing"

	"github.com/AssetMantle/schema/ids"

	"github.com/AssetMantle/modules/helpers"
	"github.com/stretchr/testify/assert"
)

func TestNewAuxiliaryRequest(t *testing.T) {
	type args struct {
		fromID                     ids.IdentityID
		toID                       ids.IdentityID
		maintainedClassificationID ids.ClassificationID
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryRequest
	}{
		{"valid", args{testFromID, testFromID, testClassificationID}, auxiliaryRequest{testFromID, testFromID, testClassificationID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAuxiliaryRequest(tt.args.fromID, tt.args.toID, tt.args.maintainedClassificationID)
			assert.Equal(t, tt.want, got, "NewAuxiliaryRequest()")
		})
	}
}

func Test_auxiliaryRequest_Validate(t *testing.T) {
	type fields struct {
		FromID                     ids.IdentityID
		ToID                       ids.IdentityID
		MaintainedClassificationID ids.ClassificationID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"nil inputs error", fields{}, true},
		{"valid", fields{testFromID, testFromID, testClassificationID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryRequest := auxiliaryRequest{
				FromID:                     tt.fields.FromID,
				ToID:                       tt.fields.ToID,
				MaintainedClassificationID: tt.fields.MaintainedClassificationID,
			}
			if err := auxiliaryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
