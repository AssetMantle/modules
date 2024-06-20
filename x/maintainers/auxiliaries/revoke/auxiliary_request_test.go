// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/ids"

	"github.com/AssetMantle/modules/helpers"
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
		{"+ve", args{testFromID, testFromID, testClassificationID}, auxiliaryRequest{testFromID, testFromID, testClassificationID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliaryRequest(tt.args.fromID, tt.args.toID, tt.args.maintainedClassificationID); !reflect.DeepEqual(got, tt.want) {
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
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{}, false},
		{"+ve", fields{testFromID, testFromID, testClassificationID}, false},
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
