/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"strings"

	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type identityID struct {
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	HashID           types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*identityID)(nil)
var _ helpers.Key = (*identityID)(nil)

func (identityID identityID) Bytes() []byte {
	return append(
		identityID.ClassificationID.Bytes(),
		identityID.HashID.Bytes()...,
	)
}
func (identityID identityID) String() string {
	var values []string
	values = append(values, identityID.ClassificationID.String())
	values = append(values, identityID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (identityID identityID) Equals(id types.ID) bool {
	return bytes.Equal(identityID.Bytes(), id.Bytes())
}
func (identityID identityID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(identityID.Bytes())
}
func (identityID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, identityID{})
}
func (identityID identityID) IsPartial() bool {
	return len(identityID.HashID.Bytes()) == 0
}
func (identityID identityID) Matches(key helpers.Key) bool {
	return identityID.Equals(identityIDFromInterface(key))
}

func NewIdentityID(classificationID types.ID, immutableProperties types.Properties) types.ID {
	return identityID{
		ClassificationID: classificationID,
		HashID:           baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID(),
	}
}
