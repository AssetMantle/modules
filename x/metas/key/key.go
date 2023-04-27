// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/AssetMantle/modules/helpers"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/AssetMantle/modules/x/metas/module"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/cosmos/cosmos-sdk/codec"
)

var _ helpers.Key = (*Key)(nil)

func (key *Key) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(key.DataID.Bytes())
}
func (*Key) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, Key{})
}
func (key *Key) IsPartial() bool {
	return len(key.DataID.GetHashID().Bytes()) == 0
}
func (key *Key) Equals(compareKey helpers.Key) bool {
	if CompareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		return key.DataID.Compare(CompareKey.DataID) == 0
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

func NewKey(dataID ids.DataID) helpers.Key {
	return &Key{
		DataID: dataID.(*baseIDs.DataID),
	}
}

func Prototype() helpers.Key { return &Key{baseIDs.PrototypeDataID().(*baseIDs.DataID)} }
