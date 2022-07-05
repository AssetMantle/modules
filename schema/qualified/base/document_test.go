package base

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"reflect"
	"testing"
)

func TestDocument_GetClassificationID(t *testing.T) {

	id := baseIDs.NewID("ID")
	classificationId := baseIDs.NewID("Data")
	testProperty := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data"))
	testMutables := Mutables{baseLists.NewPropertyList(testProperty)}
	testProperty2 := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewHeightData(baseTypes.NewHeight(123)))
	testImmutables := Immutables{baseLists.NewPropertyList(testProperty2)}

	type fields struct {
		ID               ids.ID
		ClassificationID ids.ID
		Immutables       Immutables
		Mutables         Mutables
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		{"Test for GetClassificationId", fields{id, classificationId, testImmutables, testMutables}, classificationId},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := Document{
				ID:               tt.fields.ID,
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

func TestDocument_GetID(t *testing.T) {

	id := baseIDs.NewID("ID")
	classificationId := baseIDs.NewID("Data")
	testProperty := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data"))
	testMutables := Mutables{baseLists.NewPropertyList(testProperty)}
	testProperty2 := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewHeightData(baseTypes.NewHeight(123)))
	testImmutables := Immutables{baseLists.NewPropertyList(testProperty2)}

	type fields struct {
		ID               ids.ID
		ClassificationID ids.ID
		Immutables       Immutables
		Mutables         Mutables
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		{"Test for GetID", fields{id, classificationId, testImmutables, testMutables}, id},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := Document{
				ID:               tt.fields.ID,
				ClassificationID: tt.fields.ClassificationID,
				Immutables:       tt.fields.Immutables,
				Mutables:         tt.fields.Mutables,
			}
			if got := document.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocument_GetProperty(t *testing.T) {

	id := baseIDs.NewID("ID")
	classificationId := baseIDs.NewID("Data")
	testProperty := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data"))
	testMutables := Mutables{baseLists.NewPropertyList(testProperty)}
	testProperty2 := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewHeightData(baseTypes.NewHeight(123)))
	testImmutables := Immutables{baseLists.NewPropertyList(testProperty2)}
	propertyMutable := testMutables.GetMutablePropertyList().GetProperty(testProperty.GetID())
	propertyImmutable := testImmutables.GetImmutablePropertyList().GetProperty(testProperty2.GetID())
	property := baseProperties.NewProperty(baseIDs.NewID("I"), baseData.NewStringData("D")).GetID()

	type fields struct {
		ID               ids.ID
		ClassificationID ids.ID
		Immutables       Immutables
		Mutables         Mutables
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
		{"Test for GetProperty Mutables", fields{id, classificationId, testImmutables, testMutables}, args{testProperty.GetID()}, propertyMutable},
		{"Test for GetProperty Immutables", fields{id, classificationId, testImmutables, testMutables}, args{testProperty2.GetID()}, propertyImmutable},
		{"Test for GetProperty Nil", fields{id, classificationId, testImmutables, testMutables}, args{property}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := Document{
				ID:               tt.fields.ID,
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

func TestDocument_Mutate(t *testing.T) {

	id := baseIDs.NewID("ID")
	classificationId := baseIDs.NewID("Data")
	testProperty := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data"))
	testMutables := Mutables{baseLists.NewPropertyList(testProperty)}
	testProperty2 := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewHeightData(baseTypes.NewHeight(123)))
	testImmutables := Immutables{baseLists.NewPropertyList(testProperty2)}
	mutatedProperty := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data2"))
	testMutatedProperty := Mutables{baseLists.NewPropertyList(mutatedProperty)}
	testDocument := Document{id, classificationId, testImmutables, testMutatedProperty}

	type fields struct {
		ID               ids.ID
		ClassificationID ids.ID
		Immutables       Immutables
		Mutables         Mutables
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
		{"Test for Mutate", fields{id, classificationId, testImmutables, testMutables}, args{[]properties.Property{mutatedProperty}}, testDocument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := Document{
				ID:               tt.fields.ID,
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
