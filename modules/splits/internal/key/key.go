// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/splits/internal/module"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

var _ helpers.Key = (*Key)(nil)

func (key *Key) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(key.SplitID.Bytes())
}
func (*Key) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, &Key{})
}
func (key *Key) IsPartial() bool {
	return len(key.SplitID.GetOwnableID().Bytes()) == 0
}
func (key *Key) Equals(compareKey helpers.Key) bool {
	if compareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		return key.SplitID.Compare(compareKey.SplitID) == 0
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

func NewKey(splitID ids.SplitID) helpers.Key {
	return &Key{
		SplitID: splitID.(*baseIDs.SplitID),
	}
}

func Prototype() helpers.Key {
	return &Key{baseIDs.PrototypeSplitID().(*baseIDs.SplitID)}
}
