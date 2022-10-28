// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	base2 "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
)

func createTestInput() (ids.ClassificationID, qualified.Immutables, qualified.Mutables, qualified.Document) {
	testImmutables := NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	testMutables := NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(testImmutables, testMutables)
	testDocument := base2.NewDocument(classificationID, testImmutables, testMutables)
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
		want qualified.Document
	}{
		// TODO: Add test cases.
		{"+ve", args{classificationID: classificationID, immutables: testImmutables, mutables: testMutables}, base2.document{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}},
		{"+ve with nil classificationID", args{classificationID: nil, immutables: testImmutables, mutables: testMutables}, base2.document{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}},
		{"+ve with nil immutables", args{classificationID: classificationID, immutables: nil, mutables: testMutables}, base2.document{ClassificationID: classificationID, Immutables: nil, Mutables: testMutables}},
		{"+ve with nil mutables", args{classificationID: classificationID, immutables: testImmutables, mutables: nil}, base2.document{ClassificationID: classificationID, Immutables: testImmutables, Mutables: nil}},
		{"+ve with all nil", args{}, base2.document{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base2.NewDocument(tt.args.classificationID, tt.args.immutables, tt.args.mutables); !reflect.DeepEqual(got, tt.want) {
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
			document := base2.document{
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
			document := base2.document{
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
			document := base2.document{
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
		name   string
		fields fields
		args   args
		want   properties.Property
	}{
		// TODO: Update unit test after issue # fix
		{"+ve for Immutable Property", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}, args{testImmutablePropertyID}, testImmutables.GetImmutablePropertyList().GetProperty(testImmutablePropertyID)},
		{"+ve with nil classificationID for Immutable Property", fields{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}, args{testImmutablePropertyID}, testImmutables.GetImmutablePropertyList().GetProperty(testImmutablePropertyID)},
		{"+ve with nil immutables for Immutable Property", fields{ClassificationID: classificationID, Immutables: immutables{}, Mutables: testMutables}, args{testImmutablePropertyID}, nil}, // TODO: panics for empty immutable struct
		{"+ve with nil mutables for Immutable Property", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: mutables{}}, args{testImmutablePropertyID}, testImmutables.GetImmutablePropertyList().GetProperty(testImmutablePropertyID)},
		{"+ve with all nil", fields{nil, immutables{}, mutables{}}, args{testImmutablePropertyID}, nil}, // TODO: panics for empty immutable struct
		{"+ve for Mutable Property", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}, args{testMutablePropertyID}, testMutables.GetMutablePropertyList().GetProperty(testMutablePropertyID)},
		{"+ve with nil classificationID for Mutable Property", fields{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}, args{testMutablePropertyID}, testMutables.GetMutablePropertyList().GetProperty(testMutablePropertyID)},
		{"+ve with nil immutables for Mutable Property", fields{ClassificationID: classificationID, Immutables: immutables{}, Mutables: testMutables}, args{testMutablePropertyID}, testMutables.GetMutablePropertyList().GetProperty(testMutablePropertyID)}, // TODO: panics for empty immutable struct
		{"+ve with nil mutables for Mutable Property", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: mutables{}}, args{testMutablePropertyID}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := base2.document{
				ClassificationID: tt.fields.ClassificationID,
				Immutables:       tt.fields.Immutables,
				Mutables:         tt.fields.Mutables,
			}
			if got := document.GetProperty(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
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
		name   string
		fields fields
		args   args
		want   qualified.Document
	}{
		// TODO: Update after #59 fix https://github.com/AssetMantle/modules/issues/59
		{"+ve with no mutation", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}, args{}, base2.document{ClassificationID: classificationID, Immutables: testImmutables, Mutables: testMutables}},
		{"+ve with nil classificationID", fields{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}, args{[]properties.Property{testMutateProperty}}, base2.document{ClassificationID: nil, Immutables: testImmutables, Mutables: testMutables}.Mutate(testMutateProperty)},
		{"+ve with nil immutables", fields{ClassificationID: classificationID, Immutables: nil, Mutables: testMutables}, args{[]properties.Property{testMutateProperty}}, base2.document{ClassificationID: classificationID, Immutables: nil, Mutables: testMutables}.Mutate(testMutateProperty)},
		{"+ve with nil mutables", fields{ClassificationID: classificationID, Immutables: testImmutables, Mutables: mutables{}}, args{[]properties.Property{testMutateProperty}}, base2.document{ClassificationID: classificationID, Immutables: testImmutables, Mutables: mutables{}}.Mutate(testMutateProperty)}, // TODO: fix this
		{"+ve with all nil", fields{nil, immutables{}, mutables{}}, args{[]properties.Property{testMutateProperty}}, base2.document{nil, immutables{}, mutables{}}.Mutate(testMutateProperty)},                                                                                                                    // TODO: fix this
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := base2.document{
				ClassificationID: tt.fields.ClassificationID,
				Immutables:       tt.fields.Immutables,
				Mutables:         tt.fields.Mutables,
			}
			if got := document.Mutate(tt.args.propertyList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}
