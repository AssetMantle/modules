// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type key struct {
	ids.IdentityID
}

var _ helpers.Key = (*key)(nil)

func (key key) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(key.Bytes())
}
func (key) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, key{})
}
func (key key) IsPartial() bool {
	return len(key.IdentityID.GetHashID().Bytes()) == 0
}
func (key key) Equals(compareKey helpers.Key) bool {
	if CompareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		// TODO test nil IdentityID case
		return key.IdentityID.Compare(CompareKey.IdentityID) == 0
	}
}
func keyFromInterface(i interface{}) (key, error) {
	switch value := i.(type) {
	case key:
		return value, nil
	default:
		return key{}, errorConstants.MetaDataError
	}
}

func NewKey(identityID ids.IdentityID) helpers.Key {
	return key{IdentityID: identityID}
}

func Prototype() helpers.Key {
	return key{}
}
