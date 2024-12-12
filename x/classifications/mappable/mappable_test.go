// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseData "github.com/AssetMantle/schema/data/base"
	documentsSchema "github.com/AssetMantle/schema/documents"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/qualified"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
)

func createTestInput() (ids.ClassificationID, qualified.Immutables, qualified.Mutables, *Mappable) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testMappable := &Mappable{Classification: baseDocuments.NewDocument(classificationID, immutables, mutables).(*baseDocuments.Document)}
	return classificationID, immutables, mutables, testMappable
}

func TestNewMappable(t *testing.T) {
	classificationID, immutables, mutables, testMappable := createTestInput()
	type args struct {
		Asset documentsSchema.Document
	}
	tests := []struct {
		name string
		args args
		want helpers.Mappable
	}{
		{"+ve", args{baseDocuments.NewDocument(classificationID, immutables, mutables)}, testMappable},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMappable(baseDocuments.NewClassificationFromDocument(tt.args.Asset)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Mappable
	}{
		{"+ve", &Mappable{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
