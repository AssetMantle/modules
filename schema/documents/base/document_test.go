// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func createTestInput() (ids.ClassificationID, qualified.Immutables, qualified.Mutables, documents.Document) {
	testImmutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	testMutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList()))))
	classificationID := baseIDs.NewClassificationID(testImmutables, testMutables)
	testDocument := NewDocument(classificationID, testImmutables, testMutables)
	return classificationID, testImmutables, testMutables, testDocument
}

func TestNewDocument(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
		mutables         qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want documents.Document
	}{
		// TODO: Add test cases.
		{"+ve", args{classificationID: classificationID, immutables: testImmutables, mutables: testMutables}, document{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}},
		{"+ve with nil classificationID", args{classificationID: nil, immutables: testImmutables, mutables: testMutables}, document{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}},
		{"+ve with nil immutables", args{classificationID: classificationID, immutables: nil, mutables: testMutables}, document{ClassificationID: classificationID, Immutables: nil, Mutables: testMutables}},
		{"+ve with nil mutables", args{classificationID: classificationID, immutables: testImmutables, mutables: nil}, document{ClassificationID: classificationID, Immutables: testImmutables, Mutables: nil}},
		{"+ve with all nil", args{}, document{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDocument(tt.args.classificationID, tt.args.immutables, tt.args.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_document_GetClassificationID(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	type fields struct {
		ClassificationID ids.ClassificationID
		Immutables       qualified.Immutables
		Mutables         qualified.Mutables
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ClassificationID
	}{
		// TODO: Add test cases.
		{"+ve", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}, classificationID},
		{"+ve with nil classificationID", fields{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}, nil},
		{"+ve with nil immutables", fields{ClassificationID: classificationID, Immutables: nil, Mutables: testMutables}, classificationID},
		{"+ve with nil mutables", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: nil}, classificationID},
		{"+ve with all nil", fields{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := document{
				ClassificationID: tt.fields.ClassificationID,
				Immutables:       tt.fields.Immutables,
				Mutables:         tt.fields.Mutables,
			}
			if got := document.GetClassificationID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClassificationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_document_GetImmutables(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	type fields struct {
		ClassificationID ids.ClassificationID
		Immutables       qualified.Immutables
		Mutables         qualified.Mutables
	}
	tests := []struct {
		name   string
		fields fields
		want   qualified.Immutables
	}{
		// TODO: Add test cases.
		{"+ve", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}, testImmutables},
		{"+ve with nil classificationID", fields{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}, testImmutables},
		{"+ve with nil immutables", fields{ClassificationID: classificationID, Immutables: nil, Mutables: testMutables}, nil},
		{"+ve with nil mutables", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: nil}, testImmutables},
		{"+ve with all nil", fields{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := document{
				ClassificationID: tt.fields.ClassificationID,
				Immutables:       tt.fields.Immutables,
				Mutables:         tt.fields.Mutables,
			}
			if got := document.GetImmutables(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImmutables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_document_GetMutables(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	type fields struct {
		ClassificationID ids.ClassificationID
		Immutables       qualified.Immutables
		Mutables         qualified.Mutables
	}
	tests := []struct {
		name   string
		fields fields
		want   qualified.Mutables
	}{
		// TODO: Add test cases.
		{"+ve", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}, testMutables},
		{"+ve with nil classificationID", fields{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}, testMutables},
		{"+ve with nil immutables", fields{ClassificationID: classificationID, Immutables: nil, Mutables: testMutables}, testMutables},
		{"+ve with nil mutables", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: nil}, nil},
		{"+ve with all nil", fields{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := document{
				ClassificationID: tt.fields.ClassificationID,
				Immutables:       tt.fields.Immutables,
				Mutables:         tt.fields.Mutables,
			}
			if got := document.GetMutables(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMutables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_document_GetProperty(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	testImmutablePropertyID := baseIDs.NewPropertyID(baseIDs.NewStringID("ID1"), baseIDs.NewStringID("ImmutableData"))
	testMutablePropertyID := baseIDs.NewPropertyID(baseIDs.NewStringID("ID2"), baseIDs.NewStringID("MutableData"))

	type fields struct {
		ClassificationID ids.ClassificationID
		Immutables       qualified.Immutables
		Mutables         qualified.Mutables
	}
	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      properties.Property
		wantPanic bool
	}{
		// TODO: Update unit test after issue # fix
		{"+ve for Immutable Property", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}, args{testImmutablePropertyID}, testImmutables.GetImmutablePropertyList().GetProperty(testImmutablePropertyID), false},
		{"+ve with nil classificationID for Immutable Property", fields{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}, args{testImmutablePropertyID}, testImmutables.GetImmutablePropertyList().GetProperty(testImmutablePropertyID), false},
		{"+ve with nil immutables for Immutable Property", fields{ClassificationID: classificationID, Immutables: baseQualified.NewImmutables(baseLists.NewPropertyList()), Mutables: testMutables}, args{testImmutablePropertyID}, nil, false}, // TODO: panics for empty immutable struct
		{"+ve with nil mutables for Immutable Property", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: baseQualified.NewMutables(baseLists.NewPropertyList())}, args{testImmutablePropertyID}, testImmutables.GetImmutablePropertyList().GetProperty(testImmutablePropertyID), false},
		{"+ve with all nil", fields{nil, baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList())}, args{testImmutablePropertyID}, nil, false}, // TODO: panics for empty immutable struct
		{"+ve for Mutable Property", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}, args{testMutablePropertyID}, testMutables.GetMutablePropertyList().GetProperty(testMutablePropertyID), false},
		{"+ve with nil classificationID for Mutable Property", fields{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}, args{testMutablePropertyID}, testMutables.GetMutablePropertyList().GetProperty(testMutablePropertyID), false},
		{"+ve with nil immutables for Mutable Property", fields{ClassificationID: classificationID, Immutables: baseQualified.NewImmutables(baseLists.NewPropertyList()), Mutables: testMutables}, args{testMutablePropertyID}, testMutables.GetMutablePropertyList().GetProperty(testMutablePropertyID), false}, // TODO: panics for empty immutable struct
		{"+ve with nil mutables for Mutable Property", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: baseQualified.NewMutables(baseLists.NewPropertyList())}, args{testMutablePropertyID}, nil, false},
		//Panics for empty structs
		//GetProperty() searches by traversing through Immutables and Mutables of a document, hence setting them as nil should cause a fatal panic
		{"panic case nil immutables", fields{classificationID, nil, testMutables}, args{testImmutablePropertyID}, nil, true},
		{"panic case nil mutables", fields{classificationID, testImmutables, nil}, args{testImmutablePropertyID}, nil, true},
		{"panic case nil document", fields{nil, nil, nil}, args{testImmutablePropertyID}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := document{
				ClassificationID: tt.fields.ClassificationID,
				Immutables:       tt.fields.Immutables,
				Mutables:         tt.fields.Mutables,
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					document.GetProperty(tt.args.propertyID)
				})
			} else if got := document.GetProperty(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_document_Mutate(t *testing.T) {
	classificationID, testImmutables, testMutables, _ := createTestInput()
	testMutateProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("MutableData1"))
	type fields struct {
		ClassificationID ids.ClassificationID
		Immutables       qualified.Immutables
		Mutables         qualified.Mutables
	}
	type args struct {
		propertyList []properties.Property
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      documents.Document
		wantPanic bool
	}{
		// TODO: Update after #59 fix https://github.com/AssetMantle/modules/issues/59
		{"+ve with no mutation", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}, args{}, document{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}, false},
		{"+ve with nil classificationID", fields{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}, args{[]properties.Property{testMutateProperty}}, document{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}.Mutate(testMutateProperty), false},
		{"+ve with nil immutables", fields{ClassificationID: classificationID, Immutables: nil, Mutables: testMutables}, args{[]properties.Property{testMutateProperty}}, document{ClassificationID: classificationID, Immutables: nil, Mutables: testMutables}.Mutate(testMutateProperty), false},
		{"+ve with nil mutables", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: baseQualified.NewMutables(baseLists.NewPropertyList())}, args{[]properties.Property{testMutateProperty}}, document{ClassificationID: classificationID, Immutables: testImmutables, Mutables: baseQualified.NewMutables(baseLists.NewPropertyList())}.Mutate(testMutateProperty), false}, // TODO: fix this
		{"+ve with all nil", fields{nil, baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList())}, args{[]properties.Property{testMutateProperty}}, document{nil, baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList())}.Mutate(testMutateProperty), false},                            // TODO: fix this
		{"panic case nil document", fields{nil, nil, nil}, args{[]properties.Property{testMutateProperty}}, nil, true},
		{"panic case nil mutables", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: nil}, args{[]properties.Property{testMutateProperty}}, document{ClassificationID: classificationID, Immutables: testImmutables, Mutables: baseQualified.NewMutables(baseLists.NewPropertyList())}.Mutate(testMutateProperty), true}, // TODO: fix this
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := document{
				ClassificationID: tt.fields.ClassificationID,
				Immutables:       tt.fields.Immutables,
				Mutables:         tt.fields.Mutables,
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					document.Mutate(tt.args.propertyList...)
				})
			} else if got := document.Mutate(tt.args.propertyList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}
