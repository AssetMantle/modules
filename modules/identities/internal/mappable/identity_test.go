package mappable

import (
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"reflect"
	"testing"
)

func TestNewIdentity(t *testing.T) {
	_, _, testIdentityID, immutableProperties, mutableProperties := initalizeVariables()
	type args struct {
		id                  ids.ID
		immutableProperties lists.PropertyList
		mutableProperties   lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want mappables.Identity
	}{
		// TODO: Add test cases.
		{"+ve", args{testIdentityID, immutableProperties, mutableProperties}, identity{Document: baseQualified.Document{ID: testIdentityID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}},
		{"-ve", args{testIdentityID, immutableProperties, mutableProperties}, identity{Document: baseQualified.Document{ID: testIdentityID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIdentity(tt.args.id, tt.args.immutableProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_GetAuthentication(t *testing.T) {
	_, _, testIdentityID, immutableProperties, mutableProperties := initalizeVariables()
	type fields struct {
		Document mappables.Identity
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		// TODO: Add test cases.
		{"+ve", fields{NewIdentity(testIdentityID, immutableProperties, mutableProperties)}, constants.Authentication},
		//{"-ve", fields{NewIdentity(testIdentityID, immutableProperties, mutableProperties)}, constants.Expiry},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.Document.GetAuthentication(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthentication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_GetExpiry(t *testing.T) {
	_, _, testIdentityID, immutableProperties, mutableProperties := initalizeVariables()

	type fields struct {
		Document baseQualified.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   properties.Property
	}{
		// TODO: Add test cases.
		// TODO:
		//{"+ve", fields{baseQualified.Document{ID: testIdentityID, ClassificationID: classificationID, Immutables: baseQualified.Immutables{PropertyList: defaultImmutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}, baseQualified.Document{ID: testIdentityID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}.GetProperty(constants.ExpiryProperty)},
		{"+ve for nil property", fields{baseQualified.Document{ID: testIdentityID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}, constants.Expiry},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := identity{
				Document: tt.fields.Document,
			}
			if got := identity.GetExpiry(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpiry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_GetKey(t *testing.T) {
	_, _, testIdentityID, immutableProperties, mutableProperties := initalizeVariables()
	type fields struct {
		Document baseQualified.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.Key
	}{
		// TODO: Add test cases.
		{"+ve", fields{baseQualified.Document{ID: testIdentityID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}, key.FromID(testIdentityID)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := identity{
				Document: tt.fields.Document,
			}
			if got := identity.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_RegisterCodec(t *testing.T) {
	_, _, testIdentityID, immutableProperties, mutableProperties := initalizeVariables()

	type fields struct {
		Document baseQualified.Document
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"+ve register codec", fields{baseQualified.Document{ID: testIdentityID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := identity{
				Document: tt.fields.Document,
			}
			id.RegisterCodec(tt.args.codec)
		})
	}
}

func initalizeVariables() (ids.ID, lists.PropertyList, ids.ID, lists.PropertyList, lists.PropertyList) {
	classificationID := baseIDs.NewID("classificationID")
	defaultImmutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	testIdentityID := key.NewIdentityID(classificationID, defaultImmutableProperties)
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("MutableData")))
	return classificationID, defaultImmutableProperties, testIdentityID, immutableProperties, mutableProperties
}
