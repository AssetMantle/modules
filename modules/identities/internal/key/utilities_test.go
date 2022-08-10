package key

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"reflect"
	"testing"
)

func TestFromID(t *testing.T) {
	type args struct {
		id ids.ID
	}
	tests := []struct {
		name string
		args args
		want helpers.Key
	}{
		// TODO: Add test cases.
		{"+ve empty idString", args{baseIDs.NewID("")}, identityIDFromInterface(baseIDs.NewID(""))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identityIDFromInterface(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := utilities.ReadProperties("")
	type args struct {
		i interface{}
	}
	tests := []struct {
		name      string
		args      args
		want      identityID
		wantPanic bool
	}{
		// TODO: Add test cases.
		{"+ve for id.Ids", args{NewIdentityID(classificationID, immutableProperties)}, identityID{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}, false},
		{"-ve for panic", args{immutableProperties}, identityID{baseIDs.NewID(""), baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}, true},
		{"+ve for identity{}", args{baseIDs.NewID("|")}, identityID{baseIDs.NewID(""), baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}, false},
		//{"+ve", args{baseIDs.NewID("test")}, identityID{baseIDs.NewID(""), baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("identityIDFromInterface() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := identityIDFromInterface(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("identityIDFromInterface() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_readIdentityID(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := utilities.ReadProperties("")
	type args struct {
		identityIDString string
	}
	tests := []struct {
		name string
		args args
		want ids.ID
	}{
		// TODO: Add test cases.
		{"-ve for identity{} with empty string", args{""}, identityID{baseIDs.NewID(""), baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}},
		{"+ve for identity{}", args{"|"}, identityID{baseIDs.NewID(""), baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}},
		{"+ve for identity{}", args{"classificationID|" + baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID().String()}, identityID{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readIdentityID(tt.args.identityIDString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readIdentityID() = %v, want %v", got, tt.want)
			}
		})
	}
}
