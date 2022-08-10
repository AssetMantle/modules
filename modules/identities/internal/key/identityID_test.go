package key

import (
	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/cosmos/cosmos-sdk/codec"
	"reflect"
	"testing"
)

func TestNewIdentityID(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := utilities.ReadProperties("")

	type args struct {
		classificationID    ids.ID
		immutableProperties lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", args{classificationID: classificationID, immutableProperties: immutableProperties}, identityID{ClassificationID: classificationID, HashID: baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}},
		{"empty immutables properties", args{classificationID: classificationID, immutableProperties: emptyImmutableProperties}, identityID{ClassificationID: classificationID, HashID: baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIdentityID(tt.args.classificationID, tt.args.immutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdentityID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identityID_Bytes(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := utilities.ReadProperties("")
	type fields struct {
		ClassificationID ids.ID
		HashID           ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}, append(classificationID.Bytes(), baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID().Bytes()...)},
		{"+ve empty immutable Properties", fields{classificationID, baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}, append(classificationID.Bytes(), baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID().Bytes()...)},
		//{"+ve nil", fields{classificationID, baseQualified.Immutables{PropertyList: nil}.GenerateHashID()}, append(classificationID.Bytes(), baseQualified.Immutables{PropertyList: nil}.GenerateHashID().Bytes()...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identityID := identityID{
				ClassificationID: tt.fields.ClassificationID,
				HashID:           tt.fields.HashID,
			}
			if got := identityID.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identityID_Compare(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := utilities.ReadProperties("")
	type fields struct {
		ClassificationID ids.ID
		HashID           ids.ID
	}
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{"+ve Equal", fields{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}, args{identityID{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}}, 0},
		{"+ve Not Equal", fields{classificationID, baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}, args{identityID{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}}, -1},
		//{"+ve", fields{classificationID, baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}, args{nil}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identityID := identityID{
				ClassificationID: tt.fields.ClassificationID,
				HashID:           tt.fields.HashID,
			}
			if got := identityID.Compare(tt.args.listable); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identityID_Equals(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := utilities.ReadProperties("")
	type fields struct {
		ClassificationID ids.ID
		HashID           ids.ID
	}
	type args struct {
		key helpers.Key
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"+ve Equal", fields{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}, args{FromID(NewIdentityID(classificationID, immutableProperties))}, true},
		{"-ve Not Equal", fields{classificationID, baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}, args{FromID(NewIdentityID(classificationID, immutableProperties))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identityID := identityID{
				ClassificationID: tt.fields.ClassificationID,
				HashID:           tt.fields.HashID,
			}
			if got := identityID.Equals(tt.args.key); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identityID_GenerateStoreKeyBytes(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := utilities.ReadProperties("")
	type fields struct {
		ClassificationID ids.ID
		HashID           ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{"+ve", fields{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}, module.StoreKeyPrefix.GenerateStoreKey(identityID{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}.Bytes())},
		{"+ve for empty ImmutableProperties", fields{classificationID, baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}, module.StoreKeyPrefix.GenerateStoreKey(identityID{classificationID, baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}.Bytes())},
		//{"-ve", fields{classificationID, baseQualified.Immutables{PropertyList: nil}.GenerateHashID()}, module.StoreKeyPrefix.GenerateStoreKey(identityID{classificationID, baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}.Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identityID := identityID{
				ClassificationID: tt.fields.ClassificationID,
				HashID:           tt.fields.HashID,
			}
			if got := identityID.GenerateStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStoreKeyBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identityID_IsPartial(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := utilities.ReadProperties("")
	type fields struct {
		ClassificationID ids.ID
		HashID           ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"+ve for non empty HashID", fields{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}, false},
		{"+ve for empty HashID", fields{classificationID, baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identityID := identityID{
				ClassificationID: tt.fields.ClassificationID,
				HashID:           tt.fields.HashID,
			}
			if got := identityID.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identityID_RegisterCodec(t *testing.T) {
	Codec := codec.New()
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	type fields struct {
		ClassificationID ids.ID
		HashID           ids.ID
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
		{"+ve", fields{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}, args{Codec}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := identityID{
				ClassificationID: tt.fields.ClassificationID,
				HashID:           tt.fields.HashID,
			}
			id.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_identityID_String(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := utilities.ReadProperties("")
	type fields struct {
		ClassificationID ids.ID
		HashID           ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{classificationID, baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}, "classificationID|xex68KrzcI4UhOXIC4G1OhgcVpaYBUMh_I0UMDrxARI="},
		{"+ve empty immutableProperties", fields{classificationID, baseQualified.Immutables{PropertyList: emptyImmutableProperties}.GenerateHashID()}, "classificationID|"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identityID := identityID{
				ClassificationID: tt.fields.ClassificationID,
				HashID:           tt.fields.HashID,
			}
			if got := identityID.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
