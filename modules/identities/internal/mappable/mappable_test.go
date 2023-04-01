// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/identities/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
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

func Test_identity_GetKey(t *testing.T) {
	testIdentity, _, _, _ := createTestInput()

	type fields struct {
		Document documents.Identity
	}
	tests := []struct {
		name      string
		fields    fields
		want      helpers.Key
		wantPanic bool
	}{
		{"+ve", fields{testIdentity}, key.NewKey(baseIDs.NewIdentityID(testIdentity.GetClassificationID(), testIdentity.GetImmutables())), false},
		{"panic case nil", fields{nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := &Mappable{
				Identity: tt.fields.Document.Get().(*baseDocuments.Document),
			}
			if tt.wantPanic {
				require.Panics(t, func() {
					identity.GetKey()
				})
			} else if got := identity.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
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
