/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"encoding/hex"
	"testing"

	"github.com/persistenceOne/persistenceSDK/utilities/random"

	"github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants/ids"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/mappables/qualified"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Identity_Methods(t *testing.T) {

	classificationID := base.NewID("classificationID")
	defaultImmutableProperties, _ := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	testIdentityID := key.NewIdentityID(classificationID, defaultImmutableProperties)
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("ImmutableData")))
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewStringData("MutableData")))

	testIdentity := NewIdentity(testIdentityID, immutableProperties, mutableProperties)
	require.Equal(t, testIdentity, identity{Document: qualified.Document{ID: testIdentityID, HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties}, HasMutables: baseTraits.HasMutables{Properties: mutableProperties}}})
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
				HasImmutables:    baseTraits.HasImmutables{Properties: base.NewProperties(base.NewProperty(ids.AuthenticationProperty,base.NewListData()))},
				HasMutables:      baseTraits.HasMutables{},
			}},
			args: args{address: randomAccAddress[0]},
			want: false,
		},
		{
			name: "no authentication",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseTraits.HasImmutables{},
				HasMutables:      baseTraits.HasMutables{},
			}},
			args: args{address: randomAccAddress[0]},
			want: false,
		},
		{
			name: "one authentication provisioned positive match",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseTraits.HasImmutables{Properties: base.NewProperties(base.NewProperty(ids.AuthenticationProperty,base.NewListData(base.NewAccAddressData(randomAccAddress[0]))))},
				HasMutables:      baseTraits.HasMutables{},
			}},
			args: args{address: randomAccAddress[0]},
			want: true,
		},
		{
			name: "one authentication provisioned negative match",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseTraits.HasImmutables{Properties: base.NewProperties(base.NewProperty(ids.AuthenticationProperty,base.NewListData(base.NewAccAddressData(randomAccAddress[0]))))},
				HasMutables:      baseTraits.HasMutables{},
			}},
			args: args{address: randomAccAddress[1]},
			want: false,
		},
		{
			name: "multiple authentication provisioned positive match",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseTraits.HasImmutables{Properties: base.NewProperties(base.NewProperty(ids.AuthenticationProperty, base.NewListData(base.NewAccAddressData(randomAccAddress[0]), base.NewAccAddressData(randomAccAddress[1]), base.NewAccAddressData(randomAccAddress[2]), base.NewAccAddressData(randomAccAddress[3]))))},
				HasMutables:      baseTraits.HasMutables{},
			}},
			args: args{address: randomAccAddress[3]},
			want: true,
		},
		{
			name: "multiple authentication provisioned negative match",
			fields: fields{Document: qualified.Document{
				ID:               nil,
				ClassificationID: nil,
				HasImmutables:    baseTraits.HasImmutables{Properties: base.NewProperties(base.NewProperty(ids.AuthenticationProperty,base.NewListData(base.NewAccAddressData(randomAccAddress[0]), base.NewAccAddressData(randomAccAddress[1]), base.NewAccAddressData(randomAccAddress[2]), base.NewAccAddressData(randomAccAddress[3]))))},
				HasMutables:      baseTraits.HasMutables{},
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
