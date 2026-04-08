// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	"testing"

	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/qualified"

	"github.com/AssetMantle/modules/helpers"
	"github.com/stretchr/testify/assert"
)

func TestNewAuxiliaryRequest(t *testing.T) {
	type args struct {
		maintainedClassificationID ids.ClassificationID
		toIdentityID               ids.IdentityID
		maintainedMutables         qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryRequest
	}{
		{"valid", args{testClassificationID, testFromID, mutables}, auxiliaryRequest{testClassificationID, testFromID, mutables, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAuxiliaryRequest(tt.args.maintainedClassificationID, tt.args.toIdentityID, tt.args.maintainedMutables)
			assert.Equal(t, tt.want, got, "NewAuxiliaryRequest()")
		})
	}
}

func Test_auxiliaryRequest_Validate(t *testing.T) {
	type fields struct {
		MaintainedClassificationID ids.ClassificationID
		ToIdentityID               ids.IdentityID
		MaintainedMutables         qualified.Mutables
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid", fields{testClassificationID, testFromID, mutables}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryRequest := auxiliaryRequest{
				MaintainedClassificationID: tt.fields.MaintainedClassificationID,
				ToIdentityID:               tt.fields.ToIdentityID,
				MaintainedMutables:         tt.fields.MaintainedMutables,
			}
			if err := auxiliaryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
