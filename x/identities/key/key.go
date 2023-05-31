// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/module"
)

var _ helpers.Key = (*Key)(nil)

func (key *Key) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(key.IdentityID.Bytes())
}
func (*Key) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, Key{})
}
func (key *Key) IsPartial() bool {
	return len(key.IdentityID.GetHashID().Bytes()) == 0
}
func (key *Key) Equals(compareKey helpers.Key) bool {
	if CompareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		// TODO test nil IdentityID case
		return key.IdentityID.Compare(CompareKey.IdentityID) == 0
	}
}
func keyFromInterface(i interface{}) (*Key, error) {
	switch value := i.(type) {
	case *Key:
		return value, nil
	default:
		return &Key{}, errorConstants.IncorrectFormat.Wrapf("incorrect key type")
	}
}

func NewKey(identityID ids.IdentityID) helpers.Key {
	return &Key{IdentityID: identityID.(*baseIDs.IdentityID)}
}

func Prototype() helpers.Key {
	return &Key{baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)}
}
