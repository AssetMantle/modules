// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/metas/internal/module"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type key struct {
	ids.MetaID
}

var _ helpers.Key = (*key)(nil)

func (key key) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(key.Bytes())
}
func (key) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, key{})
}
func (key key) IsPartial() bool {
	return len(key.MetaID.Bytes()) == 0
}
func (key key) Equals(compareKey helpers.Key) bool {
	return key.MetaID.Compare(keyFromInterface(compareKey)) == 0
}
func keyFromInterface(i interface{}) key {
	switch value := i.(type) {
	case key:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewKey(metaID ids.MetaID) helpers.Key {
	return key{
		MetaID: metaID,
	}
}

func Prototype() helpers.Key { return key{} }
