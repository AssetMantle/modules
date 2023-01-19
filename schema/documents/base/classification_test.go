package base

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func Test_new_Classification(t *testing.T) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))

	testClassification := NewClassification(immutables, mutables)
	type fields struct {
		immutables qualified.Immutables
		mutables   qualified.Mutables
	}
	tests := []struct {
		name      string
		fields    fields
		want      documents.Classification
		wantError bool
	}{
		{"+ve", fields{immutables, mutables}, testClassification, false},
		{"nil immutables", fields{nil, mutables}, testClassification, true},
		{"nil mutables", fields{immutables, nil}, testClassification, true},
		{"nil", fields{nil, nil}, testClassification, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantError {
					t.Errorf("lol")
				}
			}()
			if got := NewClassification(tt.fields.immutables, tt.fields.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClassification() = %v, want %v", got, tt.want)
			}
		})
	}
}
