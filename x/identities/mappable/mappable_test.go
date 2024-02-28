// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/qualified"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/helpers"
)

func createTestInput() (documents.Identity, ids.ClassificationID, qualified.Immutables, qualified.Mutables) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testIdentity := baseDocuments.NewIdentity(classificationID, immutables, mutables)

	return testIdentity, classificationID, immutables, mutables
}

func TestNewMappable(t *testing.T) {
	_, classificationID, immutables, mutables := createTestInput()
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
		mutables         qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want helpers.Mappable
	}{
		{"+ve", args{classificationID, immutables, mutables}, &Mappable{Identity: baseDocuments.NewIdentity(classificationID, immutables, mutables).Get().(*baseDocuments.Document)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMappable(baseDocuments.NewIdentity(tt.args.classificationID, tt.args.immutables, tt.args.mutables)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdentity() = %v, want %v", got, tt.want)
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

func Test_identity_RegisterCodec(t *testing.T) {
	testIdentity, _, _, _ := createTestInput()

	type fields struct {
		Document documents.Identity
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testIdentity}, args{codec.NewLegacyAmino()}},
		{"+ve nil", fields{nil}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &Mappable{
				Identity: tt.fields.Document.Get().(*baseDocuments.Document),
			}
			id.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
