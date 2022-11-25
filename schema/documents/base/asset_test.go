package base

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	documentsSchema "github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func TestNewAsset(t *testing.T) {
	classificationID, immutables, mutables, _ := createTestInput()
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
		mutables         qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want documentsSchema.Asset
	}{
		// TODO: Add test cases.
		{"+ve", args{classificationID: classificationID, immutables: immutables, mutables: mutables}, asset{Document: NewDocument(classificationID, immutables, mutables)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAsset(tt.args.classificationID, tt.args.immutables, tt.args.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAsset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asset_GetBurn(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithBurn := NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(constants.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))))))

	type fields struct {
		Document documentsSchema.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		// TODO: Add test cases.
		{"+ve", fields{Document: testDocumentWithBurn}, baseProperties.NewMesaProperty(constants.BurnHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1)))},
		{"+ve", fields{Document: testDocument}, constants.BurnHeightProperty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := asset{
				Document: tt.fields.Document,
			}
			if got := asset.GetBurn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBurn() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_asset_GetLock(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithLock := NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(constants.LockProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1))))))

	type fields struct {
		Document documentsSchema.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		// TODO: Add test cases.
		{"+ve with default lock", fields{testDocument}, constants.LockProperty},
		{"+ve with mutated", fields{testDocumentWithLock}, baseProperties.NewMesaProperty(constants.LockProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(1)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := asset{
				Document: tt.fields.Document,
			}
			if got := asset.GetLock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asset_GetSupply(t *testing.T) {
	classificationID, immutables, _, testDocument := createTestInput()
	testDocumentWithSupply := NewDocument(classificationID, immutables, baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(constants.SupplyProperty.GetKey(), baseData.NewDecData(types.NewDec(1))))))
	type fields struct {
		Document documentsSchema.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		// TODO: Add test cases.
		{"+ve", fields{testDocument}, constants.SupplyProperty},
		{"+ve", fields{testDocumentWithSupply}, baseProperties.NewMesaProperty(constants.SupplyProperty.GetKey(), baseData.NewDecData(types.NewDec(1)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := asset{
				Document: tt.fields.Document,
			}
			if got := asset.GetSupply(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSupply() = %v, want %v", got, tt.want)
			}
		})
	}
}
