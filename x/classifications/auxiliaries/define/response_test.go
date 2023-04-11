// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/x/data/base"
	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"

	"github.com/AssetMantle/modules/helpers"
)

func createTestInput() ids.ClassificationID {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)

	return classificationID
}

func TestGetClassificationIDFromResponse(t *testing.T) {
	classificationID := createTestInput()
	type args struct {
		response helpers.AuxiliaryResponse
	}
	tests := []struct {
		name    string
		args    args
		want    ids.ClassificationID
		wantErr bool
	}{
		{"+ve", args{newAuxiliaryResponse(classificationID, nil)}, classificationID, false},
		{"+ve", args{newAuxiliaryResponse(nil, nil)}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetClassificationIDFromResponse(tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClassificationIDFromResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClassificationIDFromResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryResponse_GetError(t *testing.T) {
	classificationID := createTestInput()
	type fields struct {
		Success          bool
		Error            error
		ClassificationID ids.ClassificationID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{true, nil, classificationID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryResponse := auxiliaryResponse{
				Success:          tt.fields.Success,
				Error:            tt.fields.Error,
				ClassificationID: tt.fields.ClassificationID,
			}
			if err := auxiliaryResponse.GetError(); (err != nil) != tt.wantErr {
				t.Errorf("GetError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_auxiliaryResponse_IsSuccessful(t *testing.T) {
	classificationID := createTestInput()
	type fields struct {
		Success          bool
		Error            error
		ClassificationID ids.ClassificationID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{true, nil, classificationID}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryResponse := auxiliaryResponse{
				Success:          tt.fields.Success,
				Error:            tt.fields.Error,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := auxiliaryResponse.IsSuccessful(); got != tt.want {
				t.Errorf("IsSuccessful() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newAuxiliaryResponse(t *testing.T) {
	classificationID := createTestInput()
	type args struct {
		classificationID ids.ClassificationID
		error            error
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryResponse
	}{
		{"+ve", args{classificationID, nil}, auxiliaryResponse{true, nil, classificationID}},
		{"-ve", args{classificationID, errorConstants.EntityNotFound}, auxiliaryResponse{false, errorConstants.EntityNotFound, classificationID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAuxiliaryResponse(tt.args.classificationID, tt.args.error); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAuxiliaryResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
