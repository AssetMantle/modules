// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	documentsSchema "github.com/AssetMantle/schema/go/documents"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/qualified"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/key"
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
			if got := NewMappable(tt.args.Asset); !reflect.DeepEqual(got, tt.want) {
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

func Test_mappable_RegisterCodec(t *testing.T) {
	_, _, _, testMappable := createTestInput()
	type fields struct {
		Document *Mappable
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testMappable}, args{legacyAmino: codec.NewLegacyAmino()}},
		{"+ve nil", fields{&Mappable{nil}}, args{legacyAmino: codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := &Mappable{
				Classification: tt.fields.Document.Classification,
			}
			as.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_mappable_GetKey(t *testing.T) {
	_, _, _, testMappable := createTestInput()
	type fields struct {
		Document *Mappable
	}
	tests := []struct {
		name      string
		fields    fields
		want      helpers.Key
		wantPanic bool
	}{
		{"+ve", fields{testMappable}, key.NewKey(baseIDs.NewClassificationID(testMappable.Classification.GetImmutables(), testMappable.Classification.GetMutables())), false},
		{"panic case nil", fields{&Mappable{nil}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := &Mappable{
				Classification: tt.fields.Document.Classification,
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					asset.GenerateKey()
				})
			} else if got := asset.GenerateKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
