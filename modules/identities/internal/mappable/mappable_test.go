// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/modules/identities/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	base3 "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
)

func createTestInput() (documents.Identity, ids.ClassificationID, qualified.Immutables, qualified.Mutables) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(base.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(base.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := base.NewClassificationID(immutables, mutables)
	testIdentity := base3.NewIdentity(classificationID, immutables, mutables)

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
		want mappable
	}{
		{"+ve", args{classificationID, immutables, mutables}, mappable{Identity: base3.NewIdentity(classificationID, immutables, mutables)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMappable(base3.NewIdentity(tt.args.classificationID, tt.args.immutables, tt.args.mutables)); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", fields{testIdentity}, key.NewKey(base.NewIdentityID(testIdentity.GetClassificationID(), testIdentity.GetImmutables())), false},
		{"panic case nil", fields{nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := mappable{
				Identity: tt.fields.Document,
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
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testIdentity}, args{codec.New()}},
		{"+ve nil", fields{nil}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := mappable{
				Identity: tt.fields.Document,
			}
			id.RegisterCodec(tt.args.codec)
		})
	}
}
