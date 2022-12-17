// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

type metasKey struct {
	ids.ID
}

var _ helpers.Key = (*metasKey)(nil)

func (key metasKey) RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface("KeyName", (*helpers.Key)(nil))
}

func (key metasKey) GenerateStoreKeyBytes() []byte {
	return MetasStoreKeyPrefix.GenerateStoreKey(key.Bytes())
}
func (metasKey) RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, metasKey{})
}
func (key metasKey) IsPartial() bool {
	return key.ID.(*base.ID).GetDataID().HashId == nil
}
func (key metasKey) Equals(compareKey helpers.Key) bool {
	if CompareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		return key.ID.Compare(CompareKey.ID) == 0
	}
}
func keyFromInterface(i interface{}) (metasKey, error) {
	switch value := i.(type) {
	case metasKey:
		return value, nil
	default:
		return metasKey{}, constants.MetaDataError
	}
}

func NewKey(dataID ids.ID) helpers.Key {
	return metasKey{
		ID: dataID,
	}
}

func Prototype() helpers.Key { return metasKey{} }
