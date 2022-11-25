// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	documentsSchema "github.com/AssetMantle/modules/schema/documents"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func createTestInput() (ids.ClassificationID, qualified.Immutables, qualified.Mutables, mappable) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testMappable := mappable{Classification: baseDocuments.NewDocument(classificationID, immutables, mutables)}
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
		want documentsSchema.Document
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
		{"+ve", mappable{}},
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
		Document mappable
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testMappable}, args{codec: codec.New()}},
		{"+ve nil", fields{mappable{nil}}, args{codec: codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := mappable{
				Classification: tt.fields.Document,
			}
			as.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_mappable_GetKey(t *testing.T) {
	_, _, _, testMappable := createTestInput()
	type fields struct {
		Document mappable
	}
	tests := []struct {
		name      string
		fields    fields
		want      helpers.Key
		wantPanic bool
	}{
		{"+ve", fields{testMappable}, key.NewKey(baseIDs.NewClassificationID(testMappable.GetImmutables(), testMappable.GetMutables())), false},
		{"panic case nil", fields{mappable{nil}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := mappable{
				Classification: tt.fields.Document,
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					asset.GetKey()
				})
			} else if got := asset.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
