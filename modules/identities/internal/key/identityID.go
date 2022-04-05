// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	baseTraits "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
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
func (identityID identityID) Compare(id types.ID) int {
	return bytes.Compare(identityID.Bytes(), id.Bytes())
}
func (identityID identityID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(identityID.Bytes())
}
func (identityID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, identityID{})
}
func (identityID identityID) IsPartial() bool {
	return len(identityID.HashID.Bytes()) == 0
}
func (identityID identityID) Equals(key helpers.Key) bool {
	return identityID.Compare(identityIDFromInterface(key)) == 0
}

// TODO Pass Classification & then get Classification ID
func NewIdentityID(classificationID types.ID, immutableProperties types.Properties) types.ID {
	return identityID{
		ClassificationID: classificationID,
		HashID:           baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID(),
	}
}
