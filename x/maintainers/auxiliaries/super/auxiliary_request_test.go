// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/qualified"

	"github.com/AssetMantle/modules/helpers"
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
		{"+ve", args{testClassificationID, testFromID, mutables}, auxiliaryRequest{testClassificationID, testFromID, mutables, []ids.StringID{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliaryRequest(tt.args.maintainedClassificationID, tt.args.toIdentityID, tt.args.maintainedMutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuxiliaryRequest() = %v, want %v", got, tt.want)
			}
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
		{"+ve with nil", fields{}, false},
		{"+ve", fields{testClassificationID, testFromID, mutables}, false},
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
