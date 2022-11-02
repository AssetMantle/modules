// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/maintainers/internal/module"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type key struct {
	ids.MaintainerID
}

var _ helpers.Key = (*key)(nil)

func (key key) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(key.Bytes())
}
func (key) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, key{})
}
func (key key) IsPartial() bool {
	return len(key.MaintainerID.GetClassificationID().Bytes()) == 0
}
func (key key) Equals(compareKey helpers.Key) bool {
	if CompareKey, err := keyFromInterface(key); err != nil {
		return false
	} else {
		return key.MaintainerID.Compare(CompareKey.MaintainerID) == 0
	}
}
func keyFromInterface(i interface{}) (key, error) {
	switch value := i.(type) {
	case key:
		return value, nil
	default:
		return key{}, constants.MetaDataError
	}
}

func NewKey(maintainerID ids.MaintainerID) helpers.Key {
	return key{
		MaintainerID: maintainerID,
	}
}

func Prototype() helpers.Key {
	return key{}
}
