// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"

	"github.com/AssetMantle/modules/helpers"
)

func createTestInput1() ids.ClassificationID {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)

	return classificationID
}

func TestGetClassificationIDFromResponse(t *testing.T) {
	classificationID := createTestInput1()
	type args struct {
		response helpers.AuxiliaryResponse
	}
	tests := []struct {
		name    string
		args    args
		want    ids.ClassificationID
		wantErr bool
	}{
		{"+ve", args{NewAuxiliaryResponse(classificationID)}, classificationID, false},
		{"+ve", args{NewAuxiliaryResponse(nil)}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetClassificationIDFromResponse(tt.args.response)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClassificationIDFromResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newAuxiliaryResponse(t *testing.T) {
	classificationID := createTestInput1()
	type args struct {
		classificationID ids.ClassificationID
		error            error
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryResponse
	}{
		{"+ve", args{classificationID, nil}, auxiliaryResponse{classificationID}},
		{"-ve", args{classificationID, errorConstants.EntityNotFound}, auxiliaryResponse{classificationID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliaryResponse(tt.args.classificationID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuxiliaryResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
