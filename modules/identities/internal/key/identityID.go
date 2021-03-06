/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"

	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/qualified"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var _ types.ID = (*IdentityID)(nil)
var _ helpers.Key = (*IdentityID)(nil)

func (identityID IdentityID) GetStructReference() codec.ProtoMarshaler {
	return &identityID
}
func (identityID IdentityID) Bytes() []byte {
	return append(
		identityID.ClassificationID.Bytes(),
		identityID.HashID.Bytes()...,
	)
}
func (identityID IdentityID) String() string {
	var values []string
	values = append(values, identityID.ClassificationID.String())
	values = append(values, identityID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (identityID IdentityID) Compare(id types.ID) int {
	return bytes.Compare(identityID.Bytes(), id.Bytes())
}
func (identityID IdentityID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(identityID.Bytes())
}
func (IdentityID) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, module.Name, IdentityID{})
}
func (identityID IdentityID) IsPartial() bool {
	return len(identityID.HashID.Bytes()) == 0
}
func (identityID IdentityID) Equals(key helpers.Key) bool {
	id := identityIDFromInterface(key)
	return identityID.Compare(&id) == 0
}

// TODO Pass Classification & then get Classification ID
func NewIdentityID(classificationID types.ID, immutableProperties types.Properties) types.ID {
	return &IdentityID{
		ClassificationID: classificationID,
		HashID:           baseTraits.HasImmutables{Properties: *base.NewProperties(immutableProperties.GetList()...)}.GenerateHashID(),
	}
}
