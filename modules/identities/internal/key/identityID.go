/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"strings"
)

type identityID struct {
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	HashID           types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*identityID)(nil)
var _ helpers.Key = (*identityID)(nil)

func (IdentityID identityID) Bytes() []byte {
	return append(
		IdentityID.ClassificationID.Bytes(),
		IdentityID.HashID.Bytes()...,
	)
}
func (IdentityID identityID) String() string {
	var values []string
	values = append(values, IdentityID.ClassificationID.String())
	values = append(values, IdentityID.HashID.String())
	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (IdentityID identityID) Equals(id types.ID) bool {
	return bytes.Compare(IdentityID.Bytes(), id.Bytes()) == 0
}
func (IdentityID identityID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(IdentityID.Bytes())
}
func (identityID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, identityID{})
}
func (IdentityID identityID) IsPartial() bool {
	if len(IdentityID.HashID.Bytes()) > 0 {
		return false
	}
	return true
}
func (IdentityID identityID) Matches(key helpers.Key) bool {
	switch value := key.(type) {
	case identityID:
		return bytes.Compare(IdentityID.Bytes(), value.Bytes()) == 0
	default:
		return false
	}
}

func New(id types.ID) helpers.Key {
	return identityIDFromInterface(id)
}

func NewIdentityID(classificationID types.ID, hashID types.ID) types.ID {
	return identityID{
		ClassificationID: classificationID,
		HashID:           hashID,
	}
}
