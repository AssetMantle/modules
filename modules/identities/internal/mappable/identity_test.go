// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"encoding/hex"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/random"
)

func Test_Identity_Methods(t *testing.T) {

	classificationID := baseIDs.NewID("classificationID")
	defaultImmutableProperties, _ := baseTypes.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	testIdentityID := key.NewIdentityID(classificationID, defaultImmutableProperties)
	immutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	mutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("MutableData")))

	testIdentity := NewIdentity(testIdentityID, immutableProperties, mutableProperties)
	require.Equal(t, testIdentity, identity{Document: qualified.Document{ID: testIdentityID, HasImmutables: baseQualified.HasImmutables{Properties: immutableProperties}, HasMutables: baseQualified.HasMutables{Properties: mutableProperties}}})
	require.Equal(t, testIdentity.(identity).GetID(), testIdentityID)
	require.Equal(t, testIdentity.GetImmutableProperties(), immutableProperties)
	require.Equal(t, testIdentity.GetMutableProperties(), mutableProperties)
	require.Equal(t, testIdentity.GetKey(), testIdentityID)
}

func Test_identity_IsProvisioned(t *testing.T) {
	randomAccAddress := make([]sdkTypes.AccAddress, 5)
	for i := range randomAccAddress {
		randomAccAddress[i], _ = sdkTypes.AccAddressFromHex(hex.EncodeToString([]byte(random.GenerateUniqueIdentifier())))
	}

	type fields struct {
		Document qualified.Document
	}
	type args struct {
		address types.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "empty authentication",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseQualified.HasImmutables{Properties: baseTypes.NewProperties(baseTypes.NewProperty(ids.AuthenticationProperty, baseData.NewListData()))},
				HasMutables:      baseQualified.HasMutables{},
			}},
			args: args{address: randomAccAddress[0]},
			want: false,
		},
		{
			name: "no authentication",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseQualified.HasImmutables{},
				HasMutables:      baseQualified.HasMutables{},
			}},
			args: args{address: randomAccAddress[0]},
			want: false,
		},
		{
			name: "one authentication provisioned positive match",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseQualified.HasImmutables{Properties: baseTypes.NewProperties(baseTypes.NewProperty(ids.AuthenticationProperty, baseData.NewListData(baseData.NewAccAddressData(randomAccAddress[0]))))},
				HasMutables:      baseQualified.HasMutables{},
			}},
			args: args{address: randomAccAddress[0]},
			want: true,
		},
		{
			name: "one authentication provisioned negative match",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseQualified.HasImmutables{Properties: baseTypes.NewProperties(baseTypes.NewProperty(ids.AuthenticationProperty, baseData.NewListData(baseData.NewAccAddressData(randomAccAddress[0]))))},
				HasMutables:      baseQualified.HasMutables{},
			}},
			args: args{address: randomAccAddress[1]},
			want: false,
		},
		{
			name: "multiple authentication provisioned positive match",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseQualified.HasImmutables{Properties: baseTypes.NewProperties(baseTypes.NewProperty(ids.AuthenticationProperty, baseData.NewListData(baseData.NewAccAddressData(randomAccAddress[0]), baseData.NewAccAddressData(randomAccAddress[1]), baseData.NewAccAddressData(randomAccAddress[2]), baseData.NewAccAddressData(randomAccAddress[3]))))},
				HasMutables:      baseQualified.HasMutables{},
			}},
			args: args{address: randomAccAddress[3]},
			want: true,
		},
		{
			name: "multiple authentication provisioned negative match",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseQualified.HasImmutables{Properties: baseTypes.NewProperties(baseTypes.NewProperty(ids.AuthenticationProperty, baseData.NewListData(baseData.NewAccAddressData(randomAccAddress[0]), baseData.NewAccAddressData(randomAccAddress[1]), baseData.NewAccAddressData(randomAccAddress[2]), baseData.NewAccAddressData(randomAccAddress[3]))))},
				HasMutables:      baseQualified.HasMutables{},
			}},
			args: args{address: randomAccAddress[4]},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := identity{
				Document: tt.fields.Document,
			}
			if got := identity.IsProvisioned(tt.args.address); got != tt.want {
				t.Errorf("IsProvisioned() = %v, want %v", got, tt.want)
			}
		})
	}
}
